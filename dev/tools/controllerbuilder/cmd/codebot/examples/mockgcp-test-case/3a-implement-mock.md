I'm trying to implement the proto resource google.spanner.admin.database.v1.Backup.

Run the following command
```
mkdir -p mockspanner

go run ../dev/tools/controllerbuilder/ prompt --src-dir .. --proto-dir ../.build/third_party/googleapis/ <<EOF > mockspanner/backup.go
// +tool:mockgcp-service
// http.host: spanner.googleapis.com
// proto.service: google.spanner.admin.database.v1
// proto.message: google.spanner.admin.database.v1.Backup
EOF
```
