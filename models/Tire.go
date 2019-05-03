package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type Tire struct {
	ID           objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Manufacturer string            `json:"manufacturer" bson:"manufacturer"`
	Model        string            `json:"model" bson:"model"`
	Width        int               `json:"width" bson:"width"`
	AspectRatio  int               `json:"aspectRatio" bson:"aspectRatio"`
	Diameter     int               `json:"diameter" bson:"diameter"`
}
