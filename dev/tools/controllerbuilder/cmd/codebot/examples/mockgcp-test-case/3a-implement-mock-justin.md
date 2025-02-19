I'm trying to implement the proto service google.spanner.admin.database.v1.Backup.

Run the following command
```
mkdir -p mockspanner

go run ${REPO_ROOT}/dev/tools/controllerbuilder/ prompt --src-dir ${REPO_ROOT} --proto-dir ${REPO_ROOT}/.build/third_party/googleapis/ <<EOF > mockspanner/service.go
// +tool:mockgcp-service
// http.host: spanner.googleapis.com
// proto.service: google.spanner.admin.database.v1
EOF
```
