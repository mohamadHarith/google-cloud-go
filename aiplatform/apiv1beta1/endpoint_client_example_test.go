// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package aiplatform_test

import (
	"context"

	aiplatform "cloud.google.com/go/aiplatform/apiv1beta1"
	"google.golang.org/api/iterator"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1"
)

func ExampleNewEndpointClient() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleEndpointClient_CreateEndpoint() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.CreateEndpointRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1#CreateEndpointRequest.
	}
	op, err := c.CreateEndpoint(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEndpointClient_GetEndpoint() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.GetEndpointRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1#GetEndpointRequest.
	}
	resp, err := c.GetEndpoint(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEndpointClient_ListEndpoints() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.ListEndpointsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1#ListEndpointsRequest.
	}
	it := c.ListEndpoints(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleEndpointClient_UpdateEndpoint() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.UpdateEndpointRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1#UpdateEndpointRequest.
	}
	resp, err := c.UpdateEndpoint(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEndpointClient_DeleteEndpoint() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.DeleteEndpointRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1#DeleteEndpointRequest.
	}
	op, err := c.DeleteEndpoint(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleEndpointClient_DeployModel() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.DeployModelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1#DeployModelRequest.
	}
	op, err := c.DeployModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEndpointClient_UndeployModel() {
	ctx := context.Background()
	c, err := aiplatform.NewEndpointClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.UndeployModelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1#UndeployModelRequest.
	}
	op, err := c.UndeployModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
