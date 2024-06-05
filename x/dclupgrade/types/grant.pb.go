// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/dclupgrade/grant.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Grant struct {
	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Time          int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Info          string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	SchemaVersion uint32 `protobuf:"varint,4,opt,name=schemaVersion,proto3" json:"schemaVersion,omitempty"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_a749b46e32743c76, []int{0}
}
func (m *Grant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return m.Size()
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Grant) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Grant) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Grant) GetSchemaVersion() uint32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*Grant)(nil), "zigbeealliance.distributedcomplianceledger.dclupgrade.Grant")
}

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/dclupgrade/grant.proto", fileDescriptor_a749b46e32743c76)
}

var fileDescriptor_a749b46e32743c76 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x1c, 0xc6, 0x1b, 0xef, 0x54, 0x2c, 0xdc, 0x52, 0x1c, 0xaa, 0x43, 0x28, 0xe2, 0xd0, 0xa5, 0x2d,
	0x28, 0x3e, 0xc0, 0xdd, 0xe2, 0xe0, 0x56, 0xc1, 0xc1, 0x45, 0xd2, 0xe4, 0x6f, 0x2e, 0xd0, 0x24,
	0x25, 0x49, 0x41, 0x7d, 0x03, 0x37, 0x1f, 0xc6, 0x87, 0x70, 0x3c, 0x9c, 0x1c, 0xa5, 0x7d, 0x11,
	0x69, 0x73, 0x72, 0x1e, 0xc8, 0x6d, 0x5f, 0xbe, 0x7c, 0xfc, 0x42, 0x7e, 0xe1, 0xfc, 0x45, 0xf0,
	0x0a, 0x80, 0xd4, 0xb5, 0x20, 0x8a, 0x42, 0xc1, 0x84, 0x75, 0x46, 0x54, 0xad, 0x03, 0x46, 0xb5,
	0x6c, 0x7c, 0x5b, 0x03, 0xe3, 0x60, 0x0a, 0x46, 0xeb, 0xb6, 0xe1, 0x86, 0x30, 0x28, 0xb8, 0x21,
	0xca, 0xe5, 0x8d, 0xd1, 0x4e, 0x47, 0x57, 0xdb, 0x88, 0x7c, 0x07, 0x22, 0xdf, 0x20, 0x4e, 0x4f,
	0xa8, 0xb6, 0x52, 0xdb, 0x87, 0x11, 0x52, 0xf8, 0x83, 0x27, 0x9e, 0xbd, 0xa2, 0x70, 0xff, 0x7a,
	0x78, 0x21, 0xba, 0x08, 0x0f, 0x09, 0x63, 0x06, 0xac, 0x8d, 0x51, 0x82, 0xd2, 0xa3, 0x45, 0xfc,
	0xf9, 0x9e, 0x1d, 0xaf, 0xc7, 0x73, 0x7f, 0x73, 0xeb, 0x8c, 0x50, 0xbc, 0xfc, 0x1d, 0x46, 0x51,
	0x38, 0x75, 0x42, 0x42, 0xbc, 0x97, 0xa0, 0x74, 0x52, 0x8e, 0x79, 0xe8, 0x84, 0x7a, 0xd4, 0xf1,
	0x64, 0x80, 0x94, 0x63, 0x8e, 0xce, 0xc3, 0x99, 0xa5, 0x4b, 0x90, 0xe4, 0x0e, 0x8c, 0x15, 0x5a,
	0xc5, 0xd3, 0x04, 0xa5, 0xb3, 0x72, 0xbb, 0x5c, 0xc0, 0x47, 0x87, 0xd1, 0xaa, 0xc3, 0xe8, 0xbb,
	0xc3, 0xe8, 0xad, 0xc7, 0xc1, 0xaa, 0xc7, 0xc1, 0x57, 0x8f, 0x83, 0xfb, 0x1b, 0x2e, 0xdc, 0xb2,
	0xad, 0x72, 0xaa, 0x65, 0xe1, 0x15, 0x64, 0xff, 0x69, 0xcc, 0x36, 0x12, 0xb2, 0xb5, 0xc8, 0xa7,
	0xbf, 0x2a, 0xdd, 0x73, 0x03, 0xb6, 0x3a, 0x18, 0x7f, 0x7e, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x67, 0x33, 0x51, 0x41, 0x90, 0x01, 0x00, 0x00,
}

func (m *Grant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Grant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Grant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SchemaVersion != 0 {
		i = encodeVarintGrant(dAtA, i, uint64(m.SchemaVersion))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintGrant(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Time != 0 {
		i = encodeVarintGrant(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGrant(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGrant(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Grant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGrant(uint64(l))
	}
	if m.Time != 0 {
		n += 1 + sovGrant(uint64(m.Time))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovGrant(uint64(l))
	}
	if m.SchemaVersion != 0 {
		n += 1 + sovGrant(uint64(m.SchemaVersion))
	}
	return n
}

func sovGrant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrant(x uint64) (n int) {
	return sovGrant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Grant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrant
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
			return fmt.Errorf("proto: Grant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrant
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
				return ErrInvalidLengthGrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrant
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
				return ErrInvalidLengthGrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrant
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
func skipGrant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrant
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
					return 0, ErrIntOverflowGrant
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
					return 0, ErrIntOverflowGrant
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
				return 0, ErrInvalidLengthGrant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGrant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGrant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGrant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGrant = fmt.Errorf("proto: unexpected end of group")
)
