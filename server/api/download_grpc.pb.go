// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: api/download.proto

package api

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

// FileDownloaderClient is the client API for FileDownloader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileDownloaderClient interface {
	Upload(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*Code, error)
	Download(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*Stream, error)
}

type fileDownloaderClient struct {
	cc grpc.ClientConnInterface
}

func NewFileDownloaderClient(cc grpc.ClientConnInterface) FileDownloaderClient {
	return &fileDownloaderClient{cc}
}

func (c *fileDownloaderClient) Upload(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*Code, error) {
	out := new(Code)
	err := c.cc.Invoke(ctx, "/api.FileDownloader/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileDownloaderClient) Download(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*Stream, error) {
	out := new(Stream)
	err := c.cc.Invoke(ctx, "/api.FileDownloader/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileDownloaderServer is the server API for FileDownloader service.
// All implementations must embed UnimplementedFileDownloaderServer
// for forward compatibility
type FileDownloaderServer interface {
	Upload(context.Context, *Stream) (*Code, error)
	Download(context.Context, *Stream) (*Stream, error)
	mustEmbedUnimplementedFileDownloaderServer()
}

// UnimplementedFileDownloaderServer must be embedded to have forward compatible implementations.
type UnimplementedFileDownloaderServer struct {
}

func (UnimplementedFileDownloaderServer) Upload(context.Context, *Stream) (*Code, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFileDownloaderServer) Download(context.Context, *Stream) (*Stream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedFileDownloaderServer) mustEmbedUnimplementedFileDownloaderServer() {}

// UnsafeFileDownloaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileDownloaderServer will
// result in compilation errors.
type UnsafeFileDownloaderServer interface {
	mustEmbedUnimplementedFileDownloaderServer()
}

func RegisterFileDownloaderServer(s grpc.ServiceRegistrar, srv FileDownloaderServer) {
	s.RegisterService(&FileDownloader_ServiceDesc, srv)
}

func _FileDownloader_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileDownloaderServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FileDownloader/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileDownloaderServer).Upload(ctx, req.(*Stream))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileDownloader_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileDownloaderServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FileDownloader/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileDownloaderServer).Download(ctx, req.(*Stream))
	}
	return interceptor(ctx, in, info, handler)
}

// FileDownloader_ServiceDesc is the grpc.ServiceDesc for FileDownloader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileDownloader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.FileDownloader",
	HandlerType: (*FileDownloaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _FileDownloader_Upload_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _FileDownloader_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/download.proto",
}
