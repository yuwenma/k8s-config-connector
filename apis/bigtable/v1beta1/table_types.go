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
	ResourceID *string `json:"resourceID,omitempty"`

	ChangeStreamConfig *ChangeStreamConfig `json:"changeStreamConfig,omitempty"`

	ColumnFamilies map[string]ColumnFamily `json:"columnFamilies,omitempty"`

	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	Granularity *string              `json:"granularity,omitempty"`
	InstanceRef *BigtableInstanceRef `json:"instanceRef"`

	AutomatedBackupPolicy *TableAutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`
}

type BigtableTableStatus struct {
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	ExternalRef *string `json:"externalRef,omitempty"`

	ObservedState *BigtableTableObservedState `json:"observedState,omitempty"`
	ClusterStates map[string]ClusterState     `json:"clusterStates,omitempty"`
	RestoreInfo   *RestoreInfo                `json:"restoreInfo,omitempty"`
}

type BigtableTableObservedState struct {
	RestoreInfo *RestoreInfo `json:"restoreInfo,omitempty"`

	ClusterStates map[string]ClusterState `json:"clusterStates,omitempty"`
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
	RetentionPeriod *durationpb.Duration `json:"retentionPeriod,omitempty"`
}

type ColumnFamily struct {
	GcRule    *GcRule `json:"gcRule,omitempty"`
	ValueType *Type   `json:"valueType,omitempty"`
}

type ClusterState struct {
	ReplicationState *string          `json:"replicationState,omitempty"`
	EncryptionInfo   []EncryptionInfo `json:"encryptionInfo,omitempty"`
}

type EncryptionInfo struct {
	EncryptionType *string `json:"encryptionType,omitempty"`
	KmsKeyVersion  *string `json:"kmsKeyVersion,omitempty"`
}

type GcRule struct {
	MaxNumVersions *int32               `json:"maxNumVersions,omitempty"`
	MaxAge         *durationpb.Duration `json:"maxAge,omitempty"`
	Intersection   *GcRuleIntersection  `json:"intersection,omitempty"`
	Union          *GcRuleUnion         `json:"union,omitempty"`
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
	Backup       *string `json:"backup,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	EndTime      *string `json:"endTime,omitempty"`
	SourceTable  *string `json:"sourceTable,omitempty"`
	SourceBackup *string `json:"sourceBackup,omitempty"`
}

type TableAutomatedBackupPolicy struct {
	RetentionPeriod *durationpb.Duration `json:"retentionPeriod,omitempty"`
	Frequency       *durationpb.Duration `json:"frequency,omitempty"`
}

type Type struct {
	BytesType     *TypeBytes     `json:"bytesType,omitempty"`
	StringType    *TypeString    `json:"stringType,omitempty"`
	Int64Type     *TypeInt64     `json:"int64Type,omitempty"`
	Float32Type   *TypeFloat32   `json:"float32Type,omitempty"`
	Float64Type   *TypeFloat64   `json:"float64Type,omitempty"`
	BoolType      *TypeBool      `json:"boolType,omitempty"`
	TimestampType *TypeTimestamp `json:"timestampType,omitempty"`
	DateType      *TypeDate      `json:"dateType,omitempty"`
	AggregateType *TypeAggregate `json:"aggregateType,omitempty"`
	StructType    *TypeStruct    `json:"structType,omitempty"`
	ArrayType     *TypeArray     `json:"arrayType,omitempty"`
	MapType       *TypeMap       `json:"mapType,omitempty"`
}

type TypeAggregate struct {
	InputType        *Type                          `json:"inputType,omitempty"`
	StateType        *Type                          `json:"stateType,omitempty"`
	Sum              *TypeAggregateSum              `json:"sum,omitempty"`
	HllppUniqueCount *TypeAggregateHLLPPUniqueCount `json:"hllppUniqueCount,omitempty"`
	Max              *TypeAggregateMax              `json:"max,omitempty"`
	Min              *TypeAggregateMin              `json:"min,omitempty"`
}
type TypeAggregateHLLPPUniqueCount struct {
}
type TypeAggregateMax struct {
}
type TypeAggregateMin struct {
}
type TypeAggregateSum struct {
}

type TypeArray struct {
	ElementType *Type `json:"elementType,omitempty"`
}
type TypeBool struct {
}

type TypeBytes struct {
	Encoding *string `json:"encoding,omitempty"`
}
type TypeDate struct {
}
type TypeFloat32 struct {
}
type TypeFloat64 struct {
}

type TypeInt64 struct {
	Encoding *string `json:"encoding,omitempty"`
}

type TypeMap struct {
	KeyType   *Type `json:"keyType,omitempty"`
	ValueType *Type `json:"valueType,omitempty"`
}

type TypeString struct {
	Encoding *string `json:"encoding,omitempty"`
}

type TypeStruct struct {
	Fields []TypeStructField `json:"fields,omitempty"`
}

type TypeStructField struct {
	FieldName *string `json:"fieldName,omitempty"`
	Type      *Type   `json:"type,omitempty"`
}
type TypeTimestamp struct {
}
