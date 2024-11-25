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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)
func BackupInfo_FromProto(mapCtx *direct.MapContext, in *pb.BackupInfo) *krm.BackupInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupInfo{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.SourceTable = direct.LazyPtr(in.GetSourceTable())
	out.SourceBackup = direct.LazyPtr(in.GetSourceBackup())
	return out
}
func BackupInfo_ToProto(mapCtx *direct.MapContext, in *krm.BackupInfo) *pb.BackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupInfo{}
	out.Backup = direct.ValueOf(in.Backup)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.SourceTable = direct.ValueOf(in.SourceTable)
	out.SourceBackup = direct.ValueOf(in.SourceBackup)
	return out
}
func ChangeStreamConfig_FromProto(mapCtx *direct.MapContext, in *pb.ChangeStreamConfig) *krm.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &krm.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	return out
}
func ChangeStreamConfig_ToProto(mapCtx *direct.MapContext, in *krm.ChangeStreamConfig) *pb.ChangeStreamConfig {
	if in == nil {
		return nil
	}
	out := &pb.ChangeStreamConfig{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	return out
}
func ColumnFamily_FromProto(mapCtx *direct.MapContext, in *pb.ColumnFamily) *krm.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &krm.ColumnFamily{}
	out.GcRule = GcRule_FromProto(mapCtx, in.GetGcRule())
	out.ValueType = Type_FromProto(mapCtx, in.GetValueType())
	return out
}
func ColumnFamily_ToProto(mapCtx *direct.MapContext, in *krm.ColumnFamily) *pb.ColumnFamily {
	if in == nil {
		return nil
	}
	out := &pb.ColumnFamily{}
	out.GcRule = GcRule_ToProto(mapCtx, in.GcRule)
	out.ValueType = Type_ToProto(mapCtx, in.ValueType)
	return out
}
func EncryptionInfo_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) *krm.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionInfo{}
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: EncryptionStatus
	out.KmsKeyVersion = direct.LazyPtr(in.GetKmsKeyVersion())
	return out
}
func EncryptionInfo_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionInfo) *pb.EncryptionInfo {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionInfo{}
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionInfo_EncryptionType](mapCtx, in.EncryptionType)
	// MISSING: EncryptionStatus
	out.KmsKeyVersion = direct.ValueOf(in.KmsKeyVersion)
	return out
}
func GcRule_FromProto(mapCtx *direct.MapContext, in *pb.GcRule) *krm.GcRule {
	if in == nil {
		return nil
	}
	out := &krm.GcRule{}
	out.MaxNumVersions = direct.LazyPtr(in.GetMaxNumVersions())
	out.MaxAge = direct.StringDuration_FromProto(mapCtx, in.GetMaxAge())
	out.Intersection = GcRule_Intersection_FromProto(mapCtx, in.GetIntersection())
	out.Union = GcRule_Union_FromProto(mapCtx, in.GetUnion())
	return out
}
func GcRule_ToProto(mapCtx *direct.MapContext, in *krm.GcRule) *pb.GcRule {
	if in == nil {
		return nil
	}
	out := &pb.GcRule{}
	if oneof := GcRule_MaxNumVersions_ToProto(mapCtx, in.MaxNumVersions); oneof != nil {
		out.Rule = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.MaxAge); oneof != nil {
		out.Rule = &pb.GcRule_MaxAge{MaxAge: oneof}
	}
	if oneof := GcRule_Intersection_ToProto(mapCtx, in.Intersection); oneof != nil {
		out.Rule = &pb.GcRule_Intersection_{Intersection: oneof}
	}
	if oneof := GcRule_Union_ToProto(mapCtx, in.Union); oneof != nil {
		out.Rule = &pb.GcRule_Union_{Union: oneof}
	}
	return out
}
func GcRule_Intersection_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Intersection) *krm.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &krm.GcRule_Intersection{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Intersection_ToProto(mapCtx *direct.MapContext, in *krm.GcRule_Intersection) *pb.GcRule_Intersection {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Intersection{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func GcRule_Union_FromProto(mapCtx *direct.MapContext, in *pb.GcRule_Union) *krm.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &krm.GcRule_Union{}
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, GcRule_FromProto)
	return out
}
func GcRule_Union_ToProto(mapCtx *direct.MapContext, in *krm.GcRule_Union) *pb.GcRule_Union {
	if in == nil {
		return nil
	}
	out := &pb.GcRule_Union{}
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, GcRule_ToProto)
	return out
}
func RestoreInfo_FromProto(mapCtx *direct.MapContext, in *pb.RestoreInfo) *krm.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &krm.RestoreInfo{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.BackupInfo = BackupInfo_FromProto(mapCtx, in.GetBackupInfo())
	return out
}
func RestoreInfo_ToProto(mapCtx *direct.MapContext, in *krm.RestoreInfo) *pb.RestoreInfo {
	if in == nil {
		return nil
	}
	out := &pb.RestoreInfo{}
	out.SourceType = direct.Enum_ToProto[pb.RestoreSourceType](mapCtx, in.SourceType)
	if oneof := BackupInfo_ToProto(mapCtx, in.BackupInfo); oneof != nil {
		out.SourceInfo = &pb.RestoreInfo_BackupInfo{BackupInfo: oneof}
	}
	return out
}
func Table_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.Table {
	if in == nil {
		return nil
	}
	out := &krm.Table{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	out.Granularity = direct.Enum_FromProto(mapCtx, in.GetGranularity())
	out.RestoreInfo = RestoreInfo_FromProto(mapCtx, in.GetRestoreInfo())
	out.ChangeStreamConfig = ChangeStreamConfig_FromProto(mapCtx, in.GetChangeStreamConfig())
	out.DeletionProtection = direct.LazyPtr(in.GetDeletionProtection())
	out.AutomatedBackupPolicy = Table_AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())
	return out
}
func Table_ToProto(mapCtx *direct.MapContext, in *krm.Table) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ClusterStates
	// MISSING: ColumnFamilies
	out.Granularity = direct.Enum_ToProto[pb.Table_TimestampGranularity](mapCtx, in.Granularity)
	out.RestoreInfo = RestoreInfo_ToProto(mapCtx, in.RestoreInfo)
	out.ChangeStreamConfig = ChangeStreamConfig_ToProto(mapCtx, in.ChangeStreamConfig)
	out.DeletionProtection = direct.ValueOf(in.DeletionProtection)
	if oneof := Table_AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy); oneof != nil {
		out.AutomatedBackupConfig = &pb.Table_AutomatedBackupPolicy_{AutomatedBackupPolicy: oneof}
	}
	return out
}
func Table_AutomatedBackupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Table_AutomatedBackupPolicy) *krm.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRetentionPeriod())
	out.Frequency = direct.StringDuration_FromProto(mapCtx, in.GetFrequency())
	return out
}
func Table_AutomatedBackupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.Table_AutomatedBackupPolicy) *pb.Table_AutomatedBackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Table_AutomatedBackupPolicy{}
	out.RetentionPeriod = direct.StringDuration_ToProto(mapCtx, in.RetentionPeriod)
	out.Frequency = direct.StringDuration_ToProto(mapCtx, in.Frequency)
	return out
}
func Table_ClusterState_FromProto(mapCtx *direct.MapContext, in *pb.Table_ClusterState) *krm.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &krm.Table_ClusterState{}
	out.ReplicationState = direct.Enum_FromProto(mapCtx, in.GetReplicationState())
	out.EncryptionInfo = direct.Slice_FromProto(mapCtx, in.EncryptionInfo, EncryptionInfo_FromProto)
	return out
}
func Table_ClusterState_ToProto(mapCtx *direct.MapContext, in *krm.Table_ClusterState) *pb.Table_ClusterState {
	if in == nil {
		return nil
	}
	out := &pb.Table_ClusterState{}
	out.ReplicationState = direct.Enum_ToProto[pb.Table_ClusterState_ReplicationState](mapCtx, in.ReplicationState)
	out.EncryptionInfo = direct.Slice_ToProto(mapCtx, in.EncryptionInfo, EncryptionInfo_ToProto)
	return out
}
func Type_FromProto(mapCtx *direct.MapContext, in *pb.Type) *krm.Type {
	if in == nil {
		return nil
	}
	out := &krm.Type{}
	out.BytesType = Type_Bytes_FromProto(mapCtx, in.GetBytesType())
	out.StringType = Type_String_FromProto(mapCtx, in.GetStringType())
	out.Int64Type = Type_Int64_FromProto(mapCtx, in.GetInt64Type())
	out.Float32Type = Type_Float32_FromProto(mapCtx, in.GetFloat32Type())
	out.Float64Type = Type_Float64_FromProto(mapCtx, in.GetFloat64Type())
	out.BoolType = Type_Bool_FromProto(mapCtx, in.GetBoolType())
	out.TimestampType = Type_Timestamp_FromProto(mapCtx, in.GetTimestampType())
	out.DateType = Type_Date_FromProto(mapCtx, in.GetDateType())
	out.AggregateType = Type_Aggregate_FromProto(mapCtx, in.GetAggregateType())
	out.StructType = Type_Struct_FromProto(mapCtx, in.GetStructType())
	out.ArrayType = Type_Array_FromProto(mapCtx, in.GetArrayType())
	out.MapType = Type_Map_FromProto(mapCtx, in.GetMapType())
	return out
}
func Type_ToProto(mapCtx *direct.MapContext, in *krm.Type) *pb.Type {
	if in == nil {
		return nil
	}
	out := &pb.Type{}
	if oneof := Type_Bytes_ToProto(mapCtx, in.BytesType); oneof != nil {
		out.Kind = &pb.Type_BytesType{BytesType: oneof}
	}
	if oneof := Type_String_ToProto(mapCtx, in.StringType); oneof != nil {
		out.Kind = &pb.Type_StringType{StringType: oneof}
	}
	if oneof := Type_Int64_ToProto(mapCtx, in.Int64Type); oneof != nil {
		out.Kind = &pb.Type_Int64Type{Int64Type: oneof}
	}
	if oneof := Type_Float32_ToProto(mapCtx, in.Float32Type); oneof != nil {
		out.Kind = &pb.Type_Float32Type{Float32Type: oneof}
	}
	if oneof := Type_Float64_ToProto(mapCtx, in.Float64Type); oneof != nil {
		out.Kind = &pb.Type_Float64Type{Float64Type: oneof}
	}
	if oneof := Type_Bool_ToProto(mapCtx, in.BoolType); oneof != nil {
		out.Kind = &pb.Type_BoolType{BoolType: oneof}
	}
	if oneof := Type_Timestamp_ToProto(mapCtx, in.TimestampType); oneof != nil {
		out.Kind = &pb.Type_TimestampType{TimestampType: oneof}
	}
	if oneof := Type_Date_ToProto(mapCtx, in.DateType); oneof != nil {
		out.Kind = &pb.Type_DateType{DateType: oneof}
	}
	if oneof := Type_Aggregate_ToProto(mapCtx, in.AggregateType); oneof != nil {
		out.Kind = &pb.Type_AggregateType{AggregateType: oneof}
	}
	if oneof := Type_Struct_ToProto(mapCtx, in.StructType); oneof != nil {
		out.Kind = &pb.Type_StructType{StructType: oneof}
	}
	if oneof := Type_Array_ToProto(mapCtx, in.ArrayType); oneof != nil {
		out.Kind = &pb.Type_ArrayType{ArrayType: oneof}
	}
	if oneof := Type_Map_ToProto(mapCtx, in.MapType); oneof != nil {
		out.Kind = &pb.Type_MapType{MapType: oneof}
	}
	return out
}
func Type_Aggregate_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate) *krm.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate{}
	out.InputType = Type_FromProto(mapCtx, in.GetInputType())
	out.StateType = Type_FromProto(mapCtx, in.GetStateType())
	out.Sum = Type_Aggregate_Sum_FromProto(mapCtx, in.GetSum())
	out.HllppUniqueCount = Type_Aggregate_HyperLogLogPlusPlusUniqueCount_FromProto(mapCtx, in.GetHllppUniqueCount())
	out.Max = Type_Aggregate_Max_FromProto(mapCtx, in.GetMax())
	out.Min = Type_Aggregate_Min_FromProto(mapCtx, in.GetMin())
	return out
}
func Type_Aggregate_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate) *pb.Type_Aggregate {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate{}
	out.InputType = Type_ToProto(mapCtx, in.InputType)
	out.StateType = Type_ToProto(mapCtx, in.StateType)
	if oneof := Type_Aggregate_Sum_ToProto(mapCtx, in.Sum); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Sum_{Sum: oneof}
	}
	if oneof := Type_Aggregate_HyperLogLogPlusPlusUniqueCount_ToProto(mapCtx, in.HllppUniqueCount); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_HllppUniqueCount{HllppUniqueCount: oneof}
	}
	if oneof := Type_Aggregate_Max_ToProto(mapCtx, in.Max); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Max_{Max: oneof}
	}
	if oneof := Type_Aggregate_Min_ToProto(mapCtx, in.Min); oneof != nil {
		out.Aggregator = &pb.Type_Aggregate_Min_{Min: oneof}
	}
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *krm.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_HyperLogLogPlusPlusUniqueCount_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_HyperLogLogPlusPlusUniqueCount) *pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_HyperLogLogPlusPlusUniqueCount{}
	return out
}
func Type_Aggregate_Max_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Max) *krm.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Max_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_Max) *pb.Type_Aggregate_Max {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Max{}
	return out
}
func Type_Aggregate_Min_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Min) *krm.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Min_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_Min) *pb.Type_Aggregate_Min {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Min{}
	return out
}
func Type_Aggregate_Sum_FromProto(mapCtx *direct.MapContext, in *pb.Type_Aggregate_Sum) *krm.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &krm.Type_Aggregate_Sum{}
	return out
}
func Type_Aggregate_Sum_ToProto(mapCtx *direct.MapContext, in *krm.Type_Aggregate_Sum) *pb.Type_Aggregate_Sum {
	if in == nil {
		return nil
	}
	out := &pb.Type_Aggregate_Sum{}
	return out
}
func Type_Array_FromProto(mapCtx *direct.MapContext, in *pb.Type_Array) *krm.Type_Array {
	if in == nil {
		return nil
	}
	out := &krm.Type_Array{}
	out.ElementType = Type_FromProto(mapCtx, in.GetElementType())
	return out
}
func Type_Array_ToProto(mapCtx *direct.MapContext, in *krm.Type_Array) *pb.Type_Array {
	if in == nil {
		return nil
	}
	out := &pb.Type_Array{}
	out.ElementType = Type_ToProto(mapCtx, in.ElementType)
	return out
}
func Type_Bool_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bool) *krm.Type_Bool {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bool{}
	return out
}
func Type_Bool_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bool) *pb.Type_Bool {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bool{}
	return out
}
func Type_Bytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes) *krm.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Bytes_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes) *pb.Type_Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes{}
	out.Encoding = Type_Bytes_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Bytes_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding) *krm.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes_Encoding{}
	out.Raw = Type_Bytes_Encoding_Raw_FromProto(mapCtx, in.GetRaw())
	return out
}
func Type_Bytes_Encoding_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes_Encoding) *pb.Type_Bytes_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding{}
	if oneof := Type_Bytes_Encoding_Raw_ToProto(mapCtx, in.Raw); oneof != nil {
		out.Encoding = &pb.Type_Bytes_Encoding_Raw_{Raw: oneof}
	}
	return out
}
func Type_Bytes_Encoding_Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_Bytes_Encoding_Raw) *krm.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &krm.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Bytes_Encoding_Raw_ToProto(mapCtx *direct.MapContext, in *krm.Type_Bytes_Encoding_Raw) *pb.Type_Bytes_Encoding_Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_Bytes_Encoding_Raw{}
	return out
}
func Type_Date_FromProto(mapCtx *direct.MapContext, in *pb.Type_Date) *krm.Type_Date {
	if in == nil {
		return nil
	}
	out := &krm.Type_Date{}
	return out
}
func Type_Date_ToProto(mapCtx *direct.MapContext, in *krm.Type_Date) *pb.Type_Date {
	if in == nil {
		return nil
	}
	out := &pb.Type_Date{}
	return out
}
func Type_Float32_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float32) *krm.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &krm.Type_Float32{}
	return out
}
func Type_Float32_ToProto(mapCtx *direct.MapContext, in *krm.Type_Float32) *pb.Type_Float32 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float32{}
	return out
}
func Type_Float64_FromProto(mapCtx *direct.MapContext, in *pb.Type_Float64) *krm.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &krm.Type_Float64{}
	return out
}
func Type_Float64_ToProto(mapCtx *direct.MapContext, in *krm.Type_Float64) *pb.Type_Float64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Float64{}
	return out
}
func Type_Int64_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64) *krm.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_Int64_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64) *pb.Type_Int64 {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64{}
	out.Encoding = Type_Int64_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_Int64_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding) *krm.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64_Encoding{}
	out.BigEndianBytes = Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx, in.GetBigEndianBytes())
	return out
}
func Type_Int64_Encoding_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64_Encoding) *pb.Type_Int64_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding{}
	if oneof := Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx, in.BigEndianBytes); oneof != nil {
		out.Encoding = &pb.Type_Int64_Encoding_BigEndianBytes_{BigEndianBytes: oneof}
	}
	return out
}
func Type_Int64_Encoding_BigEndianBytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_Int64_Encoding_BigEndianBytes) *krm.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_FromProto(mapCtx, in.GetBytesType())
	return out
}
func Type_Int64_Encoding_BigEndianBytes_ToProto(mapCtx *direct.MapContext, in *krm.Type_Int64_Encoding_BigEndianBytes) *pb.Type_Int64_Encoding_BigEndianBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Int64_Encoding_BigEndianBytes{}
	out.BytesType = Type_Bytes_ToProto(mapCtx, in.BytesType)
	return out
}
func Type_Map_FromProto(mapCtx *direct.MapContext, in *pb.Type_Map) *krm.Type_Map {
	if in == nil {
		return nil
	}
	out := &krm.Type_Map{}
	out.KeyType = Type_FromProto(mapCtx, in.GetKeyType())
	out.ValueType = Type_FromProto(mapCtx, in.GetValueType())
	return out
}
func Type_Map_ToProto(mapCtx *direct.MapContext, in *krm.Type_Map) *pb.Type_Map {
	if in == nil {
		return nil
	}
	out := &pb.Type_Map{}
	out.KeyType = Type_ToProto(mapCtx, in.KeyType)
	out.ValueType = Type_ToProto(mapCtx, in.ValueType)
	return out
}
func Type_String_FromProto(mapCtx *direct.MapContext, in *pb.Type_String) *krm.Type_String {
	if in == nil {
		return nil
	}
	out := &krm.Type_String{}
	out.Encoding = Type_String_Encoding_FromProto(mapCtx, in.GetEncoding())
	return out
}
func Type_String_ToProto(mapCtx *direct.MapContext, in *krm.Type_String) *pb.Type_String {
	if in == nil {
		return nil
	}
	out := &pb.Type_String{}
	out.Encoding = Type_String_Encoding_ToProto(mapCtx, in.Encoding)
	return out
}
func Type_String_Encoding_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding) *krm.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding{}
	out.Utf8Raw = Type_String_Encoding_Utf8Raw_FromProto(mapCtx, in.GetUtf8Raw())
	out.Utf8Bytes = Type_String_Encoding_Utf8Bytes_FromProto(mapCtx, in.GetUtf8Bytes())
	return out
}
func Type_String_Encoding_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding) *pb.Type_String_Encoding {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding{}
	if oneof := Type_String_Encoding_Utf8Raw_ToProto(mapCtx, in.Utf8Raw); oneof != nil {
		out.Encoding = &pb.Type_String_Encoding_Utf8Raw_{Utf8Raw: oneof}
	}
	if oneof := Type_String_Encoding_Utf8Bytes_ToProto(mapCtx, in.Utf8Bytes); oneof != nil {
		out.Encoding = &pb.Type_String_Encoding_Utf8Bytes_{Utf8Bytes: oneof}
	}
	return out
}
func Type_String_Encoding_Utf8Bytes_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Bytes) *krm.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Bytes_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding_Utf8Bytes) *pb.Type_String_Encoding_Utf8Bytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Bytes{}
	return out
}
func Type_String_Encoding_Utf8Raw_FromProto(mapCtx *direct.MapContext, in *pb.Type_String_Encoding_Utf8Raw) *krm.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &krm.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_String_Encoding_Utf8Raw_ToProto(mapCtx *direct.MapContext, in *krm.Type_String_Encoding_Utf8Raw) *pb.Type_String_Encoding_Utf8Raw {
	if in == nil {
		return nil
	}
	out := &pb.Type_String_Encoding_Utf8Raw{}
	return out
}
func Type_Struct_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct) *krm.Type_Struct {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct{}
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Type_Struct_Field_FromProto)
	return out
}
func Type_Struct_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct) *pb.Type_Struct {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct{}
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Type_Struct_Field_ToProto)
	return out
}
func Type_Struct_Field_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Field) *krm.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct_Field{}
	out.FieldName = direct.LazyPtr(in.GetFieldName())
	out.Type = Type_FromProto(mapCtx, in.GetType())
	return out
}
func Type_Struct_Field_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct_Field) *pb.Type_Struct_Field {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Field{}
	out.FieldName = direct.ValueOf(in.FieldName)
	out.Type = Type_ToProto(mapCtx, in.Type)
	return out
}
func Type_Timestamp_FromProto(mapCtx *direct.MapContext, in *pb.Type_Timestamp) *krm.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &krm.Type_Timestamp{}
	return out
}
func Type_Timestamp_ToProto(mapCtx *direct.MapContext, in *krm.Type_Timestamp) *pb.Type_Timestamp {
	if in == nil {
		return nil
	}
	out := &pb.Type_Timestamp{}
	return out
}
