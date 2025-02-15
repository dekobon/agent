// Copyright 2020-2023 Buf Technologies, Inc.
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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/alpha/registry/v1alpha1/display.proto

package registryv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// DisplayServiceName is the fully-qualified name of the DisplayService service.
	DisplayServiceName = "buf.alpha.registry.v1alpha1.DisplayService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DisplayServiceDisplayOrganizationElementsProcedure is the fully-qualified name of the
	// DisplayService's DisplayOrganizationElements RPC.
	DisplayServiceDisplayOrganizationElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayOrganizationElements"
	// DisplayServiceDisplayRepositoryElementsProcedure is the fully-qualified name of the
	// DisplayService's DisplayRepositoryElements RPC.
	DisplayServiceDisplayRepositoryElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayRepositoryElements"
	// DisplayServiceDisplayPluginElementsProcedure is the fully-qualified name of the DisplayService's
	// DisplayPluginElements RPC.
	DisplayServiceDisplayPluginElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayPluginElements"
	// DisplayServiceDisplayTemplateElementsProcedure is the fully-qualified name of the
	// DisplayService's DisplayTemplateElements RPC.
	DisplayServiceDisplayTemplateElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayTemplateElements"
	// DisplayServiceDisplayUserElementsProcedure is the fully-qualified name of the DisplayService's
	// DisplayUserElements RPC.
	DisplayServiceDisplayUserElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayUserElements"
	// DisplayServiceDisplayServerElementsProcedure is the fully-qualified name of the DisplayService's
	// DisplayServerElements RPC.
	DisplayServiceDisplayServerElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayServerElements"
	// DisplayServiceDisplayOwnerEntitledElementsProcedure is the fully-qualified name of the
	// DisplayService's DisplayOwnerEntitledElements RPC.
	DisplayServiceDisplayOwnerEntitledElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayOwnerEntitledElements"
	// DisplayServiceDisplayRepositoryEntitledElementsProcedure is the fully-qualified name of the
	// DisplayService's DisplayRepositoryEntitledElements RPC.
	DisplayServiceDisplayRepositoryEntitledElementsProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/DisplayRepositoryEntitledElements"
	// DisplayServiceListManageableRepositoryRolesProcedure is the fully-qualified name of the
	// DisplayService's ListManageableRepositoryRoles RPC.
	DisplayServiceListManageableRepositoryRolesProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/ListManageableRepositoryRoles"
	// DisplayServiceListManageableUserRepositoryRolesProcedure is the fully-qualified name of the
	// DisplayService's ListManageableUserRepositoryRoles RPC.
	DisplayServiceListManageableUserRepositoryRolesProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/ListManageableUserRepositoryRoles"
	// DisplayServiceListManageablePluginRolesProcedure is the fully-qualified name of the
	// DisplayService's ListManageablePluginRoles RPC.
	DisplayServiceListManageablePluginRolesProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/ListManageablePluginRoles"
	// DisplayServiceListManageableUserPluginRolesProcedure is the fully-qualified name of the
	// DisplayService's ListManageableUserPluginRoles RPC.
	DisplayServiceListManageableUserPluginRolesProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/ListManageableUserPluginRoles"
	// DisplayServiceListManageableTemplateRolesProcedure is the fully-qualified name of the
	// DisplayService's ListManageableTemplateRoles RPC.
	DisplayServiceListManageableTemplateRolesProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/ListManageableTemplateRoles"
	// DisplayServiceListManageableUserTemplateRolesProcedure is the fully-qualified name of the
	// DisplayService's ListManageableUserTemplateRoles RPC.
	DisplayServiceListManageableUserTemplateRolesProcedure = "/buf.alpha.registry.v1alpha1.DisplayService/ListManageableUserTemplateRoles"
)

