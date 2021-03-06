// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: send.proto

package cmd0x0002

import (
	fmt "fmt"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SendGroupMsg struct {
	MsgSeq               *uint32  `protobuf:"varint,1,req,name=MsgSeq" json:"MsgSeq,omitempty"`
	Code                 *uint32  `protobuf:"varint,2,req,name=Code" json:"Code,omitempty"`
	GroupCode            *uint64  `protobuf:"varint,3,req,name=GroupCode" json:"GroupCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendGroupMsg) Reset()         { *m = SendGroupMsg{} }
func (m *SendGroupMsg) String() string { return proto.CompactTextString(m) }
func (*SendGroupMsg) ProtoMessage()    {}
func (*SendGroupMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e4a6fa8e6deebe, []int{0}
}
func (m *SendGroupMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendGroupMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendGroupMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendGroupMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendGroupMsg.Merge(m, src)
}
func (m *SendGroupMsg) XXX_Size() int {
	return m.Size()
}
func (m *SendGroupMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SendGroupMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SendGroupMsg proto.InternalMessageInfo

func (m *SendGroupMsg) GetMsgSeq() uint32 {
	if m != nil && m.MsgSeq != nil {
		return *m.MsgSeq
	}
	return 0
}

func (m *SendGroupMsg) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *SendGroupMsg) GetGroupCode() uint64 {
	if m != nil && m.GroupCode != nil {
		return *m.GroupCode
	}
	return 0
}

func init() {
	proto.RegisterType((*SendGroupMsg)(nil), "cmd0x0002.SendGroupMsg")
}

func init() { proto.RegisterFile("send.proto", fileDescriptor_27e4a6fa8e6deebe) }

var fileDescriptor_27e4a6fa8e6deebe = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4e, 0xcd, 0x4b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xce, 0x4d, 0x31, 0xa8, 0x30, 0x30, 0x30,
	0x30, 0x52, 0x8a, 0xe0, 0xe2, 0x09, 0x4e, 0xcd, 0x4b, 0x71, 0x2f, 0xca, 0x2f, 0x2d, 0xf0, 0x2d,
	0x4e, 0x17, 0x12, 0xe3, 0x62, 0xf3, 0x2d, 0x4e, 0x0f, 0x4e, 0x2d, 0x94, 0x60, 0x54, 0x60, 0xd2,
	0xe0, 0x0d, 0x82, 0xf2, 0x84, 0x84, 0xb8, 0x58, 0x9c, 0xf3, 0x53, 0x52, 0x25, 0x98, 0xc0, 0xa2,
	0x60, 0xb6, 0x90, 0x0c, 0x17, 0x27, 0x58, 0x1f, 0x58, 0x82, 0x59, 0x81, 0x49, 0x83, 0x25, 0x08,
	0x21, 0xe0, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0xce, 0x78, 0x2c, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xde, 0x9f, 0xa8, 0x83, 0x00,
	0x00, 0x00,
}

func (m *SendGroupMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendGroupMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendGroupMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.GroupCode == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintSend(dAtA, i, uint64(*m.GroupCode))
		i--
		dAtA[i] = 0x18
	}
	if m.Code == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintSend(dAtA, i, uint64(*m.Code))
		i--
		dAtA[i] = 0x10
	}
	if m.MsgSeq == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintSend(dAtA, i, uint64(*m.MsgSeq))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSend(dAtA []byte, offset int, v uint64) int {
	offset -= sovSend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendGroupMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgSeq != nil {
		n += 1 + sovSend(uint64(*m.MsgSeq))
	}
	if m.Code != nil {
		n += 1 + sovSend(uint64(*m.Code))
	}
	if m.GroupCode != nil {
		n += 1 + sovSend(uint64(*m.GroupCode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSend(x uint64) (n int) {
	return sovSend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendGroupMsg) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSend
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
			return fmt.Errorf("proto: SendGroupMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendGroupMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgSeq", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MsgSeq = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Code = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupCode", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.GroupCode = &v
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipSend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSend
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
					return 0, ErrIntOverflowSend
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
					return 0, ErrIntOverflowSend
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
				return 0, ErrInvalidLengthSend
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSend
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSend
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSend        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSend          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSend = fmt.Errorf("proto: unexpected end of group")
)
