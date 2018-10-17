// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/inca/peer/peer.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import inca "github.com/aperturerobotics/inca"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PeerState is information about a peer stored in the database.
type PeerState struct {
	// LastObservedMessage is the last observed message from the peer.
	LastObservedMessage  *inca.NodeMessage `protobuf:"bytes,1,opt,name=last_observed_message,json=lastObservedMessage" json:"last_observed_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PeerState) Reset()         { *m = PeerState{} }
func (m *PeerState) String() string { return proto.CompactTextString(m) }
func (*PeerState) ProtoMessage()    {}
func (*PeerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_peer_2a7861f65e14dd24, []int{0}
}
func (m *PeerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerState.Unmarshal(m, b)
}
func (m *PeerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerState.Marshal(b, m, deterministic)
}
func (dst *PeerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerState.Merge(dst, src)
}
func (m *PeerState) XXX_Size() int {
	return xxx_messageInfo_PeerState.Size(m)
}
func (m *PeerState) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerState.DiscardUnknown(m)
}

var xxx_messageInfo_PeerState proto.InternalMessageInfo

func (m *PeerState) GetLastObservedMessage() *inca.NodeMessage {
	if m != nil {
		return m.LastObservedMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerState)(nil), "peer.PeerState")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/inca/peer/peer.proto", fileDescriptor_peer_2a7861f65e14dd24)
}

var fileDescriptor_peer_2a7861f65e14dd24 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x48, 0x2d, 0x2a, 0x29, 0x2d, 0x4a, 0x2d,
	0xca, 0x4f, 0xca, 0x2f, 0xc9, 0x4c, 0x2e, 0xd6, 0xcf, 0xcc, 0x4b, 0x4e, 0xd4, 0x2f, 0x48, 0x4d,
	0x2d, 0x02, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x94, 0x36, 0x41,
	0x7d, 0x20, 0x02, 0xa2, 0x45, 0x29, 0x88, 0x8b, 0x33, 0x20, 0x35, 0xb5, 0x28, 0xb8, 0x24, 0xb1,
	0x24, 0x55, 0xc8, 0x95, 0x4b, 0x34, 0x27, 0xb1, 0xb8, 0x24, 0x3e, 0x3f, 0xa9, 0x38, 0xb5, 0xa8,
	0x2c, 0x35, 0x25, 0x3e, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x50, 0x0f, 0xac, 0xd1, 0x2f, 0x3f, 0x25, 0xd5, 0x17, 0x22, 0x11, 0x24, 0x0c, 0x52,
	0xef, 0x0f, 0x55, 0x0e, 0x15, 0x4c, 0x62, 0x03, 0x1b, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x97, 0xde, 0xfa, 0xc1, 0x00, 0x00, 0x00,
}
