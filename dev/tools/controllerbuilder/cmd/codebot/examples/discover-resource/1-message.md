
export REPO_ROOT="$(git rev-parse --show-toplevel)"

# Get all proto from google client lib. Stored to 
PROTO_LIST=$REPO_ROOT/proto-list-$(date '+%m-%d').yaml
go run dev/tools/controllerbuilder/main.go iterate-types --output-file $REPO_ROOT/proto-list-$(date '+%m-%d').yaml

# Get all gcloud resource 
GCLOUD_LIST=gcloud-list.txt
mockgcp/dev/tools/print-gcloud-resource-commands >> $GCLOUD_LIST

# TargetResource
RESOURCE=DatastreamStream

# Start codebot
go run $REPO_ROOT/dev/tools/controllerbuilder/cmd/codebot/main.go 

>>> Remember that `RESOURCE=DatastreamStream`, `GCLOUD_LIST=gcloud-list.txt`
>>> Find a single line from $GCLOUD_LIST that represents resource $RESOURCE. Try reading from the `gcloud RESOURCE describe` to get more information about $$RESOURCE, and write everything you can find to $RESOURCE_info.txt   






>>> read file `gcloud-list.txt` and `proto-list.yaml`, try to map each line of `gcloud-list.txt` to a `service` `package` `proto` `kind` `protopath` section in the `proto-list.yaml`, store all you found in a new file `pair-result.yaml`

