#!/bin/bash
# Copyright 2024 Google LLC
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


set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

APIS_DIR=${REPO_ROOT}/apis/
OUTPUT_MAPPER=${REPO_ROOT}/pkg/controller/direct/

# DataFlow
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1beta1 \
    --output-api ${APIS_DIR} \
    --kind DataflowFlexTemplateJob \
    --proto-resource FlexTemplateRuntimeEnvironment

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.dataflow.v1beta3 \
    --api-version dataflow.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# SecureSourceManagerInstance
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --kind SecureSourceManagerInstance \
    --proto-resource Instance

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.securesourcemanager.v1 \
    --api-version securesourcemanager.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# RedisCluster
go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1alpha1  \
    --output-api ${APIS_DIR} \
    --kind RedisCluster \
    --proto-resource Cluster

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.cloud.redis.cluster.v1 \
    --api-version redis.cnrm.cloud.google.com/v1alpha1  \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# Bigtable

go run . generate-types  \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1  \
    --output-api ${APIS_DIR} \
    --kind BigtableInstance \
    --proto-resource Instance

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service google.bigtable.admin.v2 \
    --api-version bigtable.cnrm.cloud.google.com/v1beta1  \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# NetworkConnectivity
go run . generate-types \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
    --output-api ${APIS_DIR} \
    --kind NetworkConnectivityServiceConnectionPolicy \
    --proto-resource ServiceConnecqionPolicy

go run . generate-mapper \
    --proto-source-path ../proto-to-mapper/build/googleapis.pb \
    --service mockgcp.cloud.networkconnectivity.v1 \
    --api-version networkconnectivity.cnrm.cloud.google.com/v1alpha1 \
    --api-go-package-path github.com/GoogleCloudPlatform/k8s-config-connector/apis \
    --output-dir ${OUTPUT_MAPPER} \
    --api-dir ${APIS_DIR}

# Fix up formatting
${REPO_ROOT}/dev/tasks/fix-gofmt
