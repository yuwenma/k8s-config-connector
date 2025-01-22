// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableTableGVK = GroupVersion.WithKind("BigtableTable")

type BigtableTableSpec struct {
	ResourceID  *string              `json:"resourceID,omitempty"`
	InstanceRef *BigtableInstanceRef `json:"instanceRef,omitempty"`

	ChangeStreamConfig *ChangeStreamConfig `json:"changeStreamConfig,omitempty"`

	ColumnFamilies map[string]ColumnFamily `json:"columnFamilies,omitempty"`
	Granularity    *string                 `json:"granularity,omitempty"`
}
type BigtableTableStatus struct {
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	ExternalRef *string `json:"externalRef,omitempty"`

	ObservedState *BigtableTableObservedState `json:"observedState,omitempty"`
}

type BigtableTableObservedState struct {
	ClusterStates map[string]TableClusterState `json:"clusterStates,omitempty"`
	RestoreInfo   *RestoreInfo                 `json:"restoreInfo,omitempty"`
}
type BigtableTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableTableSpec   `json:"spec,omitempty"`
	Status BigtableTableStatus `json:"status,omitempty"`
}
type BigtableTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableTable `json:"items"`
}

func () init() {
	SchemeBuilder.Register(&BigtableTable{}, &BigtableTableList{})
}

type ChangeStreamConfig struct {
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
}
type ColumnFamily struct {
	GcRule *GcRule `json:"gcRule,omitempty"`
}
type GcRule struct {
	MaxNumVersions *int32              `json:"maxNumVersions,omitempty"`
	MaxAge         *string             `json:"maxAge,omitempty"`
	Intersection   *GcRuleIntersection `json:"intersection,omitempty"`
	Union          *GcRuleUnion        `json:"union,omitempty"`
}
type GcRuleIntersection struct {
	Rules []GcRule `json:"rules,omitempty"`
}
type GcRuleUnion struct {
	Rules []GcRule `json:"rules,omitempty"`
}
type RestoreInfo struct {
	SourceType *string     `json:"sourceType,omitempty"`
	BackupInfo *BackupInfo `json:"backupInfo,omitempty"`
}
type BackupInfo struct {
	Backup      *string `json:"backup,omitempty"`
	StartTime   *string `json:"startTime,omitempty"`
	EndTime     *string `json:"endTime,omitempty"`
	SourceTable *string `json:"sourceTable,omitempty"`
}
type TableClusterState struct {
	ReplicationState *string          `json:"replicationState,omitempty"`
	EncryptionInfo   []EncryptionInfo `json:"encryptionInfo,omitempty"`
}
type EncryptionInfo struct {
	EncryptionType *string `json:"encryptionType,omitempty"`
	KmsKeyVersion  *string `json:"kmsKeyVersion,omitempty"`
}
