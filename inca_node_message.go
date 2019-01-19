package inca

import (
	"github.com/aperturerobotics/hydra/block"
	"github.com/aperturerobotics/hydra/cid"
	"github.com/golang/protobuf/proto"
)

// MarshalBlock marshals the block to binary.
// This is the initial step of marshaling, before transformations.
func (b *NodeMessage) MarshalBlock() ([]byte, error) {
	return proto.Marshal(b)
}

// UnmarshalBlock unmarshals the block to the object.
// This is the final step of decoding, after transformations.
func (b *NodeMessage) UnmarshalBlock(data []byte) error {
	return proto.Unmarshal(data, b)
}

// ApplyRef applies a ref change with a field id.
func (b *NodeMessage) ApplyRef(id uint32, ptr *cid.BlockRef) error {
	switch id {
	case 2:
		b.GenesisRef = ptr
	case 3:
		b.PrevMsgRef = ptr
	case 5:
		b.InnerRef = ptr
	}

	return nil
}

// _ is a type assertion
var _ block.Block = ((*NodeMessage)(nil))
