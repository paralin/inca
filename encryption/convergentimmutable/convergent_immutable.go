package convergentimmutable

import (
	"github.com/aperturerobotics/pbobject"
)

// configTypeID is the type ID of the configuration.
var configTypeID = &pbobject.ObjectTypeID{TypeUuid: "/inca/encryption/convergent-immutable/config"}

// GetObjectTypeID returns the object type string, used to identify types.
func (c *Config) GetObjectTypeID() *pbobject.ObjectTypeID {
	return configTypeID
}
