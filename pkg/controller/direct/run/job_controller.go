// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package run

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "cloud.google.com/go/run/apiv2"

	runpb "cloud.google.com/go/run/apiv2/runpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func init() {
	registry.RegisterModel(krm.RunJobGVK, NewJobModel)
}

func NewJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelJob{config: *config}, nil
}

var _ directbase.Model = &modelJob{}

type modelJob struct {
	config config.ControllerConfig
}

func (m *modelJob) client(ctx context.Context) (*gcp.JobsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewJobsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Job client: %w", err)
	}
	return gcpClient, err
}

func (m *modelJob) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.RunJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewJobIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	// TODO: this should not block DELETION.
	if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
		return nil, err
	}

	// Get run GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &JobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//run.googleapis.com/"); ok {
		// Direct controller for RunJob only handles v2.
		s = strings.TrimPrefix(s, "v2/")

		var id krm.JobIdentity
		if err := id.FromExternal(s); err != nil {
			log.V(2).Error(err, "url did not match RunJob format", "url", url)
			return nil, nil
		}

		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &JobAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type JobAdapter struct {
	id        *krm.JobIdentity
	gcpClient *gcp.JobsClient
	desired   *krm.RunJob
	actual    *runpb.Job
}

var _ directbase.Adapter = &JobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *JobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Job", "name", a.id)

	req := &runpb.GetJobRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Job %q: %w", a.id, err)
	}

	a.actual = found
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *JobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Job", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := RunJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &runpb.CreateJobRequest{
		Parent: a.id.Parent().String(),
		Job:    resource,
		JobId:  a.id.ID(),
	}
	op, err := a.gcpClient.CreateJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Job %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of job %q: %w", a.id, err)
	}
	log.V(2).Info("successfully created Job", "name", a.id)

	status := &krm.RunJobStatus{}
	status.ObservedState = RunJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	specHash, err := common.HashSpec(&a.desired.Spec)
	if err != nil {
		return fmt.Errorf("calculating spec hash: %w", err)
	}
	gcpHash, err := common.HashProto(created)
	if err != nil {
		return fmt.Errorf("calculating gcp hash: %w", err)
	}
	status.LastModifiedCookie = direct.LazyPtr(fmt.Sprintf("%s/%s", specHash, gcpHash))

	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *JobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Job", "name", a.id)

	specHash, err := common.HashSpec(&a.desired.Spec)
	if err != nil {
		return fmt.Errorf("calculating spec hash: %w", err)
	}
	gcpHash, err := common.HashProto(a.actual)
	if err != nil {
		return fmt.Errorf("calculating gcp hash: %w", err)
	}
	cookie := fmt.Sprintf("%s/%s", specHash, gcpHash)

	if a.desired.Status.LastModifiedCookie != nil {
		if *a.desired.Status.LastModifiedCookie == cookie {
			log.V(2).Info("resource is up to date", "name", a.id)
			// Fast path: Update status with observed state and return.
			// The status update is important to update conditions and observedGeneration.
			status := &krm.RunJobStatus{}
			mapCtx := &direct.MapContext{}
			status.ObservedState = RunJobObservedState_FromProto(mapCtx, a.actual)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}
			status.LastModifiedCookie = a.desired.Status.LastModifiedCookie // Preserve cookie
			return updateOp.UpdateStatus(ctx, status, nil)
		}
	}

	mapCtx := &direct.MapContext{}

	var desiredPb *runpb.Job
	if _, ok := a.desired.GetAnnotations()[k8s.DefaultToGCPFieldsAnnotation]; ok {
		// If the annotation is present, we need to build an effective desired state.
		effectiveDesired, err := buildEffectiveDesiredRunJob(a.desired, a.actual)
		if err != nil {
			return err
		}
		desiredPb = effectiveDesired
	} else {
		desiredPb = RunJobSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}
	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)

		req := &runpb.UpdateJobRequest{
			Job: desiredPb,
		}
		op, err := a.gcpClient.UpdateJob(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Job %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Job %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Job", "name", a.id)
	}

	status := &krm.RunJobStatus{}
	status.ObservedState = RunJobObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	newGCPHash, err := common.HashProto(updated)
	if err != nil {
		return fmt.Errorf("calculating new gcp hash: %w", err)
	}
	status.LastModifiedCookie = direct.LazyPtr(fmt.Sprintf("%s/%s", specHash, newGCPHash))

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *JobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.RunJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RunJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = &a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.RunJobGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *JobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Job", "name", a.id)

	name := a.id.String()
	req := &runpb.DeleteJobRequest{Name: name}
	op, err := a.gcpClient.DeleteJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Job, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Job %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Job", "name", a.id)

	if _, err = op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting delete Job %s: %w", a.id, err)
	}
	return true, nil
}

