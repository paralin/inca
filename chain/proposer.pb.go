// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/inca/chain/proposer.proto

package chain

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

// ProposerState is state for a proposer.
type ProposerState struct {
	LastProposal         *inca.BlockRoundInfo `protobuf:"bytes,1,opt,name=last_proposal,json=lastProposal" json:"last_proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProposerState) Reset()         { *m = ProposerState{} }
func (m *ProposerState) String() string { return proto.CompactTextString(m) }
func (*ProposerState) ProtoMessage()    {}
func (*ProposerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposer_a3df4e6fda520661, []int{0}
}
func (m *ProposerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposerState.Unmarshal(m, b)
}
func (m *ProposerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposerState.Marshal(b, m, deterministic)
}
func (dst *ProposerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposerState.Merge(dst, src)
}
func (m *ProposerState) XXX_Size() int {
	return xxx_messageInfo_ProposerState.Size(m)
}
func (m *ProposerState) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposerState.DiscardUnknown(m)
}

var xxx_messageInfo_ProposerState proto.InternalMessageInfo

func (m *ProposerState) GetLastProposal() *inca.BlockRoundInfo {
	if m != nil {
		return m.LastProposal
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposerState)(nil), "chain.ProposerState")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/inca/chain/proposer.proto", fileDescriptor_proposer_a3df4e6fda520661)
}

var fileDescriptor_proposer_a3df4e6fda520661 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x48, 0x2d, 0x2a, 0x29, 0x2d, 0x4a, 0x2d,
	0xca, 0x4f, 0xca, 0x2f, 0xc9, 0x4c, 0x2e, 0xd6, 0xcf, 0xcc, 0x4b, 0x4e, 0xd4, 0x4f, 0xce, 0x48,
	0xcc, 0xcc, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x8b, 0x4a, 0x69, 0x13, 0xd4, 0x0d, 0x22, 0x20, 0x7a, 0x94, 0xbc, 0xb8,
	0x78, 0x03, 0xa0, 0xa6, 0x04, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x59, 0x72, 0xf1, 0xe6, 0x24, 0x16,
	0x97, 0xc4, 0x43, 0xcc, 0x4e, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd1, 0x03,
	0x6b, 0x72, 0xca, 0xc9, 0x4f, 0xce, 0x0e, 0xca, 0x2f, 0xcd, 0x4b, 0xf1, 0xcc, 0x4b, 0xcb, 0x0f,
	0xe2, 0x01, 0x29, 0x0d, 0x80, 0xaa, 0x4c, 0x62, 0x03, 0x1b, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xec, 0x94, 0x4e, 0x8d, 0xbf, 0x00, 0x00, 0x00,
}
