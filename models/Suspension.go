package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type Suspension struct {
	ID           objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Manufacturer string            `json:"manufacturer" bson:"manufacturer"`
	Model        string            `json:"model" bson:"model"`
	RideHeight   float32           `json:"rideHeight" bson:"rideHeight"`
	Camber       float32           `json:"camber" bson:"camber"`
	Damping      int               `json:"damping" bson:"damping"`
}