// DisplayServiceClient is a client for the buf.alpha.registry.v1alpha1.DisplayService service.
type DisplayServiceClient interface {
	// DisplayOrganizationElements returns which organization elements should be displayed to the user.
	DisplayOrganizationElements(context.Context, *connect_go.Request[v1alpha1.DisplayOrganizationElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOrganizationElementsResponse], error)
	// DisplayRepositoryElements returns which repository elements should be displayed to the user.
	DisplayRepositoryElements(context.Context, *connect_go.Request[v1alpha1.DisplayRepositoryElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryElementsResponse], error)
	// DisplayPluginElements returns which plugin elements should be displayed to the user.
	//
	// Deprecated: do not use.
	DisplayPluginElements(context.Context, *connect_go.Request[v1alpha1.DisplayPluginElementsRequest]) (*connect_go.Response[v1alpha1.DisplayPluginElementsResponse], error)
	// DisplayTemplateElements returns which template elements should be displayed to the user.
	//
	// Deprecated: do not use.
	DisplayTemplateElements(context.Context, *connect_go.Request[v1alpha1.DisplayTemplateElementsRequest]) (*connect_go.Response[v1alpha1.DisplayTemplateElementsResponse], error)
	// DisplayUserElements returns which user elements should be displayed to the user.
	DisplayUserElements(context.Context, *connect_go.Request[v1alpha1.DisplayUserElementsRequest]) (*connect_go.Response[v1alpha1.DisplayUserElementsResponse], error)
	// DisplayServerElements returns which server elements should be displayed to the user.
	DisplayServerElements(context.Context, *connect_go.Request[v1alpha1.DisplayServerElementsRequest]) (*connect_go.Response[v1alpha1.DisplayServerElementsResponse], error)
	// DisplayOwnerEntitledElements returns which owner elements are entitled to be displayed to the user.
	DisplayOwnerEntitledElements(context.Context, *connect_go.Request[v1alpha1.DisplayOwnerEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOwnerEntitledElementsResponse], error)
	// DisplayRepositoryEntitledElements returns which repository elements are entitled to be displayed to the user.
	DisplayRepositoryEntitledElements(context.Context, *connect_go.Request[v1alpha1.DisplayRepositoryEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryEntitledElementsResponse], error)
	// ListManageableRepositoryRoles returns which roles should be displayed
	// to the user when they are managing contributors on the repository.
	ListManageableRepositoryRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableRepositoryRolesResponse], error)
	// ListManageableUserRepositoryRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the repository.
	ListManageableUserRepositoryRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserRepositoryRolesResponse], error)
	// ListManageablePluginRoles returns which roles should be displayed
	// to the user when they are managing contributors on the plugin.
	//
	// Deprecated: do not use.
	ListManageablePluginRoles(context.Context, *connect_go.Request[v1alpha1.ListManageablePluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageablePluginRolesResponse], error)
	// ListManageableUserPluginRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the plugin.
	//
	// Deprecated: do not use.
	ListManageableUserPluginRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserPluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserPluginRolesResponse], error)
	// ListManageableTemplateRoles returns which roles should be displayed
	// to the user when they are managing contributors on the template.
	//
	// Deprecated: do not use.
	ListManageableTemplateRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableTemplateRolesResponse], error)
	// ListManageableUserTemplateRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the template.
	//
	// Deprecated: do not use.
	ListManageableUserTemplateRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserTemplateRolesResponse], error)
}

