package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type Wheel struct {
	ID           objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Manufacturer string            `json:"manufacturer" bson:"manufacturer"`
	Model        string            `json:"model" bson:"model"`
	Offset       int               `json:"offset" bson:"offset"`
	Diameter     int               `json:"diameter" bson:"diameter"`
}
