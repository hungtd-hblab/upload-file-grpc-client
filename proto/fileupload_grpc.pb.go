// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: fileupload.proto

package fileuploadv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileUploadServiceClient is the client API for FileUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileUploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileUploadService_UploadClient, error)
}

type fileUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileUploadServiceClient(cc grpc.ClientConnInterface) FileUploadServiceClient {
	return &fileUploadServiceClient{cc}
}

func (c *fileUploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileUploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileUploadService_ServiceDesc.Streams[0], "/fileupload.v1.FileUploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadServiceUploadClient{stream}
	return x, nil
}

type FileUploadService_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type fileUploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *fileUploadServiceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileUploadServiceUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileUploadServiceServer is the server API for FileUploadService service.
// All implementations must embed UnimplementedFileUploadServiceServer
// for forward compatibility
type FileUploadServiceServer interface {
	Upload(FileUploadService_UploadServer) error
	mustEmbedUnimplementedFileUploadServiceServer()
}

// UnimplementedFileUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileUploadServiceServer struct {
}

func (UnimplementedFileUploadServiceServer) Upload(FileUploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileUploadServiceServer) mustEmbedUnimplementedFileUploadServiceServer() {}

// UnsafeFileUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileUploadServiceServer will
// result in compilation errors.
type UnsafeFileUploadServiceServer interface {
	mustEmbedUnimplementedFileUploadServiceServer()
}

func RegisterFileUploadServiceServer(s grpc.ServiceRegistrar, srv FileUploadServiceServer) {
	s.RegisterService(&FileUploadService_ServiceDesc, srv)
}

func _FileUploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileUploadServiceServer).Upload(&fileUploadServiceUploadServer{stream})
}

type FileUploadService_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type fileUploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *fileUploadServiceUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileUploadServiceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileUploadService_ServiceDesc is the grpc.ServiceDesc for FileUploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileUploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fileupload.v1.FileUploadService",
	HandlerType: (*FileUploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileUploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "fileupload.proto",
}