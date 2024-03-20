// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/revoked_noc_root_certificates.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type RevokedNocRootCertificates struct {
	Subject      string         `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId string         `protobuf:"bytes,2,opt,name=subjectKeyId,proto3" json:"subjectKeyId,omitempty"`
	Certs        []*Certificate `protobuf:"bytes,3,rep,name=certs,proto3" json:"certs,omitempty"`
}

func (m *RevokedNocRootCertificates) Reset()         { *m = RevokedNocRootCertificates{} }
func (m *RevokedNocRootCertificates) String() string { return proto.CompactTextString(m) }
func (*RevokedNocRootCertificates) ProtoMessage()    {}
func (*RevokedNocRootCertificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab9c4f154f441df5, []int{0}
}
func (m *RevokedNocRootCertificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RevokedNocRootCertificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RevokedNocRootCertificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RevokedNocRootCertificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokedNocRootCertificates.Merge(m, src)
}
func (m *RevokedNocRootCertificates) XXX_Size() int {
	return m.Size()
}
func (m *RevokedNocRootCertificates) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokedNocRootCertificates.DiscardUnknown(m)
}

var xxx_messageInfo_RevokedNocRootCertificates proto.InternalMessageInfo

func (m *RevokedNocRootCertificates) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *RevokedNocRootCertificates) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *RevokedNocRootCertificates) GetCerts() []*Certificate {
	if m != nil {
		return m.Certs
	}
	return nil
}

func init() {
	proto.RegisterType((*RevokedNocRootCertificates)(nil), "zigbeealliance.distributedcomplianceledger.pki.RevokedNocRootCertificates")
}

func init() {
	proto.RegisterFile("pki/revoked_noc_root_certificates.proto", fileDescriptor_ab9c4f154f441df5)
}

var fileDescriptor_ab9c4f154f441df5 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0xc8, 0xce, 0xd4,
	0x2f, 0x4a, 0x2d, 0xcb, 0xcf, 0x4e, 0x4d, 0x89, 0xcf, 0xcb, 0x4f, 0x8e, 0x2f, 0xca, 0xcf, 0x2f,
	0x89, 0x4f, 0x4e, 0x2d, 0x2a, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0x49, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xd2, 0xab, 0xca, 0x4c, 0x4f, 0x4a, 0x4d, 0x4d, 0xcc, 0xc9, 0xc9, 0x4c,
	0xcc, 0x4b, 0x4e, 0xd5, 0x4b, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x49,
	0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4, 0xa7, 0x16, 0xe9, 0x15, 0x64, 0x67, 0x4a,
	0x89, 0x82, 0x0c, 0x46, 0x32, 0x07, 0x62, 0x8c, 0xd2, 0x5a, 0x46, 0x2e, 0xa9, 0x20, 0x88, 0x75,
	0x7e, 0xf9, 0xc9, 0x41, 0xf9, 0xf9, 0x25, 0xce, 0x48, 0x76, 0x09, 0x49, 0x70, 0xb1, 0x17, 0x97,
	0x26, 0x65, 0xa5, 0x26, 0x97, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x4a,
	0x5c, 0x3c, 0x50, 0xa6, 0x77, 0x6a, 0xa5, 0x67, 0x8a, 0x04, 0x13, 0x58, 0x1a, 0x45, 0x4c, 0x28,
	0x90, 0x8b, 0x15, 0x64, 0x63, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x35, 0x89, 0x6e,
	0xd6, 0x43, 0x72, 0x4a, 0x10, 0xc4, 0x24, 0xa7, 0xb8, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0x72, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87,
	0xd8, 0xa3, 0x0b, 0xb3, 0x48, 0x1f, 0xc9, 0x22, 0x5d, 0x84, 0x4d, 0xba, 0x10, 0xab, 0xf4, 0x2b,
	0xf4, 0x41, 0x61, 0x53, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x16, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf9, 0x0d, 0x7d, 0xa5, 0x88, 0x01, 0x00, 0x00,
}

func (m *RevokedNocRootCertificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RevokedNocRootCertificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RevokedNocRootCertificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintRevokedNocRootCertificates(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintRevokedNocRootCertificates(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintRevokedNocRootCertificates(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRevokedNocRootCertificates(dAtA []byte, offset int, v uint64) int {
	offset -= sovRevokedNocRootCertificates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RevokedNocRootCertificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovRevokedNocRootCertificates(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovRevokedNocRootCertificates(uint64(l))
	}
	if len(m.Certs) > 0 {
		for _, e := range m.Certs {
			l = e.Size()
			n += 1 + l + sovRevokedNocRootCertificates(uint64(l))
		}
	}
	return n
}

func sovRevokedNocRootCertificates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRevokedNocRootCertificates(x uint64) (n int) {
	return sovRevokedNocRootCertificates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RevokedNocRootCertificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRevokedNocRootCertificates
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
			return fmt.Errorf("proto: RevokedNocRootCertificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RevokedNocRootCertificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevokedNocRootCertificates
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
				return ErrInvalidLengthRevokedNocRootCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRevokedNocRootCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevokedNocRootCertificates
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
				return ErrInvalidLengthRevokedNocRootCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRevokedNocRootCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevokedNocRootCertificates
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
				return ErrInvalidLengthRevokedNocRootCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRevokedNocRootCertificates
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
			skippy, err := skipRevokedNocRootCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRevokedNocRootCertificates
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
func skipRevokedNocRootCertificates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRevokedNocRootCertificates
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
					return 0, ErrIntOverflowRevokedNocRootCertificates
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
					return 0, ErrIntOverflowRevokedNocRootCertificates
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
				return 0, ErrInvalidLengthRevokedNocRootCertificates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRevokedNocRootCertificates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRevokedNocRootCertificates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRevokedNocRootCertificates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRevokedNocRootCertificates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRevokedNocRootCertificates = fmt.Errorf("proto: unexpected end of group")
)