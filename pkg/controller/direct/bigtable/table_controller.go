// Copyright 2024 Google LLC
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

package bigtable

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "google.golang.org/api/bigtableadmin/v2"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableTableGVK, NewTableModel)
}

func NewTableModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTable{config: *config}, nil
}

var _ directbase.Model = &modelTable{}

type modelTable struct {
	config config.ControllerConfig
}

func (m *modelTable) client(ctx context.Context) (*gcp.Service, error) {
	client, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return client.newAdminClient(ctx)
}

func (m *modelTable) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigtableTable{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTableIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get bigtable GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TableAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTable) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type TableAdapter struct {
	id        *krm.TableIdentity
	gcpClient *gcp.Service
	desired   *krm.BigtableTable
	actual    *gcp.Table
}

var _ directbase.Adapter = &TableAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TableAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Table", "name", a.id)

	table, err := a.gcpClient.Projects.Instances.Tables.Get(a.id.Parent().String()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = table
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TableAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Table", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := BigtableTableSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &gcp.CreateTableRequest{
		Table: resource,
	}
	created, err := a.gcpClient.Projects.Instances.Tables.Create(a.id.Parent().String(), req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("Table %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Table", "name", a.id)

	status := &krm.BigtableTableStatus{}
	status.ObservedState = BigtableTableObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TableAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Table", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := BigtableTableSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	// Option 1: This option is good for proto that has `field_mask` for output-only, immutable, required/optional.
	// TODO(contributor): If choosing this option, remove the "Option 2" code.
	{
		var err error
		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
	}

	// Option 2: manually add all mutable fields.
	// TODO(contributor): If choosing this option, remove the "Option 1" code.
	{
		if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
			paths = append(paths, "display_name")
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		status := &krm.BigtableTableStatus{}
		status.ObservedState = BigtableTableObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	op, err := a.gcpClient.Projects.Instances.Tables.Patch(a.id.String(), desiredPb).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating Table %s: %w", a.id, err)
	}
	if !op.Done {
		return nil
	}
	updated, err := a.gcpClient.Projects.Instances.Tables.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating Table %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Table", "name", a.id.String())

	status := &krm.BigtableTableStatus{}
	status.ObservedState = BigtableTableObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TableAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TableAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Table", "name", a.id)

	_, err := a.gcpClient.Projects.Instances.Tables.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("deleting Table %s: %w", a.id, err)
	}
	return true, nil
}
