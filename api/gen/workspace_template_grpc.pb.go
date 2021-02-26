// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WorkspaceTemplateServiceClient is the client API for WorkspaceTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceTemplateServiceClient interface {
	// Get the generated WorkflowTemplate for a WorkspaceTemplate
	GenerateWorkspaceTemplateWorkflowTemplate(ctx context.Context, in *GenerateWorkspaceTemplateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	// Creates a WorkspaceTemplate
	CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error)
	// Updates a WorkspaceTemplate
	UpdateWorkspaceTemplate(ctx context.Context, in *UpdateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error)
	// Archives a WorkspaceTemplate
	ArchiveWorkspaceTemplate(ctx context.Context, in *ArchiveWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error)
	// Get a WorkspaceTemplate
	GetWorkspaceTemplate(ctx context.Context, in *GetWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error)
	ListWorkspaceTemplates(ctx context.Context, in *ListWorkspaceTemplatesRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplatesResponse, error)
	ListWorkspaceTemplateVersions(ctx context.Context, in *ListWorkspaceTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplateVersionsResponse, error)
	ListWorkspaceTemplatesField(ctx context.Context, in *ListWorkspaceTemplatesFieldRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplatesFieldResponse, error)
}

type workspaceTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceTemplateServiceClient(cc grpc.ClientConnInterface) WorkspaceTemplateServiceClient {
	return &workspaceTemplateServiceClient{cc}
}

