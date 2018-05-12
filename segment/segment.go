package segment

import (
	"github.com/aperturerobotics/pbobject"
)

// GetObjectTypeID returns the object type string, used to identify types.
func (s *SegmentState) GetObjectTypeID() *pbobject.ObjectTypeID {
	return &pbobject.ObjectTypeID{TypeUuid: "/inca/segment"}
}
