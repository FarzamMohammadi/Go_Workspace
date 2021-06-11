// Copyright 2021 Google LLC
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

package datacatalog_test

import (
	"context"

	datacatalog "cloud.google.com/go/datacatalog/apiv1"
	"google.golang.org/api/iterator"
	datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_SearchCatalog() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.SearchCatalogRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchCatalog(ctx, req)
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

func ExampleClient_CreateEntryGroup() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.CreateEntryGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateEntryGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetEntryGroup() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.GetEntryGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetEntryGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateEntryGroup() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.UpdateEntryGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateEntryGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteEntryGroup() {
	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.DeleteEntryGroupRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteEntryGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListEntryGroups() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.ListEntryGroupsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListEntryGroups(ctx, req)
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

func ExampleClient_CreateEntry() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.CreateEntryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateEntry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateEntry() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.UpdateEntryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateEntry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteEntry() {
	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.DeleteEntryRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteEntry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetEntry() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.GetEntryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetEntry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_LookupEntry() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.LookupEntryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.LookupEntry(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListEntries() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.ListEntriesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListEntries(ctx, req)
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

func ExampleClient_CreateTagTemplate() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.CreateTagTemplateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateTagTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetTagTemplate() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.GetTagTemplateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetTagTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateTagTemplate() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.UpdateTagTemplateRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateTagTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteTagTemplate() {
	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.DeleteTagTemplateRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTagTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_CreateTagTemplateField() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.CreateTagTemplateFieldRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateTagTemplateField(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateTagTemplateField() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.UpdateTagTemplateFieldRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateTagTemplateField(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RenameTagTemplateField() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.RenameTagTemplateFieldRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RenameTagTemplateField(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_RenameTagTemplateFieldEnumValue() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.RenameTagTemplateFieldEnumValueRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RenameTagTemplateFieldEnumValue(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteTagTemplateField() {
	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.DeleteTagTemplateFieldRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTagTemplateField(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_CreateTag() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.CreateTagRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateTag() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.UpdateTagRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteTag() {
	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.DeleteTagRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListTags() {
	// import datacatalogpb "google.golang.org/genproto/googleapis/cloud/datacatalog/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datacatalogpb.ListTagsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListTags(ctx, req)
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

func ExampleClient_SetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_TestIamPermissions() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := datacatalog.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
