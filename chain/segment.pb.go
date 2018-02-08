// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/inca/chain/segment.proto

package chain

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import storageref "github.com/aperturerobotics/storageref"
import inca "github.com/aperturerobotics/inca"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SegmentStatus is the status of a segment.
type SegmentStatus int32

const (
	// INVALID indicates the segment has been shown to be invalid.
	SegmentStatus_SegmentStatus_INVALID SegmentStatus = 0
	// DISJOINTED indicates the segment has not been connected to any trusted chain yet.
	SegmentStatus_SegmentStatus_DISJOINTED SegmentStatus = 1
	// VALID indicates the segment has been shown to be valid.
	SegmentStatus_SegmentStatus_VALID SegmentStatus = 2
)

var SegmentStatus_name = map[int32]string{
	0: "SegmentStatus_INVALID",
	1: "SegmentStatus_DISJOINTED",
	2: "SegmentStatus_VALID",
}
var SegmentStatus_value = map[string]int32{
	"SegmentStatus_INVALID":    0,
	"SegmentStatus_DISJOINTED": 1,
	"SegmentStatus_VALID":      2,
}

func (x SegmentStatus) String() string {
	return proto.EnumName(SegmentStatus_name, int32(x))
}
func (SegmentStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// SegmentState tracks state of a Segment.
type SegmentState struct {
	// Id is the identifier of the segment.
	// Usually a UUID
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Status is the status of the segment.
	Status SegmentStatus `protobuf:"varint,2,opt,name=status,enum=chain.SegmentStatus" json:"status,omitempty"`
	// HeadBlock is a reference to the HEAD of this segment.
	HeadBlock *storageref.StorageRef `protobuf:"bytes,3,opt,name=head_block,json=headBlock" json:"head_block,omitempty"`
	// TailBlock is a reference to the TAIL of this segment.
	TailBlock *storageref.StorageRef `protobuf:"bytes,4,opt,name=tail_block,json=tailBlock" json:"tail_block,omitempty"`
	// SegmentPrev is the previous segment.
	SegmentPrev string `protobuf:"bytes,5,opt,name=segment_prev,json=segmentPrev" json:"segment_prev,omitempty"`
	// SegmentNext is the next segment.
	SegmentNext string `protobuf:"bytes,6,opt,name=segment_next,json=segmentNext" json:"segment_next,omitempty"`
	// TailBlockRound is the block round info of the tail.
	TailBlockRound *inca.BlockRoundInfo `protobuf:"bytes,7,opt,name=tail_block_round,json=tailBlockRound" json:"tail_block_round,omitempty"`
}

func (m *SegmentState) Reset()                    { *m = SegmentState{} }
func (m *SegmentState) String() string            { return proto.CompactTextString(m) }
func (*SegmentState) ProtoMessage()               {}
func (*SegmentState) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *SegmentState) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SegmentState) GetStatus() SegmentStatus {
	if m != nil {
		return m.Status
	}
	return SegmentStatus_SegmentStatus_INVALID
}

func (m *SegmentState) GetHeadBlock() *storageref.StorageRef {
	if m != nil {
		return m.HeadBlock
	}
	return nil
}

func (m *SegmentState) GetTailBlock() *storageref.StorageRef {
	if m != nil {
		return m.TailBlock
	}
	return nil
}

func (m *SegmentState) GetSegmentPrev() string {
	if m != nil {
		return m.SegmentPrev
	}
	return ""
}

func (m *SegmentState) GetSegmentNext() string {
	if m != nil {
		return m.SegmentNext
	}
	return ""
}

func (m *SegmentState) GetTailBlockRound() *inca.BlockRoundInfo {
	if m != nil {
		return m.TailBlockRound
	}
	return nil
}

