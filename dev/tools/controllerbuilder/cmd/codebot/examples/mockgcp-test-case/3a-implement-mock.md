Remember this:
- CONTROLLER_BUILDER is `go run ../dev/tools/controllerbuilder/`
 `${GCLOUD_COMMAND} --project ${PROJECT}` 
- If you need to run `gcloud` command, pass in flag `--project ${PROJECT}`

I'm trying to find its proto service and the proto message for `${GCLOUD_COMMAND} --project ${PROJECT}`. Normally, the proto service is `google.cloud.${SERVICE}.v1.${RESOURCE}s`, and the proto message is `google.cloud.${SERVICE}.v1.${RESOURCE}` but this is not guarnateed.

For example, the proto service for `gcloud workflows` is  `google.cloud.workflows.v1.Workflows`, and the proto message is `google.cloud.workflows.v1.Workflow`. Or the proto service for `gcloud firestore backups` is  `google.cloud.filestore.v1.CloudFilestoreManager`, and the proto message is `google.cloud.filestore.v1.Backup`


Once you find the proto service and message, please remember the value them as ${PROTO_SERVICE} and ${PROTO_MESSAGE} in lower case, and also write the ${PROTO_SERVICE} and ${PROTO_MESSAGE} in yaml format to the corresponding `mock${SERVICE}/metadata.yaml`, the key name should be `proto.service` and `proto.message`.


If you get the ${PROTO_SERVICE} and ${PROTO_MESSAGE}, move forward.


Now I want you to run the CONTROLLER_BUILDER command to implement the proto service for ${PROTO_SERVICE}, and write the go code from the result to a go file named `mockservice`. The `mockservice` should be in a go file path: mock${SERVICE}/${RESOURCE}.go. Modify the go file name based on this rules:
- If ${SERVICE} and the last word after "." in ${PROTO_SERVICE} is different, use the last word in ${PROTO_SERVICE} after "." instead. Use lower case
- If ${RESOURCE} and the last word after "." in ${PROTO_MESSAGE} is different, use the last word in ${PROTO_MESSAGE} after "." instead. Use lower case.

Here is some examples.
# For filestore instance
```
go run ${REPO_ROOT}/dev/tools/controllerbuilder/ prompt --src-dir ${REPO_ROOT} --proto-dir ${REPO_ROOT}/.build/third_party/googleapis/ <<EOF > mockworkflows/workflow.go
// +tool:mockgcp-support
// proto.service: google.cloud.workflows.v1.Workflows
// proto.message: google.cloud.workflows.v1.Workflow
EOF
```

# For backups
```
go run ${REPO_ROOT}/dev/tools/controllerbuilder/ prompt --src-dir ${REPO_ROOT} --proto-dir ${REPO_ROOT}/.build/third_party/googleapis/ <<EOF > mockfilestore/backup.go
// +tool:mockgcp-support
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Backup
EOF
```

Some hints:

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to change the code.
* You can use VerifyCode to see any remaining compilation errors as you are iterating towards a solution.

Please add the services in the go file `mockservice` to `mock_http_roundtrip.go`. Store the `mockservice` to the same `mock${SERVICE}/metadata.yaml`

* Use the ReadFile command to read the contents of the file.
* Use the EditFile command to insert mock${PROTO_SERVICE} into the list of services.
* Don't forget to import the package!

