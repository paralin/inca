// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/inca/chain/chain.proto

/*
Package chain is a generated protocol buffer package.

It is generated from these files:
	github.com/aperturerobotics/inca/chain/chain.proto
	github.com/aperturerobotics/inca/chain/chain_config.proto
	github.com/aperturerobotics/inca/chain/proposer.proto
	github.com/aperturerobotics/inca/chain/segment.proto
	github.com/aperturerobotics/inca/chain/validator.proto

It has these top-level messages:
	ChainState
	Config
	ProposerState
	SegmentState
	ValidatorState
*/
package chain

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/aperturerobotics/storageref"

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
	LastRound uint64 `protobuf:"varint,3,opt,name=last_round,json=lastRound" json:"last_round,omitempty"`
}

func (m *ChainState) Reset()                    { *m = ChainState{} }
func (m *ChainState) String() string            { return proto.CompactTextString(m) }
func (*ChainState) ProtoMessage()               {}
func (*ChainState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func init() { proto.RegisterFile("github.com/aperturerobotics/inca/chain/chain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xcd, 0xbf, 0xca, 0xc2, 0x30,
	0x14, 0x05, 0x70, 0xfa, 0xfd, 0x83, 0xe6, 0xd3, 0x25, 0x53, 0x11, 0xc4, 0xa2, 0x4b, 0xa7, 0x16,
	0x74, 0xf0, 0x01, 0x5c, 0x9c, 0xdb, 0x07, 0x28, 0x69, 0xbc, 0x26, 0x01, 0x9b, 0x5b, 0x6f, 0x6e,
	0xdf, 0x5f, 0x92, 0x2e, 0x4e, 0x2e, 0xe1, 0x9c, 0x5f, 0x38, 0x89, 0x38, 0x1a, 0xc7, 0x76, 0x1e,
	0x6a, 0x8d, 0x63, 0xa3, 0x26, 0x20, 0x9e, 0x09, 0x08, 0x07, 0x64, 0xa7, 0x43, 0xe3, 0xbc, 0x56,
	0x8d, 0xb6, 0xca, 0xf9, 0xe5, 0xac, 0x27, 0x42, 0x46, 0xf9, 0x9b, 0xca, 0xe6, 0xfc, 0x69, 0x1a,
	0x18, 0x49, 0x19, 0x20, 0xb8, 0xbf, 0xc5, 0x65, 0xbf, 0x7f, 0x0a, 0x71, 0x89, 0x2f, 0x74, 0xac,
	0x18, 0xe4, 0x41, 0xac, 0x43, 0x0c, 0x7d, 0x00, 0x33, 0x82, 0xe7, 0x22, 0x2b, 0xb3, 0x2a, 0x6f,
	0x57, 0x09, 0xbb, 0xc5, 0xe4, 0x4e, 0xfc, 0x3f, 0x54, 0xe0, 0xde, 0x82, 0x33, 0x96, 0x8b, 0xaf,
	0x32, 0xab, 0x7e, 0x5a, 0x11, 0xe9, 0x9a, 0x44, 0x6e, 0x45, 0x6a, 0x3d, 0xe1, 0xec, 0x6f, 0xc5,
	0x77, 0xba, 0xcf, 0xa3, 0xb4, 0x11, 0x86, 0xbf, 0xf4, 0xf3, 0xe9, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x0e, 0x06, 0xa1, 0x93, 0xef, 0x00, 0x00, 0x00,
}
