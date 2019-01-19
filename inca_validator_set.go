package inca

import (
	"errors"

	"github.com/aperturerobotics/hydra/block"
	"github.com/aperturerobotics/hydra/cid"
	"github.com/golang/protobuf/proto"
)

// MarshalBlock marshals the block to binary.
// This is the initial step of marshaling, before transformations.
func (b *ValidatorSet) MarshalBlock() ([]byte, error) {
	return proto.Marshal(b)
}

// UnmarshalBlock unmarshals the block to the object.
// This is the final step of decoding, after transformations.
func (b *ValidatorSet) UnmarshalBlock(data []byte) error {
	return proto.Unmarshal(data, b)
}

// ApplyRef applies a ref change with a field id.
func (b *ValidatorSet) ApplyRef(id uint32, ptr *cid.BlockRef) error {
	return errors.New("validator set has no refs")
}

// _ is a type assertion
var _ block.Block = ((*ValidatorSet)(nil))
