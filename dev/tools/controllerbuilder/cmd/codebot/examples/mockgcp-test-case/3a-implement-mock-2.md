I'm trying to find the proto service and the proto message for `${GCLOUD_COMMAND} --project ${PROJECT}`. Now, read from `mock${SERVICE}/metadata.yaml` to get proto.service and proto.message, then use the proto.service and proto.message to replace the values in the following command, and run it. This will give you some go code. You should write the code to directory `mock${SERVICE}`.


Here is some examples.

For `gcloud filestore backups --project ${PROJECT}`, you should run
```
go run ${REPO_ROOT}/dev/tools/controllerbuilder/ prompt --src-dir ${REPO_ROOT} --proto-dir ${REPO_ROOT}/.build/third_party/googleapis/ <<EOF > mockfilestore/backup.go
// +tool:mockgcp-support
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Backup
EOF
```

This writes the go code to `mockfilestore/backup.go`.

Some hints:

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to change the code.
* You can use VerifyCode to see any remaining compilation errors as you are iterating towards a solution.

Please add the service to `mock_http_roundtrip.go`, you can look at `mockfirestore` in `mock_http_roundtrip.go` as example.

* Use the ReadFile command to read the contents of the file.
* Don't forget to import the package!

