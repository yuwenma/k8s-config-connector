// Copyright 2024 Google LLC
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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Resolver = &BigQueryConnectionConnectionRef{}

type ErrNoServerID struct{}

func (e *ErrNoServerID) Error() string {
	return "BigQueryConnectionConnection service-generated ID is not assigned yet"
}

type BigQueryConnectionConnectionRef struct {
	// A reference to an externally managed BigQueryConnectionConnection resource.
	// Should be in the format `projects/<projectID>/locations/<location>/connections/<connectionID>`.
	External string `json:"external,omitempty"`

	// The `name` of a `BigQueryConnectionConnection` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `BigQueryConnectionConnection` resource.
	Namespace string `json:"namespace,omitempty"`

	parent string
}

func (r *BigQueryConnectionConnectionRef) GetName() string {
	return r.Name
}

func (r *BigQueryConnectionConnectionRef) GetNamespace() string {
	return r.Namespace
}

func (r *BigQueryConnectionConnectionRef) GetExternal() string {
	return r.External
}

func (r *BigQueryConnectionConnectionRef) GVK() schema.GroupVersionKind {
	return BigQueryConnectionConnectionGVK
}

func (r *BigQueryConnectionConnectionRef) String() string {
	return r.External
}

func (r *BigQueryConnectionConnectionRef) Parent() string {
	return r.parent
}

func (r *BigQueryConnectionConnectionRef) ValidateExternal() error {
	if r.External == "" {
		return fmt.Errorf("external not specified")
	}
	r.External = strings.TrimPrefix(r.External, "/")
	tokens := strings.Split(r.External, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
		return fmt.Errorf("format of BigQueryConnectionConnection external=%q was not known (use projects/<projectId>/locations/<location>/connections/<connectionID>)", r.External)
	}
	return nil
}

func (r *BigQueryConnectionConnectionRef) Resolve(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) error {
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return err
	}
	location, _, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return fmt.Errorf("reading spec.location: %w", err)
	}

	// Get desired service-generated ID from spec
	desiredServiceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return fmt.Errorf("reading spec.resourceID: %w", err)
	}
	if desiredServiceID != "" {
		if _, err := uuid.Parse(desiredServiceID); err != nil {
			return fmt.Errorf("parse spec.resourceID as UUID")
		}
	}

	r.parent = "projects/" + projectID + "/locations/" + location

	// Service Generated ID does not exist
	if actualExternalRef == "" && desiredServiceID == "" {
		return &ErrNoServerID{}
	}

	if actualExternalRef == "" && desiredServiceID != "" {
		r.External = r.parent + "/connections/" + desiredServiceID
		return nil
	}
	// validate
	actualExternalRef = strings.TrimPrefix(actualExternalRef, "/")
	tokens := strings.Split(actualExternalRef, "/")
	if tokens[1] != projectID {
		return fmt.Errorf("spec.projectRef changed, expect %s, got %s", tokens[1], projectID)
	}
	if tokens[3] != location {
		return fmt.Errorf("spec.location changed, expect %s, got %s", tokens[3], location)
	}
	if desiredServiceID != "" && tokens[5] != desiredServiceID {
		// Service generated ID shall not be reset in the same BigQueryConnectionConnection.
		// TODO: what if multiple BigQueryConnectionConnection points to the same GCP Connection?
		return fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already acquired the Connection %s",
			desiredServiceID, tokens[5])
	}
	r.External = actualExternalRef
	return nil
}