func (c *workspaceTemplateServiceClient) GenerateWorkspaceTemplateWorkflowTemplate(ctx context.Context, in *GenerateWorkspaceTemplateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/GenerateWorkspaceTemplateWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error) {
	out := new(WorkspaceTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/CreateWorkspaceTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) UpdateWorkspaceTemplate(ctx context.Context, in *UpdateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error) {
	out := new(WorkspaceTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/UpdateWorkspaceTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) ArchiveWorkspaceTemplate(ctx context.Context, in *ArchiveWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error) {
	out := new(WorkspaceTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/ArchiveWorkspaceTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) GetWorkspaceTemplate(ctx context.Context, in *GetWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error) {
	out := new(WorkspaceTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/GetWorkspaceTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) ListWorkspaceTemplates(ctx context.Context, in *ListWorkspaceTemplatesRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplatesResponse, error) {
	out := new(ListWorkspaceTemplatesResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/ListWorkspaceTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) ListWorkspaceTemplateVersions(ctx context.Context, in *ListWorkspaceTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplateVersionsResponse, error) {
	out := new(ListWorkspaceTemplateVersionsResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/ListWorkspaceTemplateVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceTemplateServiceClient) ListWorkspaceTemplatesField(ctx context.Context, in *ListWorkspaceTemplatesFieldRequest, opts ...grpc.CallOption) (*ListWorkspaceTemplatesFieldResponse, error) {
	out := new(ListWorkspaceTemplatesFieldResponse)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/ListWorkspaceTemplatesField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceTemplateServiceServer is the server API for WorkspaceTemplateService service.
// All implementations must embed UnimplementedWorkspaceTemplateServiceServer
// for forward compatibility
type WorkspaceTemplateServiceServer interface {
	// Get the generated WorkflowTemplate for a WorkspaceTemplate
	GenerateWorkspaceTemplateWorkflowTemplate(context.Context, *GenerateWorkspaceTemplateWorkflowTemplateRequest) (*WorkflowTemplate, error)
	// Creates a WorkspaceTemplate
	CreateWorkspaceTemplate(context.Context, *CreateWorkspaceTemplateRequest) (*WorkspaceTemplate, error)
	// Updates a WorkspaceTemplate
	UpdateWorkspaceTemplate(context.Context, *UpdateWorkspaceTemplateRequest) (*WorkspaceTemplate, error)
	// Archives a WorkspaceTemplate
	ArchiveWorkspaceTemplate(context.Context, *ArchiveWorkspaceTemplateRequest) (*WorkspaceTemplate, error)
	// Get a WorkspaceTemplate
	GetWorkspaceTemplate(context.Context, *GetWorkspaceTemplateRequest) (*WorkspaceTemplate, error)
	ListWorkspaceTemplates(context.Context, *ListWorkspaceTemplatesRequest) (*ListWorkspaceTemplatesResponse, error)
	ListWorkspaceTemplateVersions(context.Context, *ListWorkspaceTemplateVersionsRequest) (*ListWorkspaceTemplateVersionsResponse, error)
	ListWorkspaceTemplatesField(context.Context, *ListWorkspaceTemplatesFieldRequest) (*ListWorkspaceTemplatesFieldResponse, error)
	mustEmbedUnimplementedWorkspaceTemplateServiceServer()
}

// UnimplementedWorkspaceTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceTemplateServiceServer struct {
}

func (UnimplementedWorkspaceTemplateServiceServer) GenerateWorkspaceTemplateWorkflowTemplate(context.Context, *GenerateWorkspaceTemplateWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWorkspaceTemplateWorkflowTemplate not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) CreateWorkspaceTemplate(context.Context, *CreateWorkspaceTemplateRequest) (*WorkspaceTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspaceTemplate not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) UpdateWorkspaceTemplate(context.Context, *UpdateWorkspaceTemplateRequest) (*WorkspaceTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspaceTemplate not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) ArchiveWorkspaceTemplate(context.Context, *ArchiveWorkspaceTemplateRequest) (*WorkspaceTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveWorkspaceTemplate not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) GetWorkspaceTemplate(context.Context, *GetWorkspaceTemplateRequest) (*WorkspaceTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceTemplate not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) ListWorkspaceTemplates(context.Context, *ListWorkspaceTemplatesRequest) (*ListWorkspaceTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceTemplates not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) ListWorkspaceTemplateVersions(context.Context, *ListWorkspaceTemplateVersionsRequest) (*ListWorkspaceTemplateVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceTemplateVersions not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) ListWorkspaceTemplatesField(context.Context, *ListWorkspaceTemplatesFieldRequest) (*ListWorkspaceTemplatesFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceTemplatesField not implemented")
}
func (UnimplementedWorkspaceTemplateServiceServer) mustEmbedUnimplementedWorkspaceTemplateServiceServer() {
}

// UnsafeWorkspaceTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceTemplateServiceServer will
// result in compilation errors.
type UnsafeWorkspaceTemplateServiceServer interface {
	mustEmbedUnimplementedWorkspaceTemplateServiceServer()
}

func RegisterWorkspaceTemplateServiceServer(s grpc.ServiceRegistrar, srv WorkspaceTemplateServiceServer) {
	s.RegisterService(&_WorkspaceTemplateService_serviceDesc, srv)
}

func _WorkspaceTemplateService_GenerateWorkspaceTemplateWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWorkspaceTemplateWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).GenerateWorkspaceTemplateWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/GenerateWorkspaceTemplateWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).GenerateWorkspaceTemplateWorkflowTemplate(ctx, req.(*GenerateWorkspaceTemplateWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_CreateWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).CreateWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/CreateWorkspaceTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).CreateWorkspaceTemplate(ctx, req.(*CreateWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_UpdateWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).UpdateWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/UpdateWorkspaceTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).UpdateWorkspaceTemplate(ctx, req.(*UpdateWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_ArchiveWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).ArchiveWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/ArchiveWorkspaceTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).ArchiveWorkspaceTemplate(ctx, req.(*ArchiveWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_GetWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).GetWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/GetWorkspaceTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).GetWorkspaceTemplate(ctx, req.(*GetWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_ListWorkspaceTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/ListWorkspaceTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplates(ctx, req.(*ListWorkspaceTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_ListWorkspaceTemplateVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceTemplateVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplateVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/ListWorkspaceTemplateVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplateVersions(ctx, req.(*ListWorkspaceTemplateVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceTemplateService_ListWorkspaceTemplatesField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceTemplatesFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplatesField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/ListWorkspaceTemplatesField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).ListWorkspaceTemplatesField(ctx, req.(*ListWorkspaceTemplatesFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkspaceTemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkspaceTemplateService",
	HandlerType: (*WorkspaceTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateWorkspaceTemplateWorkflowTemplate",
			Handler:    _WorkspaceTemplateService_GenerateWorkspaceTemplateWorkflowTemplate_Handler,
		},
		{
			MethodName: "CreateWorkspaceTemplate",
			Handler:    _WorkspaceTemplateService_CreateWorkspaceTemplate_Handler,
		},
		{
			MethodName: "UpdateWorkspaceTemplate",
			Handler:    _WorkspaceTemplateService_UpdateWorkspaceTemplate_Handler,
		},
		{
			MethodName: "ArchiveWorkspaceTemplate",
			Handler:    _WorkspaceTemplateService_ArchiveWorkspaceTemplate_Handler,
		},
		{
			MethodName: "GetWorkspaceTemplate",
			Handler:    _WorkspaceTemplateService_GetWorkspaceTemplate_Handler,
		},
		{
			MethodName: "ListWorkspaceTemplates",
			Handler:    _WorkspaceTemplateService_ListWorkspaceTemplates_Handler,
		},
		{
			MethodName: "ListWorkspaceTemplateVersions",
			Handler:    _WorkspaceTemplateService_ListWorkspaceTemplateVersions_Handler,
		},
		{
			MethodName: "ListWorkspaceTemplatesField",
			Handler:    _WorkspaceTemplateService_ListWorkspaceTemplatesField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace_template.proto",
}