// buildEffectiveDesiredRunJob returns a new Job proto by overlaying the desired spec
// onto the actual state for fields that are not specified in the spec but are listed
// in the `cnrm.cloud.google.com/default-to-gcp-fields` annotation.
func buildEffectiveDesiredRunJob(desired *krm.RunJob, actual *runpb.Job) (*runpb.Job, error) {
	// Start with the desired state converted to a proto.
	mapCtx := &direct.MapContext{}
	effective := RunJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	annotations := desired.GetAnnotations()
	fieldsToDefault, ok := annotations[k8s.DefaultToGCPFieldsAnnotation]
	if !ok {
		// If the annotation is not present, the initial desired state is the effective state.
		return effective, nil
	}

	fieldSet := make(map[string]bool)
	for _, field := range strings.Split(fieldsToDefault, ",") {
		fieldSet[strings.TrimSpace(field)] = true
	}

	// Now, for each field in the annotation, check if it was specified in the spec.
	// If not, copy the value from the actual state (GCP) to our effective desired state.

	if fieldSet["spec.annotations"] && desired.Spec.Annotations == nil {
		effective.Annotations = actual.Annotations
	}
	if fieldSet["spec.binaryAuthorization"] && desired.Spec.BinaryAuthorization == nil {
		effective.BinaryAuthorization = actual.BinaryAuthorization
	} else if desired.Spec.BinaryAuthorization != nil {
		if fieldSet["spec.binaryAuthorization.useDefault"] && desired.Spec.BinaryAuthorization.UseDefault == nil {
			if effective.BinaryAuthorization == nil {
				effective.BinaryAuthorization = &runpb.BinaryAuthorization{}
			}
			effective.BinaryAuthorization.BinauthzMethod = actual.GetBinaryAuthorization().GetBinauthzMethod()
		}
		if fieldSet["spec.binaryAuthorization.breakglassJustification"] && desired.Spec.BinaryAuthorization.BreakglassJustification == nil {
			if effective.BinaryAuthorization == nil {
				effective.BinaryAuthorization = &runpb.BinaryAuthorization{}
			}
			effective.BinaryAuthorization.BreakglassJustification = actual.GetBinaryAuthorization().GetBreakglassJustification()
		}
	}

	if fieldSet["spec.client"] && desired.Spec.Client == nil {
		effective.Client = actual.Client
	}
	if fieldSet["spec.clientVersion"] && desired.Spec.ClientVersion == nil {
		effective.ClientVersion = actual.ClientVersion
	}
	if fieldSet["spec.launchStage"] && desired.Spec.LaunchStage == nil {
		effective.LaunchStage = actual.LaunchStage
	}

	// Handle nested fields within 'template'
	if desired.Spec.Template == nil {
		if strings.Contains(fieldsToDefault, "spec.template") {
			// If the entire template is unspecified in the spec, but a subfield is in the annotation,
			// we should default the whole template.
			effective.Template = actual.Template
		}
		return effective, nil
	}

	if fieldSet["spec.template.annotations"] && desired.Spec.Template.Annotations == nil {
		if actual.GetTemplate() != nil {
			if effective.Template == nil {
				effective.Template = &runpb.ExecutionTemplate{}
			}
			effective.Template.Annotations = actual.Template.Annotations
		}
	}
	if fieldSet["spec.template.parallelism"] && desired.Spec.Template.Parallelism == nil {
		if actual.GetTemplate() != nil {
			if effective.Template == nil {
				effective.Template = &runpb.ExecutionTemplate{}
			}
			effective.Template.Parallelism = actual.Template.Parallelism
		}
	}
	if fieldSet["spec.template.taskCount"] && desired.Spec.Template.TaskCount == nil {
		if actual.GetTemplate() != nil {
			if effective.Template == nil {
				effective.Template = &runpb.ExecutionTemplate{}
			}
			effective.Template.TaskCount = actual.Template.TaskCount
		}
	}

	// Handle nested fields within 'template.template'
	if desired.Spec.Template.Template == nil {
		if strings.Contains(fieldsToDefault, "spec.template.template") {
			if actual.GetTemplate() != nil {
				if effective.Template == nil {
					effective.Template = &runpb.ExecutionTemplate{}
				}
				effective.Template.Template = actual.Template.Template
			}
		}
		return effective, nil
	}

	if fieldSet["spec.template.template.containers"] && len(desired.Spec.Template.Template.Containers) == 0 {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.Containers = actual.Template.Template.Containers
		}
	} else if len(desired.Spec.Template.Template.Containers) > 0 {
		// User specified containers, so we need to merge sub-fields.
		actualContainers := make(map[string]*runpb.Container)
		if actual.GetTemplate() != nil && actual.GetTemplate().GetTemplate() != nil {
			for _, c := range actual.GetTemplate().GetTemplate().GetContainers() {
				actualContainers[c.GetName()] = c
			}
		}

		// We need to iterate over the desired KRM containers to check for nil fields,
		// and apply defaults to the corresponding effective proto containers.
		for i, desiredContainer := range desired.Spec.Template.Template.Containers {
			// This check is important to prevent panics if the effective proto has fewer containers
			// than the desired spec, which shouldn't happen with the current logic but is good for safety.
			if i >= len(effective.GetTemplate().GetTemplate().GetContainers()) {
				continue
			}
			effectiveContainer := effective.GetTemplate().GetTemplate().GetContainers()[i]
			actualContainer, ok := actualContainers[direct.ValueOf(desiredContainer.Name)]
			if !ok {
				continue
			}

			if fieldSet["spec.template.template.containers[].resources"] && desiredContainer.Resources == nil {
				effectiveContainer.Resources = actualContainer.Resources
			} else if desiredContainer.Resources != nil {
				if fieldSet["spec.template.template.containers[].resources.limits"] {
					if effectiveContainer.Resources == nil {
						effectiveContainer.Resources = &runpb.ResourceRequirements{}
					}
					if effectiveContainer.Resources.Limits == nil {
						effectiveContainer.Resources.Limits = make(map[string]string)
					}
					// Merge the maps: copy keys from actual if they don't exist in effective.
					for k, v := range actualContainer.GetResources().GetLimits() {
						if _, exists := effectiveContainer.Resources.Limits[k]; !exists {
							effectiveContainer.Resources.Limits[k] = v
						}
					}
				}
			}
			if fieldSet["spec.template.template.containers[].workingDir"] && desiredContainer.WorkingDir == nil {
				effectiveContainer.WorkingDir = actualContainer.WorkingDir
			}
			if fieldSet["spec.template.template.containers[].livenessProbe"] && desiredContainer.LivenessProbe == nil {
				effectiveContainer.LivenessProbe = actualContainer.LivenessProbe
			}
			if fieldSet["spec.template.template.containers[].startupProbe"] && desiredContainer.StartupProbe == nil {
				effectiveContainer.StartupProbe = actualContainer.StartupProbe
			}
		}
	}

	if fieldSet["spec.template.template.volumes"] && len(desired.Spec.Template.Template.Volumes) == 0 {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.Volumes = actual.Template.Template.Volumes
		}
	}
	if fieldSet["spec.template.template.maxRetries"] && desired.Spec.Template.Template.MaxRetries == nil {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.Retries = actual.Template.Template.Retries
		}
	}
	if fieldSet["spec.template.template.timeout"] && desired.Spec.Template.Template.Timeout == nil {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.Timeout = actual.Template.Template.Timeout
		}
	}
	if fieldSet["spec.template.template.serviceAccountRef"] && desired.Spec.Template.Template.ServiceAccountRef == nil {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.ServiceAccount = actual.Template.Template.ServiceAccount
		}
	}
	if fieldSet["spec.template.template.executionEnvironment"] && desired.Spec.Template.Template.ExecutionEnvironment == nil {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.ExecutionEnvironment = actual.Template.Template.ExecutionEnvironment
		}
	}
	if fieldSet["spec.template.template.encryptionKeyRef"] && desired.Spec.Template.Template.EncryptionKeyRef == nil {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.EncryptionKey = actual.Template.Template.EncryptionKey
		}
	}
	if fieldSet["spec.template.template.vpcAccess"] && desired.Spec.Template.Template.VPCAccess == nil {
		if actual.GetTemplate().GetTemplate() != nil {
			if effective.GetTemplate().GetTemplate() == nil {
				effective.Template.Template = &runpb.TaskTemplate{}
			}
			effective.Template.Template.VpcAccess = actual.Template.Template.VpcAccess
		}
	}

	return effective, nil
}
