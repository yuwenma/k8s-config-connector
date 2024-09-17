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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Resolver define the required methods to resolve the external identifier of a ConfigConnector object.
type Resolver interface {
	// Resolve assigns value to the "External" field, based on the Unstructured ConfigConnector object.
	// In general, `status.externalRef` shall be used if present, otherwise it shall be built from the `spec` parent fields
	// and metadata.name (or `spec.resourceID` if applicable).
	Resolve(context.Context, client.Reader, *unstructured.Unstructured) error

	// Validate the value of the "External" field
	ValidateExternal() error
	// Return the External value
	GetExternal() string
	// Return the Name value
	GetName() string
	// Return the Namespace value
	GetNamespace() string
	// Return the GroupVersionKind of the ConfigConnector resource.
	GVK() schema.GroupVersionKind
}

func NormalizeExternal(ctx context.Context, reader client.Reader, r Resolver) (string, error) {
	if r.GetExternal() != "" && r.GetName() != "" {
		return "", fmt.Errorf("cannot specify both name and external on BigQueryConnectionConnection reference")
	}
	if r.GetExternal() != "" {
		if err := r.ValidateExternal(); err != nil {
			return "", err
		}
		return r.GetExternal(), nil
	}
	key := types.NamespacedName{Name: r.GetName(), Namespace: r.GetNamespace()}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.GVK())
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", r.GVK(), key, err)
	}

	err := r.Resolve(ctx, reader, u)
	if err != nil {
		return "", fmt.Errorf("resolve externalRef from %s %s: %w", r.GVK(), key, err)
	}
	return r.GetExternal(), nil
}
