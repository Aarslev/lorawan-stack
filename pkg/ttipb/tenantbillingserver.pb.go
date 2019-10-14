// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/tti/tenantbillingserver.proto

package ttipb

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
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
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x2f, 0xef, 0x13, 0x41,
	0x10, 0xdd, 0x85, 0xa4, 0xa2, 0x04, 0xc4, 0x91, 0x20, 0x4a, 0x33, 0x02, 0x2c, 0xb7, 0x9b, 0xd0,
	0x4f, 0x00, 0x01, 0x89, 0x01, 0x54, 0xdd, 0xde, 0xb1, 0x6c, 0x37, 0xbd, 0xee, 0x5e, 0xf6, 0xa6,
	0x57, 0xea, 0x1a, 0x54, 0x25, 0x09, 0x06, 0x49, 0x50, 0x95, 0x95, 0x95, 0x95, 0x95, 0x4d, 0x30,
	0x95, 0xbd, 0x5d, 0x44, 0x65, 0x65, 0x25, 0xe9, 0xf5, 0x20, 0xbf, 0x7f, 0x6e, 0x26, 0x6f, 0xe6,
	0xbd, 0x79, 0xf3, 0xda, 0x3c, 0xb3, 0x4e, 0x4c, 0x84, 0x89, 0x0b, 0x14, 0xe9, 0x90, 0x8b, 0x5c,
	0x73, 0x44, 0xcd, 0x51, 0x1a, 0x61, 0x30, 0xd1, 0x59, 0xa6, 0x8d, 0x2a, 0xa4, 0x2b, 0xa5, 0x63,
	0xb9, 0xb3, 0x68, 0xa3, 0x47, 0x88, 0x9a, 0x35, 0x4b, 0xac, 0xec, 0x75, 0x5e, 0x29, 0x8d, 0x83,
	0x71, 0xc2, 0x52, 0x3b, 0xe2, 0xd2, 0x94, 0x76, 0x9a, 0x3b, 0xfb, 0x65, 0xca, 0xeb, 0xe1, 0x34,
	0x56, 0xd2, 0xc4, 0xa5, 0xc8, 0xf4, 0x27, 0x81, 0x92, 0xdf, 0x2a, 0x2e, 0x94, 0x9d, 0xf8, 0x0a,
	0x85, 0xb2, 0xca, 0x5e, 0x96, 0x93, 0xf1, 0xe7, 0xba, 0xab, 0x9b, 0xba, 0x6a, 0xc6, 0xbb, 0xca,
	0x5a, 0x95, 0xc9, 0xfa, 0x56, 0x61, 0x8c, 0x45, 0x81, 0xda, 0x9a, 0xa2, 0x41, 0x9f, 0x36, 0xe8,
	0x7f, 0x0e, 0x39, 0xca, 0x71, 0xda, 0x80, 0xcf, 0xef, 0x76, 0x9b, 0xe9, 0x54, 0x9a, 0xa2, 0x39,
	0xe7, 0x65, 0xbf, 0x7d, 0xff, 0x63, 0x52, 0x44, 0x1f, 0xda, 0xad, 0xf7, 0x32, 0xb7, 0x0e, 0xa3,
	0x2e, 0xbb, 0xee, 0x99, 0xbd, 0x93, 0x28, 0x9d, 0x36, 0xea, 0x8d, 0x40, 0xd1, 0x79, 0xc2, 0x2e,
	0x8a, 0xec, 0x9f, 0x22, 0x7b, 0x7b, 0x56, 0x7c, 0xf6, 0xf8, 0xeb, 0xef, 0x3f, 0xdf, 0xef, 0x3d,
	0x8c, 0x1e, 0x70, 0x4c, 0x0a, 0xee, 0x6a, 0xaa, 0xd7, 0xbf, 0xe8, 0xa6, 0x02, 0xba, 0xad, 0x80,
	0xee, 0x2a, 0x20, 0xfb, 0x0a, 0xc8, 0xa1, 0x02, 0x72, 0xac, 0x80, 0x9c, 0x2a, 0xa0, 0x33, 0x0f,
	0x74, 0xee, 0x81, 0x2c, 0x3c, 0xd0, 0xa5, 0x07, 0xb2, 0xf2, 0x40, 0xd6, 0x1e, 0xc8, 0xc6, 0x03,
	0xdd, 0x7a, 0xa0, 0x3b, 0x0f, 0x64, 0xef, 0x81, 0x1e, 0x3c, 0x90, 0xa3, 0x07, 0x7a, 0xf2, 0x40,
	0x66, 0x01, 0xc8, 0x3c, 0x00, 0xfd, 0x16, 0x80, 0xfc, 0x08, 0x40, 0x7f, 0x06, 0x20, 0x8b, 0x00,
	0x64, 0x19, 0x80, 0xae, 0x02, 0xd0, 0x75, 0x00, 0xda, 0x7f, 0xa1, 0x2c, 0xc3, 0x81, 0xc4, 0xc1,
	0x39, 0x4d, 0x66, 0x24, 0x4e, 0xac, 0x1b, 0xde, 0x88, 0x3e, 0x1f, 0xaa, 0xf3, 0x33, 0xf2, 0x24,
	0x69, 0xd5, 0x4e, 0x7a, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa3, 0x4f, 0x85, 0x1c, 0x02,
	0x00, 0x00,
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