// NewDisplayServiceClient constructs a client for the buf.alpha.registry.v1alpha1.DisplayService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDisplayServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DisplayServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &displayServiceClient{
		displayOrganizationElements: connect_go.NewClient[v1alpha1.DisplayOrganizationElementsRequest, v1alpha1.DisplayOrganizationElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayOrganizationElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		displayRepositoryElements: connect_go.NewClient[v1alpha1.DisplayRepositoryElementsRequest, v1alpha1.DisplayRepositoryElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayRepositoryElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		displayPluginElements: connect_go.NewClient[v1alpha1.DisplayPluginElementsRequest, v1alpha1.DisplayPluginElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayPluginElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		displayTemplateElements: connect_go.NewClient[v1alpha1.DisplayTemplateElementsRequest, v1alpha1.DisplayTemplateElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayTemplateElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		displayUserElements: connect_go.NewClient[v1alpha1.DisplayUserElementsRequest, v1alpha1.DisplayUserElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayUserElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		displayServerElements: connect_go.NewClient[v1alpha1.DisplayServerElementsRequest, v1alpha1.DisplayServerElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayServerElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		displayOwnerEntitledElements: connect_go.NewClient[v1alpha1.DisplayOwnerEntitledElementsRequest, v1alpha1.DisplayOwnerEntitledElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayOwnerEntitledElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		displayRepositoryEntitledElements: connect_go.NewClient[v1alpha1.DisplayRepositoryEntitledElementsRequest, v1alpha1.DisplayRepositoryEntitledElementsResponse](
			httpClient,
			baseURL+DisplayServiceDisplayRepositoryEntitledElementsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listManageableRepositoryRoles: connect_go.NewClient[v1alpha1.ListManageableRepositoryRolesRequest, v1alpha1.ListManageableRepositoryRolesResponse](
			httpClient,
			baseURL+DisplayServiceListManageableRepositoryRolesProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listManageableUserRepositoryRoles: connect_go.NewClient[v1alpha1.ListManageableUserRepositoryRolesRequest, v1alpha1.ListManageableUserRepositoryRolesResponse](
			httpClient,
			baseURL+DisplayServiceListManageableUserRepositoryRolesProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listManageablePluginRoles: connect_go.NewClient[v1alpha1.ListManageablePluginRolesRequest, v1alpha1.ListManageablePluginRolesResponse](
			httpClient,
			baseURL+DisplayServiceListManageablePluginRolesProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listManageableUserPluginRoles: connect_go.NewClient[v1alpha1.ListManageableUserPluginRolesRequest, v1alpha1.ListManageableUserPluginRolesResponse](
			httpClient,
			baseURL+DisplayServiceListManageableUserPluginRolesProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listManageableTemplateRoles: connect_go.NewClient[v1alpha1.ListManageableTemplateRolesRequest, v1alpha1.ListManageableTemplateRolesResponse](
			httpClient,
			baseURL+DisplayServiceListManageableTemplateRolesProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		listManageableUserTemplateRoles: connect_go.NewClient[v1alpha1.ListManageableUserTemplateRolesRequest, v1alpha1.ListManageableUserTemplateRolesResponse](
			httpClient,
			baseURL+DisplayServiceListManageableUserTemplateRolesProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// displayServiceClient implements DisplayServiceClient.
type displayServiceClient struct {
	displayOrganizationElements       *connect_go.Client[v1alpha1.DisplayOrganizationElementsRequest, v1alpha1.DisplayOrganizationElementsResponse]
	displayRepositoryElements         *connect_go.Client[v1alpha1.DisplayRepositoryElementsRequest, v1alpha1.DisplayRepositoryElementsResponse]
	displayPluginElements             *connect_go.Client[v1alpha1.DisplayPluginElementsRequest, v1alpha1.DisplayPluginElementsResponse]
	displayTemplateElements           *connect_go.Client[v1alpha1.DisplayTemplateElementsRequest, v1alpha1.DisplayTemplateElementsResponse]
	displayUserElements               *connect_go.Client[v1alpha1.DisplayUserElementsRequest, v1alpha1.DisplayUserElementsResponse]
	displayServerElements             *connect_go.Client[v1alpha1.DisplayServerElementsRequest, v1alpha1.DisplayServerElementsResponse]
	displayOwnerEntitledElements      *connect_go.Client[v1alpha1.DisplayOwnerEntitledElementsRequest, v1alpha1.DisplayOwnerEntitledElementsResponse]
	displayRepositoryEntitledElements *connect_go.Client[v1alpha1.DisplayRepositoryEntitledElementsRequest, v1alpha1.DisplayRepositoryEntitledElementsResponse]
	listManageableRepositoryRoles     *connect_go.Client[v1alpha1.ListManageableRepositoryRolesRequest, v1alpha1.ListManageableRepositoryRolesResponse]
	listManageableUserRepositoryRoles *connect_go.Client[v1alpha1.ListManageableUserRepositoryRolesRequest, v1alpha1.ListManageableUserRepositoryRolesResponse]
	listManageablePluginRoles         *connect_go.Client[v1alpha1.ListManageablePluginRolesRequest, v1alpha1.ListManageablePluginRolesResponse]
	listManageableUserPluginRoles     *connect_go.Client[v1alpha1.ListManageableUserPluginRolesRequest, v1alpha1.ListManageableUserPluginRolesResponse]
	listManageableTemplateRoles       *connect_go.Client[v1alpha1.ListManageableTemplateRolesRequest, v1alpha1.ListManageableTemplateRolesResponse]
	listManageableUserTemplateRoles   *connect_go.Client[v1alpha1.ListManageableUserTemplateRolesRequest, v1alpha1.ListManageableUserTemplateRolesResponse]
}

// DisplayOrganizationElements calls
// buf.alpha.registry.v1alpha1.DisplayService.DisplayOrganizationElements.
func (c *displayServiceClient) DisplayOrganizationElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayOrganizationElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOrganizationElementsResponse], error) {
	return c.displayOrganizationElements.CallUnary(ctx, req)
}

// DisplayRepositoryElements calls
// buf.alpha.registry.v1alpha1.DisplayService.DisplayRepositoryElements.
func (c *displayServiceClient) DisplayRepositoryElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayRepositoryElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryElementsResponse], error) {
	return c.displayRepositoryElements.CallUnary(ctx, req)
}

// DisplayPluginElements calls buf.alpha.registry.v1alpha1.DisplayService.DisplayPluginElements.
//
// Deprecated: do not use.
func (c *displayServiceClient) DisplayPluginElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayPluginElementsRequest]) (*connect_go.Response[v1alpha1.DisplayPluginElementsResponse], error) {
	return c.displayPluginElements.CallUnary(ctx, req)
}

// DisplayTemplateElements calls buf.alpha.registry.v1alpha1.DisplayService.DisplayTemplateElements.
//
// Deprecated: do not use.
func (c *displayServiceClient) DisplayTemplateElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayTemplateElementsRequest]) (*connect_go.Response[v1alpha1.DisplayTemplateElementsResponse], error) {
	return c.displayTemplateElements.CallUnary(ctx, req)
}

// DisplayUserElements calls buf.alpha.registry.v1alpha1.DisplayService.DisplayUserElements.
func (c *displayServiceClient) DisplayUserElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayUserElementsRequest]) (*connect_go.Response[v1alpha1.DisplayUserElementsResponse], error) {
	return c.displayUserElements.CallUnary(ctx, req)
}

// DisplayServerElements calls buf.alpha.registry.v1alpha1.DisplayService.DisplayServerElements.
func (c *displayServiceClient) DisplayServerElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayServerElementsRequest]) (*connect_go.Response[v1alpha1.DisplayServerElementsResponse], error) {
	return c.displayServerElements.CallUnary(ctx, req)
}

// DisplayOwnerEntitledElements calls
// buf.alpha.registry.v1alpha1.DisplayService.DisplayOwnerEntitledElements.
func (c *displayServiceClient) DisplayOwnerEntitledElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayOwnerEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOwnerEntitledElementsResponse], error) {
	return c.displayOwnerEntitledElements.CallUnary(ctx, req)
}

