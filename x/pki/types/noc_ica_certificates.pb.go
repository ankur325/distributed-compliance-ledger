// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/pki/noc_ica_certificates.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type NocIcaCertificates struct {
	Vid   int32          `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty" validate:"gte=1,lte=65535"`
	Certs []*Certificate `protobuf:"bytes,2,rep,name=certs,proto3" json:"certs,omitempty"`
}

func (m *NocIcaCertificates) Reset()         { *m = NocIcaCertificates{} }
func (m *NocIcaCertificates) String() string { return proto.CompactTextString(m) }
func (*NocIcaCertificates) ProtoMessage()    {}
func (*NocIcaCertificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_59301d987fd5414e, []int{0}
}
func (m *NocIcaCertificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NocIcaCertificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NocIcaCertificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NocIcaCertificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NocIcaCertificates.Merge(m, src)
}
func (m *NocIcaCertificates) XXX_Size() int {
	return m.Size()
}
func (m *NocIcaCertificates) XXX_DiscardUnknown() {
	xxx_messageInfo_NocIcaCertificates.DiscardUnknown(m)
}

var xxx_messageInfo_NocIcaCertificates proto.InternalMessageInfo

func (m *NocIcaCertificates) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *NocIcaCertificates) GetCerts() []*Certificate {
	if m != nil {
		return m.Certs
	}
	return nil
}

func init() {
	proto.RegisterType((*NocIcaCertificates)(nil), "zigbeealliance.distributedcomplianceledger.pki.NocIcaCertificates")
}

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/pki/noc_ica_certificates.proto", fileDescriptor_59301d987fd5414e)
}

var fileDescriptor_59301d987fd5414e = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0xac, 0xca, 0x4c, 0x4f,
	0x4a, 0x4d, 0x4d, 0xcc, 0xc9, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x4f, 0xc9, 0x2c, 0x2e, 0x29,
	0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x49, 0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4,
	0xa7, 0x16, 0xe9, 0x17, 0x64, 0x67, 0xea, 0xe7, 0xe5, 0x27, 0xc7, 0x67, 0x26, 0x27, 0xc6, 0x27,
	0xa7, 0x16, 0x95, 0x64, 0xa6, 0x65, 0x26, 0x27, 0x96, 0xa4, 0x16, 0xeb, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0xe9, 0xa1, 0x1a, 0xa5, 0x87, 0xc7, 0x28, 0xbd, 0x82, 0xec, 0x4c, 0x29, 0x91, 0xf4,
	0xfc, 0xf4, 0x7c, 0xb0, 0x56, 0x7d, 0x10, 0x0b, 0x62, 0x8a, 0x94, 0x03, 0x89, 0x0e, 0x42, 0x72,
	0x08, 0xc4, 0x04, 0xa5, 0x99, 0x8c, 0x5c, 0x42, 0x7e, 0xf9, 0xc9, 0x9e, 0xc9, 0x89, 0xce, 0x48,
	0x8e, 0x14, 0x32, 0xe0, 0x62, 0x2e, 0xcb, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x75, 0x92,
	0xfb, 0x74, 0x4f, 0x5e, 0xaa, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0xd5, 0x4a, 0x29, 0xbd,
	0x24, 0xd5, 0xd6, 0x50, 0x27, 0xa7, 0x24, 0xd5, 0xd6, 0xcc, 0xd4, 0xd4, 0xd8, 0x54, 0x29, 0x08,
	0xa4, 0x54, 0x28, 0x90, 0x8b, 0x15, 0x64, 0x7a, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0x35, 0x89, 0x1e, 0xd4, 0x43, 0xb2, 0x3e, 0x08, 0x62, 0x92, 0x53, 0xdc, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xb9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x43, 0xec, 0xd1, 0xc5, 0x16, 0x06, 0xba, 0x08, 0x9b, 0x74, 0xa1, 0xa1, 0x50, 0x01,
	0x0e, 0x87, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x10, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xbf, 0x0d, 0x78, 0x1d, 0xd7, 0x01, 0x00, 0x00,
}

func (m *NocIcaCertificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NocIcaCertificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NocIcaCertificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Certs) > 0 {
		for iNdEx := len(m.Certs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNocIcaCertificates(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Vid != 0 {
		i = encodeVarintNocIcaCertificates(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNocIcaCertificates(dAtA []byte, offset int, v uint64) int {
	offset -= sovNocIcaCertificates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NocIcaCertificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovNocIcaCertificates(uint64(m.Vid))
	}
	if len(m.Certs) > 0 {
		for _, e := range m.Certs {
			l = e.Size()
			n += 1 + l + sovNocIcaCertificates(uint64(l))
		}
	}
	return n
}

func sovNocIcaCertificates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNocIcaCertificates(x uint64) (n int) {
	return sovNocIcaCertificates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NocIcaCertificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNocIcaCertificates
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
			return fmt.Errorf("proto: NocIcaCertificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NocIcaCertificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNocIcaCertificates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNocIcaCertificates
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
				return ErrInvalidLengthNocIcaCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNocIcaCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certs = append(m.Certs, &Certificate{})
			if err := m.Certs[len(m.Certs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNocIcaCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNocIcaCertificates
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
func skipNocIcaCertificates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNocIcaCertificates
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
					return 0, ErrIntOverflowNocIcaCertificates
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
					return 0, ErrIntOverflowNocIcaCertificates
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
				return 0, ErrInvalidLengthNocIcaCertificates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNocIcaCertificates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNocIcaCertificates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNocIcaCertificates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNocIcaCertificates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNocIcaCertificates = fmt.Errorf("proto: unexpected end of group")
)