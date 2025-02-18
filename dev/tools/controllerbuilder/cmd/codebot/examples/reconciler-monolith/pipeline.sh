#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#!/bin/bash
set -x
set -e
set -o nounset

echo "KIND: $KIND, the config connector kind" # e.g. CloudbuildWorkerPool
echo "PROTO: $PROTO, the gcp proto name" # e.g. WorkerPool
echo "PACKAGE: $PACKAGE, the gcp proto package or service name" # e.g. google.devtools.cloudbuild.v1
echo "SERVICE: $SERVICE, the config connector api group name" # e.g. cloudbuild

REPO_ROOT="$(git rev-parse --show-toplevel)"
CB_DIR=dev/tools/controllerbuilder
UNCOMMIT_ROOT="/usr/local/google/home/yuwenma/go/src/github.com/GoogleCloudPlatform/k8s-config-connector"
INS_PATH="dev/tools/controllerbuilder/cmd/codebot/examples/reconciler-monolith/"

PASS_LOG="${UNCOMMIT_ROOT}/compiled-resource-success-17.log" 
FAIL_LOG="${UNCOMMIT_ROOT}/compiled-resource-fail-17.log" 
FEATURE_BRANCH_PREFIX="resource-"
BRANCH=$(git branch --show)
CODEBOT_LOG="${UNCOMMIT_ROOT}/compiled-resource-codebot-17-${KIND}.log" 

# # This script should be run under the mock feature branch.
# # Verify the git repo right resource
if [[ "$BRANCH" != ${FEATURE_BRANCH_PREFIX}* ]]; then
    echo "[ERROR] git HEAD is not in a feature branch"
    exit 1
fi 

RESOURCE=${BRANCH#"${FEATURE_BRANCH_PREFIX}"}
echo "LLM to add reconciler for resource ${RESOURCE}"


function git_commit() {
    git add apis/"${SERVICE}"/v1alpha1/
    git add pkg/controller/direct/register/register.go 
    git add pkg/controller/direct/"${SERVICE}"/
    git add config/crds/resources/
    git add go.mod
    git add go.sum
    git add go.work
    git commit -m "$1"
    echo "[${KIND}] git-commit $1"  >> "${PASS_LOG}"
}

function reset_branch() {
    # rm -f "${REPO_ROOT}"/pkg/controller/direct/"${SERVICE}"/"${PROTO,,}"_controller.go
    # echo "Deleted ${REPO_ROOT}/pkg/controller/direct/${SERVICE}/${PROTO,,}_controller.go" 
    git remote add upstream git@github.com:GoogleCloudPlatform/k8s-config-connector.git | echo "failed to add git remote upstream"
    git fetch upstream 
        
    git checkout upstream/master
    git branch -D "${BRANCH}"
    git checkout -b "${BRANCH}"
}

function push_to_remote() {
    git push -f origin "${BRANCH}":"${BRANCH}"
    result=$(gh pr create --reviewer yuwenma --base master --title "llm: compilable ${KIND} reconciler" --fill --head "${BRANCH}" 2>&1)
    echo "[${KIND}] created PR: ${result}"  >> "${PASS_LOG}"
    return 0
}


function run_worker() {
    attempts=1
    while [[ attempts -le 4 ]]; do  
        PACKAGE="${PACKAGE}" KIND="${KIND}" PROTO="${PROTO}" SERVICE="${SERVICE}" REPO_ROOT=${REPO_ROOT} "${INS_PATH}"/1-add-reconciler.sh
        if [[ $? -eq 0 ]]; then
            git_commit "add basic ${KIND} reconciler"
        fi 

        if go build -o /dev/null cmd/manager/main.go
        then
            echo "[${KIND}] passed after ${attempts} attempts"  >> "${PASS_LOG}"
            exit 0
        fi

        ins="${INS_PATH}"/2-llm-add-parents.txt
        msg=$(cat $ins)
        msg=$(eval "echo \"$msg\"")
        codebot --ui-type "bash" --proto-dir "${REPO_ROOT}"/.build/third_party/googleapis/  --project "${PROJECT}"  <<EOF 2>&1 | tee -a "${CODEBOT_LOG}" 
"${msg}" 
EOF
        if [[ $? -eq 0 ]]; then
           git_commit "[llm] add parent for ${KIND}"
        fi 


        codebot --ui-type "bash" --proto-dir "${REPO_ROOT}"/.build/third_party/googleapis/  --project "${PROJECT}" <<EOF 2>&1 | tee -a "${CODEBOT_LOG}"  
"${msg}" 
EOF
        if [[ $? -eq 0 ]]; then
           git_commit "[llm] add parent for ${KIND}"
        fi 

        ins="${INS_PATH}"/3-llm-gcp-imports.txt
        msg=$(cat $ins)
        msg=$(eval "echo \"$msg\"")
        codebot --ui-type "bash" --proto-dir "${REPO_ROOT}"/.build/third_party/googleapis/  --project "${PROJECT}" <<EOF 2>&1 | tee -a "${CODEBOT_LOG}" 
${msg}
EOF
        if [[ $? -eq 0 ]]; then
            git_commit "[llm] fix gcp imports for ${KIND}"
        fi 

        ins="${INS_PATH}"/4-llm-fix-gobuild.txt
        msg=$(cat $ins)
        msg=$(eval "echo \"$msg\"")
        codebot --ui-type "bash" --proto-dir "${REPO_ROOT}"/.build/third_party/googleapis/  --project "${PROJECT}" <<EOF 2>&1 | tee -a "${CODEBOT_LOG}" 
"${msg}"
EOF
        if [[ $? -eq 0 ]]; then
            git_commit "[llm] fix gobuilds for ${KIND}"
        fi 

    ((attempts++))
    done 
    return 1
}


# Call LLM codebot to generate reconciler code
# reset_branch

if run_worker 
then
    make fmt
    make ready-pr

    git_commit "make read-pr"
    push_to_remote
    echo "[${KIND}] ALL PASS"  >> "${PASS_LOG}"
else 
    echo "Unfortunately, this run failed" 
    exit 1
fi