// DisplayRepositoryEntitledElements calls
// buf.alpha.registry.v1alpha1.DisplayService.DisplayRepositoryEntitledElements.
func (c *displayServiceClient) DisplayRepositoryEntitledElements(ctx context.Context, req *connect_go.Request[v1alpha1.DisplayRepositoryEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryEntitledElementsResponse], error) {
	return c.displayRepositoryEntitledElements.CallUnary(ctx, req)
}

// ListManageableRepositoryRoles calls
// buf.alpha.registry.v1alpha1.DisplayService.ListManageableRepositoryRoles.
func (c *displayServiceClient) ListManageableRepositoryRoles(ctx context.Context, req *connect_go.Request[v1alpha1.ListManageableRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableRepositoryRolesResponse], error) {
	return c.listManageableRepositoryRoles.CallUnary(ctx, req)
}

// ListManageableUserRepositoryRoles calls
// buf.alpha.registry.v1alpha1.DisplayService.ListManageableUserRepositoryRoles.
func (c *displayServiceClient) ListManageableUserRepositoryRoles(ctx context.Context, req *connect_go.Request[v1alpha1.ListManageableUserRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserRepositoryRolesResponse], error) {
	return c.listManageableUserRepositoryRoles.CallUnary(ctx, req)
}

// ListManageablePluginRoles calls
// buf.alpha.registry.v1alpha1.DisplayService.ListManageablePluginRoles.
//
// Deprecated: do not use.
func (c *displayServiceClient) ListManageablePluginRoles(ctx context.Context, req *connect_go.Request[v1alpha1.ListManageablePluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageablePluginRolesResponse], error) {
	return c.listManageablePluginRoles.CallUnary(ctx, req)
}