func init() {
	proto.RegisterType((*SegmentState)(nil), "chain.SegmentState")
	proto.RegisterEnum("chain.SegmentStatus", SegmentStatus_name, SegmentStatus_value)
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/inca/chain/segment.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0xb4, 0x95, 0x6e, 0x6b, 0x29, 0x6b, 0xd5, 0xb5, 0x78, 0xa8, 0x9e, 0x8a, 0x4a,
	0x02, 0x55, 0xf1, 0x26, 0x28, 0xf5, 0x10, 0x91, 0x2a, 0x89, 0x78, 0x0d, 0x9b, 0x64, 0xda, 0x2e,
	0xb6, 0xd9, 0xb0, 0x99, 0x94, 0xfe, 0x3c, 0x7f, 0x9a, 0x64, 0x13, 0x34, 0xb9, 0xd4, 0xcb, 0x32,
	0xbc, 0x79, 0x1f, 0xef, 0x31, 0x4b, 0x6e, 0xe7, 0x02, 0x17, 0x59, 0x60, 0x85, 0x72, 0x65, 0xf3,
	0x04, 0x14, 0x66, 0x0a, 0x94, 0x0c, 0x24, 0x8a, 0x30, 0xb5, 0x45, 0x1c, 0x72, 0x3b, 0x5c, 0x70,
	0x11, 0xdb, 0x29, 0xcc, 0x57, 0x10, 0xa3, 0x95, 0x28, 0x89, 0x92, 0x36, 0xb4, 0x38, 0xb8, 0xdf,
	0x06, 0xa7, 0x28, 0x15, 0x9f, 0x83, 0x82, 0x59, 0x65, 0x2c, 0xf8, 0xc1, 0xd5, 0xbf, 0xa9, 0xf9,
	0x53, 0x98, 0x2f, 0xbe, 0x4d, 0xd2, 0xf1, 0x8a, 0x78, 0x0f, 0x39, 0x02, 0xed, 0x12, 0x53, 0x44,
	0xcc, 0x18, 0x1a, 0xa3, 0x96, 0x6b, 0x8a, 0x88, 0x5e, 0x93, 0x66, 0x8a, 0x1c, 0xb3, 0x94, 0x99,
	0x43, 0x63, 0xd4, 0x1d, 0xf7, 0x2d, 0x5d, 0xcf, 0xaa, 0x40, 0x59, 0xea, 0x96, 0x1e, 0x7a, 0x47,
	0xc8, 0x02, 0x78, 0xe4, 0x07, 0x4b, 0x19, 0x7e, 0xb1, 0xdd, 0xa1, 0x31, 0x6a, 0x8f, 0x8f, 0xad,
	0x4a, 0x45, 0xaf, 0x18, 0x5d, 0x98, 0xb9, 0xad, 0xdc, 0xf9, 0x94, 0x1b, 0x73, 0x0c, 0xb9, 0x58,
	0x96, 0xd8, 0xde, 0x76, 0x2c, 0x77, 0x16, 0xd8, 0x39, 0xe9, 0x94, 0xa7, 0xf3, 0x13, 0x05, 0x6b,
	0xd6, 0xd0, 0xad, 0xdb, 0xa5, 0xf6, 0xae, 0x60, 0x5d, 0xb5, 0xc4, 0xb0, 0x41, 0xd6, 0xac, 0x59,
	0xa6, 0xb0, 0x41, 0xfa, 0x40, 0x7a, 0x7f, 0xe1, 0xbe, 0x92, 0x59, 0x1c, 0xb1, 0x7d, 0x5d, 0xa1,
	0x6f, 0xe9, 0x4b, 0xe9, 0x30, 0x37, 0xd7, 0x9d, 0x78, 0x26, 0xdd, 0xee, 0x6f, 0x01, 0xad, 0x5d,
	0x72, 0x72, 0x50, 0x3b, 0x06, 0x3d, 0x25, 0x47, 0x35, 0xc1, 0x77, 0xa6, 0x9f, 0x8f, 0xaf, 0xce,
	0xa4, 0xb7, 0x43, 0xcf, 0x08, 0xab, 0xaf, 0x26, 0x8e, 0xf7, 0xf2, 0xe6, 0x4c, 0x3f, 0x9e, 0x27,
	0x3d, 0x83, 0x9e, 0x90, 0xc3, 0xfa, 0xb6, 0xc0, 0xcc, 0xa0, 0xa9, 0x3f, 0xeb, 0xe6, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0xd5, 0x7d, 0x56, 0x51, 0x02, 0x00, 0x00,
}