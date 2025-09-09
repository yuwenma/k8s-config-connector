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

package sql

import (
	"reflect"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func InstancesMatch(desired *api.DatabaseInstance, actual *api.DatabaseInstance, diff *structuredreporting.Diff, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.DatabaseVersion == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.DatabaseVersion != actual.DatabaseVersion {
		diff.AddField(".databaseVersion", actual.DatabaseVersion, desired.DatabaseVersion)
		return false
	}
	if ignoreUnspecified && desired.DiskEncryptionConfiguration == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !DiskEncryptionConfigurationsMatch(desired.DiskEncryptionConfiguration, actual.DiskEncryptionConfiguration, ignoreUnspecified) {
		diff.AddField(".diskEncryptionConfiguration", actual.DiskEncryptionConfiguration, desired.DiskEncryptionConfiguration)
		return false
	}
	// Ignore GeminiConfig. It is not supported in KRM API.
	if ignoreUnspecified && desired.InstanceType == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.InstanceType != actual.InstanceType {
		diff.AddField(".instanceType", actual.InstanceType, desired.InstanceType)
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.MaintenanceVersion == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.MaintenanceVersion != actual.MaintenanceVersion {
		diff.AddField(".maintenanceVersion", actual.MaintenanceVersion, desired.MaintenanceVersion)
		return false
	}
	if ignoreUnspecified && desired.MasterInstanceName == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.MasterInstanceName != actual.MasterInstanceName {
		diff.AddField(".masterInstanceName", actual.MasterInstanceName, desired.MasterInstanceName)
		return false
	}
	// Ignore MaxDiskSize. It is not supported in KRM API.
	if ignoreUnspecified && desired.Name == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Name != actual.Name {
		diff.AddField(".name", actual.Name, desired.Name)
		return false
	}
	// Ignore OnPremisesConfiguration. It is not supported in KRM API.
	if ignoreUnspecified && desired.Region == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Region != actual.Region {
		diff.AddField(".region", actual.Region, desired.Region)
		return false
	}
	if ignoreUnspecified && desired.ReplicaConfiguration == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !ReplicaConfigurationsMatch(desired.ReplicaConfiguration, actual.ReplicaConfiguration, ignoreUnspecified) {
		diff.AddField(".replicaConfiguration", actual.ReplicaConfiguration, desired.ReplicaConfiguration)
		return false
	}
	// Ignore ReplicationCluster. It is not supported in KRM API.
	// Ignore RootPassword. It is not exported.
	if ignoreUnspecified && desired.Settings == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !SettingsMatch(desired.Settings, actual.Settings, diff, ignoreUnspecified) {
		return false
	}
	// Ignore SqlNetworkArchitecture. It is not supported in KRM API.
	// Ignore SwitchTransactionLogsToCloudStorageEnabled. It is not supported in KRM API.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DiskEncryptionConfigurationsMatch(desired *api.DiskEncryptionConfiguration, actual *api.DiskEncryptionConfiguration, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.KmsKeyName == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.KmsKeyName != actual.KmsKeyName {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func ReplicaConfigurationsMatch(desired *api.ReplicaConfiguration, actual *api.ReplicaConfiguration, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	// Ignore CascadableReplica. It is not supported in KRM API.
	if ignoreUnspecified && desired.FailoverTarget == false {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.FailoverTarget != actual.FailoverTarget {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.MysqlReplicaConfiguration == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !MysqlReplicaConfigurationsMatch(desired.MysqlReplicaConfiguration, actual.MysqlReplicaConfiguration, ignoreUnspecified) {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func SettingsMatch(desired *api.Settings, actual *api.Settings, diff *structuredreporting.Diff, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		diff.AddField(".settings", actual, desired)
		return false
	}
	if ignoreUnspecified && desired.ActivationPolicy == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ActivationPolicy != actual.ActivationPolicy {
		diff.AddField(".settings.activationPolicy", actual.ActivationPolicy, desired.ActivationPolicy)
		return false
	}
	if ignoreUnspecified && desired.ActiveDirectoryConfig == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !ActiveDirectoryConfigsMatch(desired.ActiveDirectoryConfig, actual.ActiveDirectoryConfig, ignoreUnspecified) {
		diff.AddField(".settings.activeDirectoryConfig", actual.ActiveDirectoryConfig, desired.ActiveDirectoryConfig)
		return false
	}
	if ignoreUnspecified && desired.AdvancedMachineFeatures == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !AdvancedMachineFeaturesMatch(desired.AdvancedMachineFeatures, actual.AdvancedMachineFeatures, ignoreUnspecified) {
		diff.AddField(".settings.advancedMachineFeatures", actual.AdvancedMachineFeatures, desired.AdvancedMachineFeatures)
		return false
	}
	if ignoreUnspecified && desired.AuthorizedGaeApplications == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !slicesMatch(desired.AuthorizedGaeApplications, actual.AuthorizedGaeApplications) {
		diff.AddField(".settings.authorizedGaeApplications", actual.AuthorizedGaeApplications, desired.AuthorizedGaeApplications)
		return false
	}
	if ignoreUnspecified && desired.AvailabilityType == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.AvailabilityType != actual.AvailabilityType {
		diff.AddField(".settings.availabilityType", actual.AvailabilityType, desired.AvailabilityType)
		return false
	}
	if ignoreUnspecified && desired.BackupConfiguration == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !BackupConfigurationsMatch(desired.BackupConfiguration, actual.BackupConfiguration, ignoreUnspecified) {
		diff.AddField(".settings.backupConfiguration", actual.BackupConfiguration, desired.BackupConfiguration)
		return false
	}
	if ignoreUnspecified && desired.Collation == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Collation != actual.Collation {
		diff.AddField(".settings.collation", actual.Collation, desired.Collation)
		return false
	}
	if ignoreUnspecified && desired.ConnectorEnforcement == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ConnectorEnforcement != actual.ConnectorEnforcement {
		diff.AddField(".settings.connectorEnforcement", actual.ConnectorEnforcement, desired.ConnectorEnforcement)
		return false
	}
	// Ignore CrashSafeReplicationEnabled. It is only applicable to first-gen instances.
	if ignoreUnspecified && desired.DataCacheConfig == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !DataCacheConfigsMatch(desired.DataCacheConfig, actual.DataCacheConfig, ignoreUnspecified) {
		diff.AddField(".settings.dataCacheConfig", actual.DataCacheConfig, desired.DataCacheConfig)
		return false
	}
	if ignoreUnspecified && desired.DataDiskSizeGb == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.DataDiskSizeGb != actual.DataDiskSizeGb {
		diff.AddField(".settings.dataDiskSizeGb", actual.DataDiskSizeGb, desired.DataDiskSizeGb)
		return false
	}
	if ignoreUnspecified && desired.DataDiskType == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.DataDiskType != actual.DataDiskType {
		diff.AddField(".settings.dataDiskType", actual.DataDiskType, desired.DataDiskType)
		return false
	}
	if ignoreUnspecified && desired.DatabaseFlags == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !DatabaseFlagListsMatch(desired.DatabaseFlags, actual.DatabaseFlags, ignoreUnspecified) {
		diff.AddField(".settings.databaseFlags", actual.DatabaseFlags, desired.DatabaseFlags)
		return false
	}
	// Ignore DatabaseReplicationEnabled. It is not supported in KRM API.
	if ignoreUnspecified && !desired.DeletionProtectionEnabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.DeletionProtectionEnabled != actual.DeletionProtectionEnabled {
		diff.AddField(".settings.deletionProtectionEnabled", actual.DeletionProtectionEnabled, desired.DeletionProtectionEnabled)
		return false
	}
	if ignoreUnspecified && desired.DenyMaintenancePeriods == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !DenyMaintenancePeriodListsMatch(desired.DenyMaintenancePeriods, actual.DenyMaintenancePeriods, ignoreUnspecified) {
		diff.AddField(".settings.denyMaintenancePeriods", actual.DenyMaintenancePeriods, desired.DenyMaintenancePeriods)
		return false
	}
	if ignoreUnspecified && desired.Edition == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Edition != actual.Edition {
		diff.AddField(".settings.edition", actual.Edition, desired.Edition)
		return false
	}
	// Ignore EnableDataplexIntegration. It is not supported in KRM API.
	// Ignore EnableGoogleMlIntegration. It is not supported in KRM API.
	if ignoreUnspecified && desired.InsightsConfig == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !InsightsConfigsMatch(desired.InsightsConfig, actual.InsightsConfig, ignoreUnspecified) {
		diff.AddField(".settings.insightsConfig", actual.InsightsConfig, desired.InsightsConfig)
		return false
	}
	if ignoreUnspecified && desired.IpConfiguration == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !IpConfigurationsMatch(desired.IpConfiguration, actual.IpConfiguration, ignoreUnspecified) {
		diff.AddField(".settings.ipConfiguration", actual.IpConfiguration, desired.IpConfiguration)
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.LocationPreference == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !LocationPreferencesMatch(desired.LocationPreference, actual.LocationPreference, ignoreUnspecified) {
		diff.AddField(".settings.locationPreference", actual.LocationPreference, desired.LocationPreference)
		return false
	}
	if ignoreUnspecified && desired.MaintenanceWindow == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !MaintenanceWindowsMatch(desired.MaintenanceWindow, actual.MaintenanceWindow, ignoreUnspecified) {
		diff.AddField(".settings.maintenanceWindow", actual.MaintenanceWindow, desired.MaintenanceWindow)
		return false
	}
	if ignoreUnspecified && desired.PasswordValidationPolicy == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !PasswordValidationPoliciesMatch(desired.PasswordValidationPolicy, actual.PasswordValidationPolicy, ignoreUnspecified) {
		diff.AddField(".settings.passwordValidationPolicy", actual.PasswordValidationPolicy, desired.PasswordValidationPolicy)
		return false
	}
	if ignoreUnspecified && desired.PricingPlan == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.PricingPlan != actual.PricingPlan {
		diff.AddField(".settings.pricingPlan", actual.PricingPlan, desired.PricingPlan)
		return false
	}
	if ignoreUnspecified && desired.ReplicationType == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ReplicationType != actual.ReplicationType {
		diff.AddField(".settings.replicationType", actual.ReplicationType, desired.ReplicationType)
		return false
	}
	if ignoreUnspecified && desired.SettingsVersion == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.SettingsVersion != actual.SettingsVersion {
		diff.AddField(".settings.settingsVersion", actual.SettingsVersion, desired.SettingsVersion)
		return false
	}
	if ignoreUnspecified && desired.SqlServerAuditConfig == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !SqlServerAuditConfigsMatch(desired.SqlServerAuditConfig, actual.SqlServerAuditConfig, ignoreUnspecified) {
		diff.AddField(".settings.sqlServerAuditConfig", actual.SqlServerAuditConfig, desired.SqlServerAuditConfig)
		return false
	}
	if ignoreUnspecified && desired.StorageAutoResize == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !StorageAutoResizesMatch(desired.StorageAutoResize, actual.StorageAutoResize, ignoreUnspecified) {
		diff.AddField(".settings.storageAutoResize", actual.StorageAutoResize, desired.StorageAutoResize)
		return false
	}
	if ignoreUnspecified && desired.StorageAutoResizeLimit == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.StorageAutoResizeLimit != actual.StorageAutoResizeLimit {
		diff.AddField(".settings.storageAutoResizeLimit", actual.StorageAutoResizeLimit, desired.StorageAutoResizeLimit)
		return false
	}
	if ignoreUnspecified && desired.Tier == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Tier != actual.Tier {
		diff.AddField(".settings.tier", actual.Tier, desired.Tier)
		return false
	}
	if ignoreUnspecified && desired.TimeZone == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.TimeZone != actual.TimeZone {
		diff.AddField(".settings.timeZone", actual.TimeZone, desired.TimeZone)
		return false
	}
	if ignoreUnspecified && desired.UserLabels == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !reflect.DeepEqual(desired.UserLabels, actual.UserLabels) {
		diff.AddField(".settings.userLabels", actual.UserLabels, desired.UserLabels)
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

// slicesMatch checks if two slices are equal, matching with reflect.DeepEqual.
// As a special-case, the empty slice is treated the same as the nil slice
func slicesMatch[T any](desired []T, actual []T) bool {
	if len(desired) != len(actual) {
		return false
	}
	if len(desired) == 0 && len(actual) == 0 {
		return true
	}
	return reflect.DeepEqual(desired, actual)
}

func MysqlReplicaConfigurationsMatch(desired *api.MySqlReplicaConfiguration, actual *api.MySqlReplicaConfiguration, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.CaCertificate == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.CaCertificate != actual.CaCertificate {
		return false
	}
	if ignoreUnspecified && desired.ClientCertificate == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ClientCertificate != actual.ClientCertificate {
		return false
	}
	if ignoreUnspecified && desired.ClientKey == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ClientKey != actual.ClientKey {
		return false
	}
	if ignoreUnspecified && desired.ConnectRetryInterval == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ConnectRetryInterval != actual.ConnectRetryInterval {
		return false
	}
	if ignoreUnspecified && desired.DumpFilePath == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.DumpFilePath != actual.DumpFilePath {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.MasterHeartbeatPeriod == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.MasterHeartbeatPeriod != actual.MasterHeartbeatPeriod {
		return false
	}
	// Ignore Password. It is not exported.
	if ignoreUnspecified && desired.SslCipher == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.SslCipher != actual.SslCipher {
		return false
	}
	if ignoreUnspecified && desired.Username == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Username != actual.Username {
		return false
	}
	if ignoreUnspecified && !desired.VerifyServerCertificate {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.VerifyServerCertificate != actual.VerifyServerCertificate {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func ActiveDirectoryConfigsMatch(desired *api.SqlActiveDirectoryConfig, actual *api.SqlActiveDirectoryConfig, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.Domain == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Domain != actual.Domain {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func AdvancedMachineFeaturesMatch(desired *api.AdvancedMachineFeatures, actual *api.AdvancedMachineFeatures, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.ThreadsPerCore == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ThreadsPerCore != actual.ThreadsPerCore {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func BackupConfigurationsMatch(desired *api.BackupConfiguration, actual *api.BackupConfiguration, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.BackupRetentionSettings == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !BackupRetentionSettingsMatch(desired.BackupRetentionSettings, actual.BackupRetentionSettings, ignoreUnspecified) {
		return false
	}
	if ignoreUnspecified && !desired.BinaryLogEnabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.BinaryLogEnabled != actual.BinaryLogEnabled {
		return false
	}
	if ignoreUnspecified && !desired.Enabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Enabled != actual.Enabled {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.Location == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Location != actual.Location {
		return false
	}
	if ignoreUnspecified && !desired.PointInTimeRecoveryEnabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.PointInTimeRecoveryEnabled != actual.PointInTimeRecoveryEnabled {
		return false
	}
	// Ignore StartTime if it is not set. empty string is not a valid start time.
	if ignoreUnspecified && desired.StartTime == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.StartTime != "" && desired.StartTime != actual.StartTime {
		return false
	}
	// Ignore TransactionLogRetentionDays if it is not set. 0 is not a valid transaction log retention days.
	if ignoreUnspecified && desired.TransactionLogRetentionDays == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.TransactionLogRetentionDays != 0 && desired.TransactionLogRetentionDays != actual.TransactionLogRetentionDays {
		return false
	}

	// Ignore ReplicationLogArchivingEnabled. It is not supported in KRM API.
	// Ignore TransactionalLogStorageState. It is not supported in KRM API.
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}
func BackupRetentionSettingsMatch(desired *api.BackupRetentionSettings, actual *api.BackupRetentionSettings, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.RetainedBackups == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.RetainedBackups != actual.RetainedBackups {
		return false
	}
	if ignoreUnspecified && desired.RetentionUnit == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.RetentionUnit != actual.RetentionUnit {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DataCacheConfigsMatch(desired *api.DataCacheConfig, actual *api.DataCacheConfig, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && !desired.DataCacheEnabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.DataCacheEnabled != actual.DataCacheEnabled {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DatabaseFlagListsMatch(desired []*api.DatabaseFlags, actual []*api.DatabaseFlags, ignoreUnspecified bool) bool {
	if ignoreUnspecified && desired == nil {
		return true
	}
	if len(desired) != len(actual) {
		return false
	}
	for i := 0; i < len(desired); i++ {
		if !DatabaseFlagsMatch(desired[i], actual[i], ignoreUnspecified) {
			return false
		}
	}
	return true
}

func DatabaseFlagsMatch(desired *api.DatabaseFlags, actual *api.DatabaseFlags, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.Name == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Name != actual.Name {
		return false
	}
	if ignoreUnspecified && desired.Value == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Value != actual.Value {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func DenyMaintenancePeriodListsMatch(desired []*api.DenyMaintenancePeriod, actual []*api.DenyMaintenancePeriod, ignoreUnspecified bool) bool {
	if ignoreUnspecified && desired == nil {
		return true
	}
	if len(desired) != len(actual) {
		return false
	}
	for i := 0; i < len(desired); i++ {
		if !DenyMaintenancePeriodsMatch(desired[i], actual[i], ignoreUnspecified) {
			return false
		}
	}
	return true
}

func DenyMaintenancePeriodsMatch(desired *api.DenyMaintenancePeriod, actual *api.DenyMaintenancePeriod, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.EndDate == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.EndDate != actual.EndDate {
		return false
	}
	if ignoreUnspecified && desired.StartDate == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.StartDate != actual.StartDate {
		return false
	}
	if ignoreUnspecified && desired.Time == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Time != actual.Time {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func InsightsConfigsMatch(desired *api.InsightsConfig, actual *api.InsightsConfig, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && !desired.QueryInsightsEnabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.QueryInsightsEnabled != actual.QueryInsightsEnabled {
		return false
	}
	if ignoreUnspecified && desired.QueryPlansPerMinute == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.QueryPlansPerMinute != actual.QueryPlansPerMinute {
		return false
	}
	if ignoreUnspecified && desired.QueryStringLength == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.QueryStringLength != actual.QueryStringLength {
		return false
	}
	if ignoreUnspecified && !desired.RecordApplicationTags {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.RecordApplicationTags != actual.RecordApplicationTags {
		return false
	}
	if ignoreUnspecified && !desired.RecordClientAddress {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.RecordClientAddress != actual.RecordClientAddress {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func IpConfigurationsMatch(desired *api.IpConfiguration, actual *api.IpConfiguration, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.AllocatedIpRange == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.AllocatedIpRange != actual.AllocatedIpRange {
		return false
	}
	if ignoreUnspecified && desired.AuthorizedNetworks == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !AclEntryListsMatch(desired.AuthorizedNetworks, actual.AuthorizedNetworks, ignoreUnspecified) {
		return false
	}
	if ignoreUnspecified && !desired.EnablePrivatePathForGoogleCloudServices {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.EnablePrivatePathForGoogleCloudServices != actual.EnablePrivatePathForGoogleCloudServices {
		return false
	}
	if ignoreUnspecified && !desired.Ipv4Enabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Ipv4Enabled != actual.Ipv4Enabled {
		return false
	}
	if ignoreUnspecified && desired.PrivateNetwork == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.PrivateNetwork != actual.PrivateNetwork {
		return false
	}
	if ignoreUnspecified && desired.PscConfig == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !PscConfigsMatch(desired.PscConfig, actual.PscConfig, ignoreUnspecified) {
		return false
	}
	if ignoreUnspecified && !desired.RequireSsl {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.RequireSsl != actual.RequireSsl {
		return false
	}
	// Ignore ServerCaMode. It is not supported in KRM API.
	if ignoreUnspecified && desired.SslMode == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.SslMode != actual.SslMode {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

// AclEntriesByName implements sort.Interface for []*api.AclEntry based on the Name field.
type AclEntriesByName []*api.AclEntry

func (a AclEntriesByName) Len() int           { return len(a) }
func (a AclEntriesByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AclEntriesByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func AclEntryListsMatch(desired []*api.AclEntry, actual []*api.AclEntry, ignoreUnspecified bool) bool {
	if ignoreUnspecified && desired == nil {
		return true
	}
	if len(desired) != len(actual) {
		return false
	}
	// We mustiterate over the AclEntry lists in sorted order,
	// so that the comparison is deterministic.
	sort.Sort(AclEntriesByName(desired))
	sort.Sort(AclEntriesByName(actual))
	// Compare the AclEntry lists.
	for i := 0; i < len(desired); i++ {
		if !AclEntriesMatch(desired[i], actual[i], ignoreUnspecified) {
			return false
		}
	}
	return true
}

func AclEntriesMatch(desired *api.AclEntry, actual *api.AclEntry, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.ExpirationTime == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ExpirationTime != actual.ExpirationTime {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.Name == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Name != actual.Name {
		return false
	}
	if ignoreUnspecified && desired.Value == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Value != actual.Value {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func PscConfigsMatch(desired *api.PscConfig, actual *api.PscConfig, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.AllowedConsumerProjects == nil {
		// If the desired field is unspecified, ignore the diff.
	} else if !reflect.DeepEqual(desired.AllowedConsumerProjects, actual.AllowedConsumerProjects) {
		return false
	}
	if ignoreUnspecified && !desired.PscEnabled {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.PscEnabled != actual.PscEnabled {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func LocationPreferencesMatch(desired *api.LocationPreference, actual *api.LocationPreference, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.FollowGaeApplication == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.FollowGaeApplication != actual.FollowGaeApplication {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.SecondaryZone == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.SecondaryZone != actual.SecondaryZone {
		return false
	}
	if ignoreUnspecified && desired.Zone == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Zone != actual.Zone {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func MaintenanceWindowsMatch(desired *api.MaintenanceWindow, actual *api.MaintenanceWindow, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.Day == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Day != actual.Day {
		return false
	}
	if ignoreUnspecified && desired.Hour == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Hour != actual.Hour {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.UpdateTrack == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.UpdateTrack != actual.UpdateTrack {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func PasswordValidationPoliciesMatch(desired *api.PasswordValidationPolicy, actual *api.PasswordValidationPolicy, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.Complexity == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Complexity != actual.Complexity {
		return false
	}
	// Ignore DisallowCompromisedCredentials. It is not supported in KRM API.
	if ignoreUnspecified && !desired.DisallowUsernameSubstring {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.DisallowUsernameSubstring != actual.DisallowUsernameSubstring {
		return false
	}
	if ignoreUnspecified && !desired.EnablePasswordPolicy {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.EnablePasswordPolicy != actual.EnablePasswordPolicy {
		return false
	}
	if ignoreUnspecified && desired.MinLength == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.MinLength != actual.MinLength {
		return false
	}
	if ignoreUnspecified && desired.PasswordChangeInterval == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.PasswordChangeInterval != actual.PasswordChangeInterval {
		return false
	}
	if ignoreUnspecified && desired.ReuseInterval == 0 {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.ReuseInterval != actual.ReuseInterval {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func SqlServerAuditConfigsMatch(desired *api.SqlServerAuditConfig, actual *api.SqlServerAuditConfig, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if ignoreUnspecified && desired.Bucket == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.Bucket != actual.Bucket {
		return false
	}
	// Ignore Kind. It is sometimes not set in API responses.
	if ignoreUnspecified && desired.RetentionInterval == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.RetentionInterval != actual.RetentionInterval {
		return false
	}
	if ignoreUnspecified && desired.UploadInterval == "" {
		// If the desired field is unspecified, ignore the diff.
	} else if desired.UploadInterval != actual.UploadInterval {
		return false
	}
	// Ignore ForceSendFields. Assume it is set correctly in desired.
	// Ignore NullFields. Assume it is set correctly in desired.
	return true
}

func StorageAutoResizesMatch(desired *bool, actual *bool, ignoreUnspecified bool) bool {
	if desired == nil && actual == nil {
		return true
	}
	if ignoreUnspecified && desired == nil {
		// If the desired field is unspecified, ignore the diff.
		return true
	}
	if !PointersMatch(desired, actual) {
		return false
	}
	if *desired != *actual {
		return false
	}
	return true
}

func PointersMatch[T any](desired *T, actual *T) bool {
	if (desired == nil && actual != nil) || (desired != nil && actual == nil) {
		// Pointers are not matching if one is nil and the other is not nil.
		return false
	}
	// Otherwise, they match. Either both are nil, or both are not nil.
	return true
}