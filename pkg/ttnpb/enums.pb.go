// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/enums.proto

package ttnpb

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
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

type DownlinkPathConstraint int32

const (
	// Indicates that the gateway can be selected for downlink without constraints by the Network Server.
	DOWNLINK_PATH_CONSTRAINT_NONE DownlinkPathConstraint = 0
	// Indicates that the gateway can be selected for downlink only if no other or better gateway can be selected.
	DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER DownlinkPathConstraint = 1
	// Indicates that this gateway will never be selected for downlink, even if that results in no available downlink path.
	DOWNLINK_PATH_CONSTRAINT_NEVER DownlinkPathConstraint = 2
)

var DownlinkPathConstraint_name = map[int32]string{
	0: "DOWNLINK_PATH_CONSTRAINT_NONE",
	1: "DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER",
	2: "DOWNLINK_PATH_CONSTRAINT_NEVER",
}

var DownlinkPathConstraint_value = map[string]int32{
	"DOWNLINK_PATH_CONSTRAINT_NONE":         0,
	"DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER": 1,
	"DOWNLINK_PATH_CONSTRAINT_NEVER":        2,
}

func (DownlinkPathConstraint) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{0}
}

// State enum defines states that an entity can be in.
type State int32

const (
	// Denotes that the entity has been requested and is pending review by an admin.
	STATE_REQUESTED State = 0
	// Denotes that the entity has been reviewed and approved by an admin.
	STATE_APPROVED State = 1
	// Denotes that the entity has been reviewed and rejected by an admin.
	STATE_REJECTED State = 2
	// Denotes that the entity has been flagged and is pending review by an admin.
	STATE_FLAGGED State = 3
	// Denotes that the entity has been reviewed and suspended by an admin.
	STATE_SUSPENDED State = 4
)

var State_name = map[int32]string{
	0: "STATE_REQUESTED",
	1: "STATE_APPROVED",
	2: "STATE_REJECTED",
	3: "STATE_FLAGGED",
	4: "STATE_SUSPENDED",
}

var State_value = map[string]int32{
	"STATE_REQUESTED": 0,
	"STATE_APPROVED":  1,
	"STATE_REJECTED":  2,
	"STATE_FLAGGED":   3,
	"STATE_SUSPENDED": 4,
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{1}
}

type ClusterRole int32

const (
	ClusterRole_NONE                         ClusterRole = 0
	ClusterRole_ENTITY_REGISTRY              ClusterRole = 1
	ClusterRole_ACCESS                       ClusterRole = 2
	ClusterRole_GATEWAY_SERVER               ClusterRole = 3
	ClusterRole_NETWORK_SERVER               ClusterRole = 4
	ClusterRole_APPLICATION_SERVER           ClusterRole = 5
	ClusterRole_JOIN_SERVER                  ClusterRole = 6
	ClusterRole_CRYPTO_SERVER                ClusterRole = 7
	ClusterRole_DEVICE_TEMPLATE_CONVERTER    ClusterRole = 8
	ClusterRole_DEVICE_CLAIMING_SERVER       ClusterRole = 9
	ClusterRole_GATEWAY_CONFIGURATION_SERVER ClusterRole = 10
	ClusterRole_QR_CODE_GENERATOR            ClusterRole = 11
	ClusterRole_TENANT_BILLING_SERVER        ClusterRole = 101
)

var ClusterRole_name = map[int32]string{
	0:   "NONE",
	1:   "ENTITY_REGISTRY",
	2:   "ACCESS",
	3:   "GATEWAY_SERVER",
	4:   "NETWORK_SERVER",
	5:   "APPLICATION_SERVER",
	6:   "JOIN_SERVER",
	7:   "CRYPTO_SERVER",
	8:   "DEVICE_TEMPLATE_CONVERTER",
	9:   "DEVICE_CLAIMING_SERVER",
	10:  "GATEWAY_CONFIGURATION_SERVER",
	11:  "QR_CODE_GENERATOR",
	101: "TENANT_BILLING_SERVER",
}

var ClusterRole_value = map[string]int32{
	"NONE":                         0,
	"ENTITY_REGISTRY":              1,
	"ACCESS":                       2,
	"GATEWAY_SERVER":               3,
	"NETWORK_SERVER":               4,
	"APPLICATION_SERVER":           5,
	"JOIN_SERVER":                  6,
	"CRYPTO_SERVER":                7,
	"DEVICE_TEMPLATE_CONVERTER":    8,
	"DEVICE_CLAIMING_SERVER":       9,
	"GATEWAY_CONFIGURATION_SERVER": 10,
	"QR_CODE_GENERATOR":            11,
	"TENANT_BILLING_SERVER":        101,
}