// ListManageableUserPluginRoles calls
// buf.alpha.registry.v1alpha1.DisplayService.ListManageableUserPluginRoles.
//
// Deprecated: do not use.
func (c *displayServiceClient) ListManageableUserPluginRoles(ctx context.Context, req *connect_go.Request[v1alpha1.ListManageableUserPluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserPluginRolesResponse], error) {
	return c.listManageableUserPluginRoles.CallUnary(ctx, req)
}

// ListManageableTemplateRoles calls
// buf.alpha.registry.v1alpha1.DisplayService.ListManageableTemplateRoles.
//
// Deprecated: do not use.
func (c *displayServiceClient) ListManageableTemplateRoles(ctx context.Context, req *connect_go.Request[v1alpha1.ListManageableTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableTemplateRolesResponse], error) {
	return c.listManageableTemplateRoles.CallUnary(ctx, req)
}

// ListManageableUserTemplateRoles calls
// buf.alpha.registry.v1alpha1.DisplayService.ListManageableUserTemplateRoles.
//
// Deprecated: do not use.
func (c *displayServiceClient) ListManageableUserTemplateRoles(ctx context.Context, req *connect_go.Request[v1alpha1.ListManageableUserTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserTemplateRolesResponse], error) {
	return c.listManageableUserTemplateRoles.CallUnary(ctx, req)
}

