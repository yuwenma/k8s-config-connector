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

// +tool:fuzz-gen
// proto.message: google.cloud.workstations.v1.WorkstationCluster
// proto.message: google.cloud.workstations.v1.WorkstationConfig
// proto.message: google.cloud.workstations.v1.Workstation
// api.group: workstations.cnrm.cloud.google.com

package workstations

import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzerWithKind("WorkstationCluster", workstationclusterFuzzer())
	fuzztesting.RegisterKRMFuzzerWithKind("WorkstationConfig", workstationConfigFuzzer())
	fuzztesting.RegisterKRMFuzzerWithKind("Workstation", workstationFuzzer())
}

func workstationclusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.WorkstationCluster{},
		WorkstationClusterSpec_FromProto, WorkstationClusterSpec_ToProto,
		WorkstationClusterObservedState_FromProto, WorkstationClusterObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".degraded")
	f.UnimplementedFields.Insert(".conditions")
	f.UnimplementedFields.Insert(".private_cluster_config.cluster_hostname")
	f.UnimplementedFields.Insert(".private_cluster_config.service_attachment_uri")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".private_cluster_config")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".subnetwork")
	f.SpecFields.Insert(".network")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".control_plane_ip")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".uid")

	return f
}

func workstationConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.WorkstationConfig{},
		WorkstationConfigSpec_FromProto, WorkstationConfigSpec_ToProto,
		WorkstationConfigObservedState_FromProto, WorkstationConfigObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".conditions")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".idle_timeout")
	f.SpecFields.Insert(".running_timeout")
	f.SpecFields.Insert(".host")
	f.SpecFields.Insert(".persistent_directories")
	f.SpecFields.Insert(".container")
	f.SpecFields.Insert(".encryption_key")
	f.SpecFields.Insert(".readiness_checks")
	f.SpecFields.Insert(".replica_zones")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".host.gce_instance.pooled_instances")
	f.StatusFields.Insert(".degraded")

	return f
}

func workstationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Workstation{},
		WorkstationSpec_FromProto, WorkstationSpec_ToProto,
		WorkstationObservedState_FromProto, WorkstationObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".reconciling")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".start_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".host")

	return f
}