func (ClusterRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{2}
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
}

func init() { proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb)
}

var fileDescriptor_e36318a1e2f407cb = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xb1, 0x4f, 0x1b, 0x49,
	0x14, 0xc6, 0x67, 0x8c, 0xe1, 0xb8, 0x41, 0x07, 0xcb, 0x20, 0x90, 0x40, 0xc7, 0xd3, 0x5d, 0xa4,
	0x14, 0x41, 0xc1, 0x2e, 0xf8, 0x0b, 0x96, 0xdd, 0x87, 0x59, 0x30, 0xb3, 0xcb, 0xec, 0x60, 0xe4,
	0x34, 0x2b, 0x1b, 0x39, 0xb6, 0x65, 0xb3, 0x6b, 0xd9, 0xe3, 0xd0, 0x52, 0x52, 0x52, 0xa6, 0x8c,
	0x92, 0x06, 0x29, 0x0d, 0x25, 0x25, 0x25, 0x25, 0x25, 0x25, 0xbb, 0xdb, 0x50, 0x52, 0x52, 0x46,
	0x36, 0x76, 0x08, 0x05, 0xdd, 0xbc, 0xdf, 0x7c, 0xef, 0xbd, 0xef, 0x93, 0x66, 0xd8, 0x6a, 0x3b,
	0xea, 0x56, 0x4e, 0x2a, 0xe1, 0x7a, 0x4f, 0x57, 0x8e, 0x5a, 0xf9, 0x4a, 0xa7, 0x99, 0xaf, 0x85,
	0xfd, 0xe3, 0x5e, 0xae, 0xd3, 0x8d, 0x74, 0xc4, 0x67, 0xb5, 0x0e, 0x73, 0x23, 0x49, 0xee, 0xcb,
	0xc6, 0xca, 0x7a, 0xbd, 0xa9, 0x1b, 0xfd, 0x6a, 0xee, 0x28, 0x3a, 0xce, 0xd7, 0xa3, 0x7a, 0x94,
	0x1f, 0xca, 0xaa, 0xfd, 0xcf, 0xc3, 0x6a, 0x58, 0x0c, 0x4f, 0xcf, 0xed, 0x6b, 0xe7, 0x94, 0x2d,
	0xd9, 0xd1, 0x49, 0xd8, 0x6e, 0x86, 0x2d, 0xaf, 0xa2, 0x1b, 0x56, 0x14, 0xf6, 0x74, 0xb7, 0xd2,
	0x0c, 0x35, 0xff, 0x9f, 0xad, 0xda, 0xee, 0xa1, 0x28, 0x3a, 0x62, 0x37, 0xf0, 0x4c, 0xb5, 0x1d,
	0x58, 0xae, 0xf0, 0x95, 0x34, 0x1d, 0xa1, 0x02, 0xe1, 0x0a, 0x34, 0x08, 0xff, 0xc0, 0xde, 0xbf,
	0x29, 0xf1, 0x24, 0x6e, 0xa1, 0x0c, 0x5c, 0xb5, 0x8d, 0xd2, 0xa0, 0xfc, 0x1d, 0x83, 0xb7, 0xa7,
	0x61, 0x09, 0xa5, 0x91, 0x59, 0xc9, 0x9e, 0xfd, 0x00, 0xb2, 0xd6, 0x65, 0x93, 0xbe, 0xae, 0xe8,
	0x1a, 0x5f, 0x60, 0x73, 0xbe, 0x32, 0x15, 0x06, 0x12, 0xf7, 0x0f, 0xd0, 0x57, 0x68, 0x1b, 0x84,
	0x73, 0x36, 0xfb, 0x0c, 0x4d, 0xcf, 0x93, 0x6e, 0x09, 0x6d, 0x83, 0xbe, 0x30, 0x89, 0x3b, 0x68,
	0x0d, 0x74, 0x19, 0x3e, 0xcf, 0xfe, 0x79, 0x66, 0x5b, 0x45, 0xb3, 0x50, 0x40, 0xdb, 0x98, 0x78,
	0x99, 0xe7, 0x1f, 0xf8, 0x1e, 0x0a, 0x1b, 0x6d, 0x23, 0x3b, 0xda, 0xf9, 0x33, 0xc3, 0x66, 0xac,
	0x76, 0xbf, 0xa7, 0x6b, 0x5d, 0x19, 0xb5, 0x6b, 0x7c, 0x9a, 0x65, 0x47, 0x11, 0x17, 0xd8, 0x1c,
	0x0a, 0xe5, 0xa8, 0x72, 0x20, 0xb1, 0xe0, 0xf8, 0x4a, 0x96, 0x0d, 0xca, 0x19, 0x9b, 0x32, 0x2d,
	0x0b, 0x7d, 0xdf, 0xc8, 0x0c, 0x96, 0x17, 0x4c, 0x85, 0x87, 0x66, 0x39, 0xf0, 0x51, 0x0e, 0x82,
	0x4c, 0x0c, 0x98, 0x40, 0x75, 0xe8, 0xca, 0xdd, 0x31, 0xcb, 0xf2, 0x25, 0xc6, 0x4d, 0xcf, 0x2b,
	0x3a, 0x96, 0xa9, 0x1c, 0x57, 0x8c, 0xf9, 0x24, 0x9f, 0x63, 0x33, 0x3b, 0xae, 0xf3, 0x1b, 0x4c,
	0x0d, 0x9c, 0x5b, 0xb2, 0xec, 0x29, 0x77, 0x8c, 0xfe, 0xe2, 0xab, 0x6c, 0xd9, 0xc6, 0x92, 0x63,
	0x61, 0xa0, 0x70, 0xcf, 0x2b, 0x0e, 0x32, 0x58, 0xae, 0x28, 0xa1, 0x54, 0x28, 0x8d, 0x69, 0xbe,
	0xc2, 0x96, 0x46, 0xd7, 0x56, 0xd1, 0x74, 0xf6, 0x1c, 0x51, 0x18, 0xb7, 0xfe, 0xcd, 0xff, 0x63,
	0xff, 0x8e, 0xed, 0x59, 0xae, 0xd8, 0x72, 0x0a, 0x07, 0xf2, 0x95, 0x01, 0xc6, 0x17, 0xd9, 0xfc,
	0xbe, 0x0c, 0x2c, 0xd7, 0xc6, 0xa0, 0x80, 0x02, 0xa5, 0xa9, 0x5c, 0x69, 0xcc, 0xf0, 0x65, 0xb6,
	0xa8, 0x50, 0x98, 0x42, 0x05, 0x9b, 0x4e, 0xb1, 0xf8, 0xc7, 0xcc, 0xda, 0xe6, 0x77, 0x7a, 0x13,
	0x03, 0xbd, 0x8d, 0x81, 0xde, 0xc5, 0x40, 0xee, 0x63, 0x20, 0x0f, 0x31, 0x90, 0xc7, 0x18, 0xc8,
	0x53, 0x0c, 0xf4, 0x34, 0x01, 0x7a, 0x96, 0x00, 0xb9, 0x48, 0x80, 0x5e, 0x26, 0x40, 0xae, 0x12,
	0x20, 0xd7, 0x09, 0x90, 0x9b, 0x04, 0xe8, 0x6d, 0x02, 0xf4, 0x2e, 0x01, 0x72, 0x9f, 0x00, 0x7d,
	0x48, 0x80, 0x3c, 0x26, 0x40, 0x9f, 0x12, 0x20, 0xa7, 0x29, 0x90, 0xb3, 0x14, 0xe8, 0x79, 0x0a,
	0xe4, 0x6b, 0x0a, 0xf4, 0x5b, 0x0a, 0xe4, 0x22, 0x05, 0x72, 0x99, 0x02, 0xbd, 0x4a, 0x81, 0x5e,
	0xa7, 0x40, 0x3f, 0x7d, 0xac, 0x47, 0x39, 0xdd, 0xa8, 0xe9, 0x46, 0x33, 0xac, 0xf7, 0x72, 0x61,
	0x4d, 0x9f, 0x44, 0xdd, 0x56, 0xfe, 0xf5, 0xf7, 0xe8, 0xb4, 0xea, 0x79, 0xad, 0xc3, 0x4e, 0xb5,
	0x3a, 0x35, 0x7c, 0xe0, 0x1b, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0x20, 0xbd, 0xd6, 0x40,
	0x03, 0x00, 0x00,
}

func (x DownlinkPathConstraint) String() string {
	s, ok := DownlinkPathConstraint_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x State) String() string {
	s, ok := State_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ClusterRole) String() string {
	s, ok := ClusterRole_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
