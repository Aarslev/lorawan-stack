// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/tti/tenantbillingserver.proto

package ttipb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/tti/tenantbillingserver.proto", fileDescriptor_a10320d6120e9a0d)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/tti/tenantbillingserver.proto", fileDescriptor_a10320d6120e9a0d)
}

var fileDescriptor_a10320d6120e9a0d = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4f, 0x72, 0x31,
	0x14, 0xc5, 0x7b, 0xbf, 0x2f, 0x61, 0xc0, 0xe8, 0x80, 0x89, 0x03, 0x92, 0x3b, 0xe0, 0x6c, 0x9b,
	0xc8, 0x7f, 0x60, 0x74, 0x74, 0x51, 0x27, 0xb6, 0x3e, 0x52, 0x1f, 0x0d, 0xcf, 0xf6, 0xa5, 0xef,
	0x0a, 0x71, 0x23, 0x4e, 0x8c, 0x26, 0x2e, 0x8e, 0x2e, 0x26, 0x8c, 0x8c, 0x8c, 0x8c, 0x8c, 0x24,
	0x2e, 0x8c, 0xbc, 0xd6, 0x81, 0x91, 0x91, 0xd1, 0xf0, 0xc0, 0x41, 0xe3, 0xd8, 0x9c, 0x73, 0xfa,
	0xbb, 0xf7, 0xdc, 0xb2, 0x48, 0xac, 0x93, 0x3d, 0x69, 0x4e, 0x33, 0x92, 0xad, 0x8e, 0x90, 0xa9,
	0x16, 0x44, 0x5a, 0x90, 0x32, 0xd2, 0x50, 0xa4, 0x93, 0x44, 0x9b, 0x38, 0x53, 0xae, 0xab, 0x1c,
	0x4f, 0x9d, 0x25, 0x5b, 0x39, 0x20, 0xd2, 0x7c, 0x17, 0xe2, 0xdd, 0x46, 0xb5, 0x16, 0x5b, 0x1b,
	0x27, 0xaa, 0x48, 0x4a, 0x63, 0x2c, 0x49, 0xd2, 0xd6, 0x64, 0x5b, 0x77, 0xf5, 0x78, 0xa7, 0x16,
	0xaf, 0xe8, 0xe1, 0x4e, 0xa8, 0xfb, 0x94, 0x1e, 0x77, 0xe2, 0xc9, 0xdf, 0xec, 0x44, 0xb7, 0x94,
	0xc9, 0xd4, 0xd6, 0x74, 0xd6, 0x2c, 0xff, 0xbf, 0x8d, 0xb2, 0xca, 0x4d, 0xb9, 0x74, 0xad, 0x52,
	0xeb, 0xa8, 0x52, 0xe3, 0x3f, 0x27, 0xe0, 0x57, 0x8a, 0x94, 0xd3, 0x26, 0xbe, 0x90, 0x24, 0xab,
	0x47, 0x7c, 0x4b, 0xe4, 0xdf, 0x44, 0x7e, 0xb9, 0x21, 0xd6, 0x0f, 0x9f, 0x3e, 0x3e, 0x5f, 0xfe,
	0xed, 0xd7, 0xf7, 0x04, 0x45, 0x99, 0x70, 0xc5, 0x57, 0xe7, 0xef, 0x30, 0xcd, 0x11, 0x66, 0x39,
	0xc2, 0x3c, 0x47, 0xb6, 0xc8, 0x91, 0x2d, 0x73, 0x64, 0xab, 0x1c, 0xd9, 0x3a, 0x47, 0xe8, 0x7b,
	0x84, 0x81, 0x47, 0x36, 0xf4, 0x08, 0x23, 0x8f, 0x6c, 0xec, 0x91, 0x4d, 0x3c, 0xb2, 0xa9, 0x47,
	0x98, 0x79, 0x84, 0xb9, 0x47, 0xb6, 0xf0, 0x08, 0x4b, 0x8f, 0x6c, 0xe5, 0x11, 0xd6, 0x1e, 0x59,
	0x3f, 0x20, 0x1b, 0x04, 0x84, 0xe7, 0x80, 0xec, 0x35, 0x20, 0xbc, 0x05, 0x64, 0xc3, 0x80, 0x6c,
	0x14, 0x10, 0xc6, 0x01, 0x61, 0x12, 0x10, 0x9a, 0x22, 0xb6, 0x9c, 0xda, 0x8a, 0xda, 0x9b, 0x6e,
	0xb9, 0x51, 0xd4, 0xb3, 0xae, 0xf3, 0xeb, 0x10, 0xdd, 0x86, 0x48, 0x3b, 0xf1, 0xa6, 0x8f, 0x34,
	0x8a, 0x4a, 0xc5, 0x32, 0x8d, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x39, 0x39, 0x10, 0xad,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TbsClient is the client API for Tbs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TbsClient interface {
	Report(ctx context.Context, in *MeteringData, opts ...grpc.CallOption) (*types.Empty, error)
}

type tbsClient struct {
	cc *grpc.ClientConn
}

func NewTbsClient(cc *grpc.ClientConn) TbsClient {
	return &tbsClient{cc}
}

func (c *tbsClient) Report(ctx context.Context, in *MeteringData, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/tti.lorawan.v3.Tbs/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TbsServer is the server API for Tbs service.
type TbsServer interface {
	Report(context.Context, *MeteringData) (*types.Empty, error)
}

// UnimplementedTbsServer can be embedded to have forward compatible implementations.
type UnimplementedTbsServer struct {
}

func (*UnimplementedTbsServer) Report(ctx context.Context, req *MeteringData) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterTbsServer(s *grpc.Server, srv TbsServer) {
	s.RegisterService(&_Tbs_serviceDesc, srv)
}

func _Tbs_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeteringData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TbsServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tti.lorawan.v3.Tbs/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TbsServer).Report(ctx, req.(*MeteringData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tbs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tti.lorawan.v3.Tbs",
	HandlerType: (*TbsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _Tbs_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/tti/tenantbillingserver.proto",
}
