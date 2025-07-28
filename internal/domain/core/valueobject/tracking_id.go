package valueobject

import (
	"github.com/google/uuid"
	sharedVO "github.com/leninner/shared/domain/valueobject"
)

type TrackingID struct {
	sharedVO.WithID[uuid.UUID]
}

func NewTrackingID() TrackingID {
	return TrackingID{WithID: sharedVO.WithID[uuid.UUID]{ID: uuid.New()}}
}

func TrackingIDFromUUID(id uuid.UUID) TrackingID {
	return TrackingID{WithID: sharedVO.WithID[uuid.UUID]{ID: id}}
}
