// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/inca/block/block.proto

/*
Package block is a generated protocol buffer package.

It is generated from these files:
	github.com/aperturerobotics/inca/block/block.proto

It has these top-level messages:
	State
*/
package block

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import storageref "github.com/aperturerobotics/storageref"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// State contains state for a block.
type State struct {
	// NextBlock is the known next block in the sequence.
	NextBlock *storageref.StorageRef `protobuf:"bytes,1,opt,name=next_block,json=nextBlock" json:"next_block,omitempty"`
	// SegmentId is the ID of the block segment.
	SegmentId string `protobuf:"bytes,2,opt,name=segment_id,json=segmentId" json:"segment_id,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *State) GetNextBlock() *storageref.StorageRef {
	if m != nil {
		return m.NextBlock
	}
	return nil
}

func (m *State) GetSegmentId() string {
	if m != nil {
		return m.SegmentId
	}
	return ""
}

func init() {
	proto.RegisterType((*State)(nil), "block.State")
}

func init() { proto.RegisterFile("github.com/aperturerobotics/inca/block/block.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x48, 0x2d, 0x2a, 0x29, 0x2d, 0x4a, 0x2d,
	0xca, 0x4f, 0xca, 0x2f, 0xc9, 0x4c, 0x2e, 0xd6, 0xcf, 0xcc, 0x4b, 0x4e, 0xd4, 0x4f, 0xca, 0xc9,
	0x4f, 0xce, 0x86, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x94, 0x39,
	0x3e, 0xad, 0xc5, 0x25, 0xf9, 0x45, 0x89, 0xe9, 0xa9, 0x45, 0xa9, 0x69, 0x48, 0x4c, 0x88, 0x7e,
	0xa5, 0x58, 0x2e, 0xd6, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x53, 0x2e, 0xae, 0xbc, 0xd4, 0x8a,
	0x92, 0x78, 0xb0, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x62, 0x7a, 0x48, 0xea, 0x83,
	0x21, 0xcc, 0xa0, 0xd4, 0xb4, 0x20, 0x4e, 0x90, 0x4a, 0x27, 0x90, 0x42, 0x21, 0x59, 0x2e, 0xae,
	0xe2, 0xd4, 0xf4, 0xdc, 0xd4, 0xbc, 0x92, 0xf8, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x4e, 0xa8, 0x88, 0x67, 0x4a, 0x12, 0x1b, 0xd8, 0x16, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x97, 0x0a, 0xdc, 0x0b, 0xdb, 0x00, 0x00, 0x00,
}
