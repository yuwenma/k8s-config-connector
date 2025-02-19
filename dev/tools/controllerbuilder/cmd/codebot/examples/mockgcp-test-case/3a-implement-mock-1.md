I'm trying to find the proto service and the proto message for `${GCLOUD_COMMAND} --project ${PROJECT}`. 

Remember this:
- CONTROLLER_BUILDER is `go run ../dev/tools/controllerbuilder/`
- If you need to run `gcloud` command, pass in flag `--project ${PROJECT}`
- If you find a `metadata.yaml` under the `mock*` directory, you can read the file to get proto information or other information. 


Now, let's find the proto service and proto message. Write read from `mock${SERVICE}/metadata.yaml` first. If you cannot find anything, we can find it via `${GCLOUD_COMMAND}`.

For example, the proto service for `gcloud firestore` is  `google.cloud.filestore.v1.CloudFilestoreManager`, and the proto message is `google.cloud.filestore.v1.Backup`, the directory is `mockfirestore/`.


Once you find the proto service and message, remember the value in the `mock${SERVICE}/metadata.yaml`, and move forward.


Now I want you to run the CONTROLLER_BUILDER command to implement the proto service in golang for `${GCLOUD_COMMAND}`, and the command gives you some go code and I want you to write the code under mock${SERVICE} appropriately.


Here is some examples.

For `gcloud filestore backups`, you should run
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

