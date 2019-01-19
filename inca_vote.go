package inca

import (
	"errors"

	"github.com/aperturerobotics/bifrost/hash"
	"github.com/aperturerobotics/bifrost/peer"
	"github.com/aperturerobotics/hydra/block"
	"github.com/aperturerobotics/hydra/cid"
	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-crypto"
)

// MarshalBlock marshals the block to binary.
// This is the initial step of marshaling, before transformations.
func (b *Vote) MarshalBlock() ([]byte, error) {
	return proto.Marshal(b)
}

// UnmarshalBlock unmarshals the block to the object.
// This is the final step of decoding, after transformations.
func (b *Vote) UnmarshalBlock(data []byte) error {
	return proto.Unmarshal(data, b)
}

// ApplyRef applies a ref change with a field id.
func (b *Vote) ApplyRef(id uint32, ptr *cid.BlockRef) error {
	return errors.New("vote has no refs")
}

// SignBlockHeader signs the block header and stores the data in the Vote.
func (b *Vote) SignBlockHeader(privKey crypto.PrivKey, bhRef *cid.BlockRef) error {
	if privKey == nil {
		return errors.New("priv key cannot be empty")
	}

	// marshal the bhRef
	dat, err := bhRef.MarshalKey()
	if err != nil {
		return err
	}

	bhSig, err := peer.NewSignature(
		privKey,
		hash.HashType_HashType_SHA256,
		dat,
		false,
	)
	if err != nil {
		return err
	}
	b.ValidatorSignature = bhSig
	return nil
}

// _ is a type assertion
var _ block.Block = ((*Vote)(nil))