// DisplayServiceHandler is an implementation of the buf.alpha.registry.v1alpha1.DisplayService
// service.
type DisplayServiceHandler interface {
	// DisplayOrganizationElements returns which organization elements should be displayed to the user.
	DisplayOrganizationElements(context.Context, *connect_go.Request[v1alpha1.DisplayOrganizationElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOrganizationElementsResponse], error)
	// DisplayRepositoryElements returns which repository elements should be displayed to the user.
	DisplayRepositoryElements(context.Context, *connect_go.Request[v1alpha1.DisplayRepositoryElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryElementsResponse], error)
	// DisplayPluginElements returns which plugin elements should be displayed to the user.
	//
	// Deprecated: do not use.
	DisplayPluginElements(context.Context, *connect_go.Request[v1alpha1.DisplayPluginElementsRequest]) (*connect_go.Response[v1alpha1.DisplayPluginElementsResponse], error)
	// DisplayTemplateElements returns which template elements should be displayed to the user.
	//
	// Deprecated: do not use.
	DisplayTemplateElements(context.Context, *connect_go.Request[v1alpha1.DisplayTemplateElementsRequest]) (*connect_go.Response[v1alpha1.DisplayTemplateElementsResponse], error)
	// DisplayUserElements returns which user elements should be displayed to the user.
	DisplayUserElements(context.Context, *connect_go.Request[v1alpha1.DisplayUserElementsRequest]) (*connect_go.Response[v1alpha1.DisplayUserElementsResponse], error)
	// DisplayServerElements returns which server elements should be displayed to the user.
	DisplayServerElements(context.Context, *connect_go.Request[v1alpha1.DisplayServerElementsRequest]) (*connect_go.Response[v1alpha1.DisplayServerElementsResponse], error)
	// DisplayOwnerEntitledElements returns which owner elements are entitled to be displayed to the user.
	DisplayOwnerEntitledElements(context.Context, *connect_go.Request[v1alpha1.DisplayOwnerEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOwnerEntitledElementsResponse], error)
	// DisplayRepositoryEntitledElements returns which repository elements are entitled to be displayed to the user.
	DisplayRepositoryEntitledElements(context.Context, *connect_go.Request[v1alpha1.DisplayRepositoryEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryEntitledElementsResponse], error)
	// ListManageableRepositoryRoles returns which roles should be displayed
	// to the user when they are managing contributors on the repository.
	ListManageableRepositoryRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableRepositoryRolesResponse], error)
	// ListManageableUserRepositoryRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the repository.
	ListManageableUserRepositoryRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserRepositoryRolesResponse], error)
	// ListManageablePluginRoles returns which roles should be displayed
	// to the user when they are managing contributors on the plugin.
	//
	// Deprecated: do not use.
	ListManageablePluginRoles(context.Context, *connect_go.Request[v1alpha1.ListManageablePluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageablePluginRolesResponse], error)
	// ListManageableUserPluginRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the plugin.
	//
	// Deprecated: do not use.
	ListManageableUserPluginRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserPluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserPluginRolesResponse], error)
	// ListManageableTemplateRoles returns which roles should be displayed
	// to the user when they are managing contributors on the template.
	//
	// Deprecated: do not use.
	ListManageableTemplateRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableTemplateRolesResponse], error)
	// ListManageableUserTemplateRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the template.
	//
	// Deprecated: do not use.
	ListManageableUserTemplateRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserTemplateRolesResponse], error)
}

// NewDisplayServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDisplayServiceHandler(svc DisplayServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(DisplayServiceDisplayOrganizationElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayOrganizationElementsProcedure,
		svc.DisplayOrganizationElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceDisplayRepositoryElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayRepositoryElementsProcedure,
		svc.DisplayRepositoryElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceDisplayPluginElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayPluginElementsProcedure,
		svc.DisplayPluginElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceDisplayTemplateElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayTemplateElementsProcedure,
		svc.DisplayTemplateElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceDisplayUserElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayUserElementsProcedure,
		svc.DisplayUserElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceDisplayServerElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayServerElementsProcedure,
		svc.DisplayServerElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceDisplayOwnerEntitledElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayOwnerEntitledElementsProcedure,
		svc.DisplayOwnerEntitledElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceDisplayRepositoryEntitledElementsProcedure, connect_go.NewUnaryHandler(
		DisplayServiceDisplayRepositoryEntitledElementsProcedure,
		svc.DisplayRepositoryEntitledElements,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceListManageableRepositoryRolesProcedure, connect_go.NewUnaryHandler(
		DisplayServiceListManageableRepositoryRolesProcedure,
		svc.ListManageableRepositoryRoles,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceListManageableUserRepositoryRolesProcedure, connect_go.NewUnaryHandler(
		DisplayServiceListManageableUserRepositoryRolesProcedure,
		svc.ListManageableUserRepositoryRoles,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceListManageablePluginRolesProcedure, connect_go.NewUnaryHandler(
		DisplayServiceListManageablePluginRolesProcedure,
		svc.ListManageablePluginRoles,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceListManageableUserPluginRolesProcedure, connect_go.NewUnaryHandler(
		DisplayServiceListManageableUserPluginRolesProcedure,
		svc.ListManageableUserPluginRoles,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceListManageableTemplateRolesProcedure, connect_go.NewUnaryHandler(
		DisplayServiceListManageableTemplateRolesProcedure,
		svc.ListManageableTemplateRoles,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	mux.Handle(DisplayServiceListManageableUserTemplateRolesProcedure, connect_go.NewUnaryHandler(
		DisplayServiceListManageableUserTemplateRolesProcedure,
		svc.ListManageableUserTemplateRoles,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	))
	return "/buf.alpha.registry.v1alpha1.DisplayService/", mux
}

// UnimplementedDisplayServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDisplayServiceHandler struct{}

func (UnimplementedDisplayServiceHandler) DisplayOrganizationElements(context.Context, *connect_go.Request[v1alpha1.DisplayOrganizationElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOrganizationElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayOrganizationElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) DisplayRepositoryElements(context.Context, *connect_go.Request[v1alpha1.DisplayRepositoryElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayRepositoryElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) DisplayPluginElements(context.Context, *connect_go.Request[v1alpha1.DisplayPluginElementsRequest]) (*connect_go.Response[v1alpha1.DisplayPluginElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayPluginElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) DisplayTemplateElements(context.Context, *connect_go.Request[v1alpha1.DisplayTemplateElementsRequest]) (*connect_go.Response[v1alpha1.DisplayTemplateElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayTemplateElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) DisplayUserElements(context.Context, *connect_go.Request[v1alpha1.DisplayUserElementsRequest]) (*connect_go.Response[v1alpha1.DisplayUserElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayUserElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) DisplayServerElements(context.Context, *connect_go.Request[v1alpha1.DisplayServerElementsRequest]) (*connect_go.Response[v1alpha1.DisplayServerElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayServerElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) DisplayOwnerEntitledElements(context.Context, *connect_go.Request[v1alpha1.DisplayOwnerEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayOwnerEntitledElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayOwnerEntitledElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) DisplayRepositoryEntitledElements(context.Context, *connect_go.Request[v1alpha1.DisplayRepositoryEntitledElementsRequest]) (*connect_go.Response[v1alpha1.DisplayRepositoryEntitledElementsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.DisplayRepositoryEntitledElements is not implemented"))
}

func (UnimplementedDisplayServiceHandler) ListManageableRepositoryRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableRepositoryRolesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.ListManageableRepositoryRoles is not implemented"))
}

func (UnimplementedDisplayServiceHandler) ListManageableUserRepositoryRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserRepositoryRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserRepositoryRolesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.ListManageableUserRepositoryRoles is not implemented"))
}

func (UnimplementedDisplayServiceHandler) ListManageablePluginRoles(context.Context, *connect_go.Request[v1alpha1.ListManageablePluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageablePluginRolesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.ListManageablePluginRoles is not implemented"))
}

func (UnimplementedDisplayServiceHandler) ListManageableUserPluginRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserPluginRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserPluginRolesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.ListManageableUserPluginRoles is not implemented"))
}

func (UnimplementedDisplayServiceHandler) ListManageableTemplateRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableTemplateRolesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.ListManageableTemplateRoles is not implemented"))
}

func (UnimplementedDisplayServiceHandler) ListManageableUserTemplateRoles(context.Context, *connect_go.Request[v1alpha1.ListManageableUserTemplateRolesRequest]) (*connect_go.Response[v1alpha1.ListManageableUserTemplateRolesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("buf.alpha.registry.v1alpha1.DisplayService.ListManageableUserTemplateRoles is not implemented"))
}
