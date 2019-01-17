// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/inca/chain/chain.proto

package chain

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ChainState manages the chain state in the database.
type ChainState struct {
	// StateSegment is the current active segment used for state.
	StateSegment string `protobuf:"bytes,1,opt,name=state_segment,json=stateSegment" json:"state_segment,omitempty"`
	// LastHeight is the last processed height.
	LastHeight uint64 `protobuf:"varint,2,opt,name=last_height,json=lastHeight" json:"last_height,omitempty"`
	// LastRound is the last processed round.
	LastRound            uint64   `protobuf:"varint,3,opt,name=last_round,json=lastRound" json:"last_round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainState) Reset()         { *m = ChainState{} }
func (m *ChainState) String() string { return proto.CompactTextString(m) }
func (*ChainState) ProtoMessage()    {}
func (*ChainState) Descriptor() ([]byte, []int) {
	return fileDescriptor_chain_005209ff69c5a025, []int{0}
}
func (m *ChainState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainState.Unmarshal(m, b)
}
func (m *ChainState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainState.Marshal(b, m, deterministic)
}
func (dst *ChainState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainState.Merge(dst, src)
}
func (m *ChainState) XXX_Size() int {
	return xxx_messageInfo_ChainState.Size(m)
}
func (m *ChainState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainState.DiscardUnknown(m)
}

var xxx_messageInfo_ChainState proto.InternalMessageInfo

func (m *ChainState) GetStateSegment() string {
	if m != nil {
		return m.StateSegment
	}
	return ""
}

func (m *ChainState) GetLastHeight() uint64 {
	if m != nil {
		return m.LastHeight
	}
	return 0
}

func (m *ChainState) GetLastRound() uint64 {
	if m != nil {
		return m.LastRound
	}
	return 0
}

func init() {
	proto.RegisterType((*ChainState)(nil), "chain.ChainState")
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/inca/chain/chain.proto", fileDescriptor_chain_005209ff69c5a025)
}

var fileDescriptor_chain_005209ff69c5a025 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xce, 0xbf, 0x0e, 0x82, 0x30,
	0x10, 0x06, 0xf0, 0xe0, 0xbf, 0x84, 0x53, 0x97, 0x4e, 0x2c, 0x46, 0xa2, 0x0b, 0x13, 0x24, 0xfa,
	0x08, 0x2e, 0xce, 0xf0, 0x00, 0xa4, 0xd4, 0x0b, 0x6d, 0x22, 0x3d, 0x6c, 0x8f, 0xf7, 0x37, 0x3d,
	0x96, 0xcb, 0xf7, 0xfd, 0x6e, 0xf9, 0xe0, 0x31, 0x3a, 0xb6, 0xcb, 0x50, 0x1b, 0x9a, 0x1a, 0x3d,
	0x63, 0xe0, 0x25, 0x60, 0xa0, 0x81, 0xd8, 0x99, 0xd8, 0x38, 0x6f, 0x74, 0x63, 0xac, 0x76, 0x7e,
	0xbd, 0xf5, 0x1c, 0x88, 0x49, 0xed, 0xa5, 0xdc, 0x7e, 0x00, 0xaf, 0x14, 0x3a, 0xd6, 0x8c, 0xea,
	0x0e, 0xe7, 0x98, 0x42, 0x1f, 0x71, 0x9c, 0xd0, 0x73, 0x91, 0x95, 0x59, 0x95, 0xb7, 0x27, 0xc1,
	0x6e, 0x35, 0x75, 0x85, 0xe3, 0x57, 0x47, 0xee, 0x2d, 0xba, 0xd1, 0x72, 0xb1, 0x29, 0xb3, 0x6a,
	0xd7, 0x42, 0xa2, 0xb7, 0x88, 0xba, 0x80, 0xb4, 0x3e, 0xd0, 0xe2, 0x3f, 0xc5, 0x56, 0xfe, 0x79,
	0x92, 0x36, 0xc1, 0x70, 0x90, 0x01, 0xcf, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xb5, 0x64,
	0x5f, 0xb6, 0x00, 0x00, 0x00,
}
