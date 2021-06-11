// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/events/content.proto

package events

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_opencontainers_go_digest "github.com/opencontainers/go-digest"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type ContentDelete struct {
	Digest               github_com_opencontainers_go_digest.Digest `protobuf:"bytes,1,opt,name=digest,proto3,customtype=github.com/opencontainers/go-digest.Digest" json:"digest"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *ContentDelete) Reset()      { *m = ContentDelete{} }
func (*ContentDelete) ProtoMessage() {}
func (*ContentDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfb34b8b808e2ecd, []int{0}
}
func (m *ContentDelete) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentDelete.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentDelete.Merge(m, src)
}
func (m *ContentDelete) XXX_Size() int {
	return m.Size()
}
func (m *ContentDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentDelete.DiscardUnknown(m)
}

var xxx_messageInfo_ContentDelete proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ContentDelete)(nil), "containerd.events.ContentDelete")
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/events/content.proto", fileDescriptor_dfb34b8b808e2ecd)
}

var fileDescriptor_dfb34b8b808e2ecd = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0x41, 0x66, 0x26, 0x16, 0x64, 0xea, 0xa7, 0x96, 0xa5, 0xe6, 0x95, 0x14, 0x83, 0x45, 0x53,
	0xf3, 0x4a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x11, 0x8a, 0xf4, 0x20, 0x0a, 0xa4,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44, 0xa1, 0x94, 0x03, 0x41, 0x3b,
	0xc0, 0xea, 0x92, 0x4a, 0xd3, 0xf4, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0xf4, 0xd3, 0x32, 0x53,
	0x73, 0x52, 0x0a, 0x12, 0x4b, 0x32, 0x20, 0x26, 0x28, 0x45, 0x73, 0xf1, 0x3a, 0x43, 0xec, 0x76,
	0x49, 0xcd, 0x49, 0x2d, 0x49, 0x15, 0xf2, 0xe2, 0x62, 0x4b, 0xc9, 0x4c, 0x4f, 0x2d, 0x2e, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x32, 0x3a, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0x2d,
	0x24, 0xab, 0xf2, 0x0b, 0x52, 0xf3, 0xe0, 0x76, 0x14, 0xeb, 0xa7, 0xe7, 0xeb, 0x42, 0xb4, 0xe8,
	0xb9, 0x80, 0xa9, 0x20, 0xa8, 0x09, 0x4e, 0x01, 0x27, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7,
	0xd0, 0xf0, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x5c, 0xf0, 0x45, 0x8e, 0x31, 0xca, 0x88, 0x84, 0x00, 0xb2, 0x86, 0x50, 0x11, 0x0c, 0x11,
	0x8c, 0x49, 0x6c, 0x60, 0x97, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x78, 0x99, 0xee,
	0x61, 0x01, 0x00, 0x00,
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContentDelete) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	case "digest":
		return string(m.Digest), len(m.Digest) > 0
	}
	return "", false
}
func (m *ContentDelete) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentDelete) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentDelete) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Digest) > 0 {
		i -= len(m.Digest)
		copy(dAtA[i:], m.Digest)
		i = encodeVarintContent(dAtA, i, uint64(len(m.Digest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContent(dAtA []byte, offset int, v uint64) int {
	offset -= sovContent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContentDelete) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Digest)
	if l > 0 {
		n += 1 + l + sovContent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovContent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContent(x uint64) (n int) {
	return sovContent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ContentDelete) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContentDelete{`,
		`Digest:` + fmt.Sprintf("%v", this.Digest) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ContentDelete) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContent
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
			return fmt.Errorf("proto: ContentDelete: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentDelete: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Digest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContent
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
				return ErrInvalidLengthContent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Digest = github_com_opencontainers_go_digest.Digest(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthContent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipContent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContent
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
					return 0, ErrIntOverflowContent
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
					return 0, ErrIntOverflowContent
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
				return 0, ErrInvalidLengthContent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContent = fmt.Errorf("proto: unexpected end of group")
)
