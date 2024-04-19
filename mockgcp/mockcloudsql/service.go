// Copyright 2022 Google LLC
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

package mockcloudsql

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked privateca service.
type MockInstanceService struct {
	*common.MockEnvironment
	storage    storage.Storage
	operations *operations.Operations
	v1         *InstanceV1
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockInstanceService {
	s := &MockInstanceService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1 = &InstanceV1{MockInstanceService: s}
	return s
}

func (s *MockInstanceService) ExpectedHost() string {
	return "sqladmin.googleapis.com"
}

func (s *MockInstanceService) Register(grpcServer *grpc.Server) {
	pb.RegisterSqlInstancesServiceServer(grpcServer, s.v1)
}

func (s *MockInstanceService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux := runtime.NewServeMux()

	if err := pb.RegisterSqlInstancesServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}

	return mux, nil
}
