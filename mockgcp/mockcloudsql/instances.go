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

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1"
)

func toSqlInstanceOperation(old *longrunning.Operation) *pb.Operation {
	return &pb.Operation{}
}

type InstanceV1 struct {
	*MockInstanceService
	pb.UnimplementedSqlInstancesServiceServer
}

func (s *InstanceV1) Get(ctx context.Context, req *pb.SqlInstancesGetRequest) (*pb.DatabaseInstance, error) {
	name, err := s.parseInstancename(req.Instance)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *InstanceV1) Insert(ctx context.Context, req *pb.SqlInstancesInsertRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.Project + "/instances/" + req.Body.Name
	name, err := s.parseInstancename(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.Body).(*pb.DatabaseInstance)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if o, err := s.operations.NewLRO(ctx); err != nil {
		return nil, err
	} else {
		return toSqlInstanceOperation(o), err
	}
}

func (s *InstanceV1) Update(ctx context.Context, req *pb.SqlInstancesUpdateRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.Project + "/instances/" + req.Instance
	name, err := s.parseInstancename(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// NOTE: The SQL instance UPDATE method does not support update_mast. We assume the UPDATE can update everything.
	obj = proto.Clone(req.Body).(*pb.DatabaseInstance)
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if o, err := s.operations.NewLRO(ctx); err != nil {
		return nil, err
	} else {
		return toSqlInstanceOperation(o), err
	}
}

func (s *InstanceV1) Delete(ctx context.Context, req *pb.SqlInstancesDeleteRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.Project + "/instances/" + req.Instance
	name, err := s.parseInstancename(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.DatabaseInstance{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	if o, err := s.operations.NewLRO(ctx); err != nil {
		return nil, err
	} else {
		return toSqlInstanceOperation(o), err
	}
}
