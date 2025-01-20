package benchmark

type Resource struct {
	Kind    string
	RawFile string
	Proto   string
}

// Ground Truth resources
// TODO: We should exclude these ground truth from the gemnimi input datapoints.
var GroundTruthResources = []Resource{
	{
		Kind:    "BigQueryConnectionConnection",
		RawFile: "bigqueryconnection/v1beta1/connection_types.go",
		Proto:   "google.cloud.bigquery.connection.v1.Connection",
	},
	{
		Kind:    "KMSAutokeyConfig",
		RawFile: "kms/v1beta1/autokeyconfig_types.go",
		Proto:   "google.cloud.kms.v1.AutokeyConfig",
	},
	{
		Kind:    "BigQueryDataTransferConfig",
		RawFile: "bigquerydatatransfer/v1beta1/transferconfig_types.go",
		Proto:   "google.cloud.bigquery.datatransfer.v1.TransferConfig",
	},
	{
		Kind:    "CloudBuildWorkerPool",
		RawFile: "cloudbuild/v1beta1/workerpool_types.go",
		Proto:   "google.devtools.cloudbuild.v1.WorkerPool",
	},
	{
		Kind:    "DataformRepository",
		RawFile: "dataform/v1beta1/repository_types.go",
		Proto:   "google.cloud.dataform.v1beta1.Repository",
	},
}
