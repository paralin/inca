package inca

import (
	"github.com/aperturerobotics/pbobject"
)

// GetObjectTypeID returns the object type string, used to identify types.
func (g *Genesis) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/genesis")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (g *NodeMessage) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/node-message")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (g *ValidatorSet) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/validator-set")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (g *ChainConfig) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/chain-config")
}

// GetObjectTypeID returns the object type string, used to identify types.
func (b *BlockHeader) GetObjectTypeID() *pbobject.ObjectTypeID {
	return &pbobject.ObjectTypeID{TypeUuid: "/inca/block-header"}
}

// GetObjectTypeID returns the object type string, used to identify types.
func (b *Block) GetObjectTypeID() *pbobject.ObjectTypeID {
	return &pbobject.ObjectTypeID{TypeUuid: "/inca/block"}
}

// GetObjectTypeID returns the object type string, used to identify types.
func (b *Vote) GetObjectTypeID() *pbobject.ObjectTypeID {
	return &pbobject.ObjectTypeID{TypeUuid: "/inca/vote"}
}
