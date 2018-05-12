package chain

import (
	"github.com/aperturerobotics/pbobject"
)

// GetObjectTypeID returns the object type string, used to identify types.
func (g *ChainState) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/chain-state")
}
