// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/tti/application_api_key.proto

package ttipb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	ttnpb "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
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

type ApplicationAPIKey struct {
	ApplicationIDs       ttnpb.ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids"`
	APIKey               string                       `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ApplicationAPIKey) Reset()      { *m = ApplicationAPIKey{} }
func (*ApplicationAPIKey) ProtoMessage() {}
func (*ApplicationAPIKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_d23aba2d743b94a9, []int{0}
}
func (m *ApplicationAPIKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplicationAPIKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplicationAPIKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplicationAPIKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationAPIKey.Merge(m, src)
}
func (m *ApplicationAPIKey) XXX_Size() int {
	return m.Size()
}
func (m *ApplicationAPIKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationAPIKey.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationAPIKey proto.InternalMessageInfo

func (m *ApplicationAPIKey) GetApplicationIDs() ttnpb.ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIDs
	}
	return ttnpb.ApplicationIdentifiers{}
}

func (m *ApplicationAPIKey) GetAPIKey() string {
	if m != nil {
		return m.APIKey
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationAPIKey)(nil), "tti.lorawan.v3.ApplicationAPIKey")
	golang_proto.RegisterType((*ApplicationAPIKey)(nil), "tti.lorawan.v3.ApplicationAPIKey")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/tti/application_api_key.proto", fileDescriptor_d23aba2d743b94a9)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/tti/application_api_key.proto", fileDescriptor_d23aba2d743b94a9)
}

var fileDescriptor_d23aba2d743b94a9 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x31, 0x6b, 0x1b, 0x41,
	0x10, 0x85, 0x67, 0x44, 0xa2, 0x10, 0x05, 0x14, 0x5b, 0x95, 0x71, 0x31, 0x32, 0x09, 0x04, 0x13,
	0xd0, 0x2e, 0x44, 0xbf, 0x40, 0x47, 0x1a, 0xe3, 0x26, 0xb8, 0x4c, 0x63, 0x56, 0xd2, 0xfa, 0xb4,
	0x9c, 0xb2, 0xbb, 0xdc, 0xad, 0xcf, 0xb9, 0xce, 0xa5, 0x49, 0x95, 0x32, 0x65, 0x9a, 0x80, 0x4b,
	0x95, 0x2e, 0x5d, 0xaa, 0x74, 0xe9, 0x4a, 0xf8, 0x76, 0x1b, 0x97, 0x2e, 0x8d, 0xab, 0x90, 0x3b,
	0x81, 0x8f, 0xa8, 0x9b, 0x81, 0xf7, 0x3d, 0xe6, 0x0d, 0xaf, 0xc3, 0xe7, 0x26, 0x15, 0x67, 0x42,
	0x0f, 0x32, 0x27, 0x26, 0x09, 0x17, 0x56, 0x71, 0xe7, 0x14, 0x17, 0xd6, 0xce, 0xd5, 0x44, 0x38,
	0x65, 0xf4, 0xb1, 0xb0, 0xea, 0x38, 0x91, 0x05, 0xb3, 0xa9, 0x71, 0xa6, 0xd7, 0x75, 0x4e, 0xb1,
	0x35, 0xc4, 0xf2, 0xe1, 0xee, 0x28, 0x56, 0x6e, 0x76, 0x3a, 0x66, 0x13, 0xf3, 0x8d, 0x4b, 0x9d,
	0x9b, 0xc2, 0xa6, 0xe6, 0x7b, 0xc1, 0x2b, 0xf1, 0x64, 0x10, 0x4b, 0x3d, 0xc8, 0xc5, 0x5c, 0x4d,
	0x85, 0x93, 0x7c, 0x63, 0xa8, 0x2d, 0x77, 0x07, 0x0d, 0x8b, 0xd8, 0xc4, 0xa6, 0x86, 0xc7, 0xa7,
	0x27, 0xd5, 0x56, 0x2d, 0xd5, 0xb4, 0x96, 0xbf, 0xdf, 0x3c, 0x59, 0x4d, 0xa5, 0x76, 0xea, 0x44,
	0xc9, 0x34, 0xab, 0x45, 0xef, 0x16, 0xd8, 0xd9, 0x1e, 0x3d, 0x87, 0x18, 0x7d, 0x39, 0x38, 0x94,
	0x45, 0xcf, 0x74, 0xde, 0x36, 0x93, 0xa9, 0x69, 0xb6, 0x83, 0x7b, 0xb8, 0xff, 0xe6, 0xd3, 0x07,
	0xe6, 0x9c, 0x6e, 0xc4, 0x62, 0x0d, 0xf6, 0xe0, 0xd9, 0x3c, 0xa2, 0xa7, 0xe8, 0xe5, 0x0f, 0x6c,
	0x6d, 0xe1, 0x72, 0xd5, 0x07, 0xbf, 0xea, 0x77, 0x9b, 0xba, 0xcf, 0xd9, 0x51, 0x57, 0x34, 0xb9,
	0xac, 0xf7, 0xb1, 0xf3, 0x6a, 0xfd, 0xbe, 0x9d, 0xd6, 0x1e, 0xee, 0xbf, 0x8e, 0xb6, 0x9f, 0xa2,
	0x17, 0x69, 0x6b, 0x0b, 0xfd, 0xaa, 0xdf, 0xae, 0x8f, 0x3a, 0x6a, 0x0b, 0xab, 0x0e, 0x65, 0x11,
	0xfd, 0xc1, 0x65, 0x49, 0x78, 0x53, 0x12, 0xde, 0x96, 0x04, 0x77, 0x25, 0xc1, 0x7d, 0x49, 0xf0,
	0x50, 0x12, 0x3c, 0x96, 0x84, 0xe7, 0x9e, 0xf0, 0xc2, 0x13, 0x5c, 0x7a, 0xc2, 0x85, 0x27, 0xb8,
	0xf2, 0x04, 0xd7, 0x9e, 0x60, 0xe9, 0x09, 0x6f, 0x3c, 0xe1, 0xad, 0x27, 0xb8, 0xf3, 0x84, 0xf7,
	0x9e, 0xe0, 0xc1, 0x13, 0x3e, 0x7a, 0x82, 0xf3, 0x40, 0x70, 0x11, 0x08, 0x7f, 0x06, 0x82, 0x5f,
	0x81, 0xf0, 0x77, 0x20, 0xb8, 0x0c, 0x04, 0x8b, 0x40, 0x78, 0x15, 0x08, 0xaf, 0x03, 0xe1, 0x57,
	0x1e, 0x1b, 0xe6, 0x66, 0xd2, 0xcd, 0x94, 0x8e, 0x33, 0xa6, 0xa5, 0x3b, 0x33, 0x69, 0xf2, 0x5f,
	0x2d, 0xf2, 0x21, 0xb7, 0x49, 0xfc, 0xaf, 0x19, 0x76, 0x3c, 0x6e, 0x57, 0x1f, 0x1e, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x31, 0xec, 0x07, 0x62, 0x3b, 0x02, 0x00, 0x00,
}

func (this *ApplicationAPIKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApplicationAPIKey)
	if !ok {
		that2, ok := that.(ApplicationAPIKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApplicationIDs.Equal(&that1.ApplicationIDs) {
		return false
	}
	if this.APIKey != that1.APIKey {
		return false
	}
	return true
}
func (m *ApplicationAPIKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationAPIKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplicationAPIKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.APIKey) > 0 {
		i -= len(m.APIKey)
		copy(dAtA[i:], m.APIKey)
		i = encodeVarintApplicationApiKey(dAtA, i, uint64(len(m.APIKey)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ApplicationIDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintApplicationApiKey(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintApplicationApiKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovApplicationApiKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedApplicationAPIKey(r randyApplicationApiKey, easy bool) *ApplicationAPIKey {
	this := &ApplicationAPIKey{}
	v1 := ttnpb.NewPopulatedApplicationIdentifiers(r, easy)
	this.ApplicationIDs = *v1
	this.APIKey = randStringApplicationApiKey(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyApplicationApiKey interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneApplicationApiKey(r randyApplicationApiKey) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringApplicationApiKey(r randyApplicationApiKey) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneApplicationApiKey(r)
	}
	return string(tmps)
}
func randUnrecognizedApplicationApiKey(r randyApplicationApiKey, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldApplicationApiKey(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldApplicationApiKey(dAtA []byte, r randyApplicationApiKey, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateApplicationApiKey(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateApplicationApiKey(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateApplicationApiKey(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateApplicationApiKey(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateApplicationApiKey(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateApplicationApiKey(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateApplicationApiKey(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ApplicationAPIKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ApplicationIDs.Size()
	n += 1 + l + sovApplicationApiKey(uint64(l))
	l = len(m.APIKey)
	if l > 0 {
		n += 1 + l + sovApplicationApiKey(uint64(l))
	}
	return n
}

func sovApplicationApiKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApplicationApiKey(x uint64) (n int) {
	return sovApplicationApiKey((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *ApplicationAPIKey) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ApplicationAPIKey{`,
		`ApplicationIDs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ApplicationIDs), "ApplicationIdentifiers", "ttnpb.ApplicationIdentifiers", 1), `&`, ``, 1) + `,`,
		`APIKey:` + fmt.Sprintf("%v", this.APIKey) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApplicationApiKey(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ApplicationAPIKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplicationApiKey
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ApplicationAPIKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationAPIKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationApiKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApplicationApiKey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationApiKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ApplicationIDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationApiKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApplicationApiKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationApiKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplicationApiKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplicationApiKey
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApplicationApiKey
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApplicationApiKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplicationApiKey
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApplicationApiKey
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApplicationApiKey
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApplicationApiKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApplicationApiKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApplicationApiKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApplicationApiKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplicationApiKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApplicationApiKey = fmt.Errorf("proto: unexpected end of group")
)
