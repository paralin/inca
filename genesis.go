package inca

import (
	"github.com/aperturerobotics/pbobject"
)

// GetObjectTypeID returns the object type string, used to identify types.
func (g *Genesis) GetObjectTypeID() *pbobject.ObjectTypeID {
	return pbobject.NewObjectTypeID("/inca/genesis")
}
