package inca

import (
	"github.com/aperturerobotics/hydra/block"
	"github.com/aperturerobotics/hydra/cid"
	"github.com/golang/protobuf/proto"
)

// MarshalBlock marshals the block to binary.
// This is the initial step of marshaling, before transformations.
func (b *Block) MarshalBlock() ([]byte, error) {
	return proto.Marshal(b)
}

// UnmarshalBlock unmarshals the block to the object.
// This is the final step of decoding, after transformations.
func (b *Block) UnmarshalBlock(data []byte) error {
	return proto.Unmarshal(data, b)
}

// ApplyRef applies a ref change with a field id.
// The reference may be nil if the child block is nil.
func (b *Block) ApplyRef(id uint32, ptr *cid.BlockRef) error {
	switch id {
	case 1:
		b.BlockHeaderRef = ptr
	case 2:
		if ptr.GetEmpty() {
			b.VoteRefs = nil
		} else {
			b.VoteRefs = append(b.VoteRefs, ptr)
		}
	}
	return nil
}

// _ is a type assertion
var _ block.Block = ((*Block)(nil))
