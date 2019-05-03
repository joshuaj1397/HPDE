package model

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type Car struct {
	ID           objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Manufacturer string            `json:"manufacturer" bson:"manufacturer"`
	Model        string            `json:"model" bson:"model"`
	Year         int               `json:"year" bson:"year"`
	Image        string            `json:"image" bson:"image"`
	Mods         struct {
		Tires struct {
			FrontLeft  objectid.ObjectID `json:"frontLeft" bson:"frontLeft"`
			FrontRight objectid.ObjectID `json:"frontRight" bson:"frontRight"`
			RearLeft   objectid.ObjectID `json:"rearLeft" bson:"rearLeft"`
			RearRight  objectid.ObjectID `json:"rearRight" bson:"rearRight"`
		} `json:"tires" bson:"tires"`
		Wheels struct {
			FrontLeft  objectid.ObjectID `json:"frontLeft" bson:"frontLeft"`
			FrontRight objectid.ObjectID `json:"frontRight" bson:"frontRight"`
			RearLeft   objectid.ObjectID `json:"rearLeft" bson:"rearLeft"`
			RearRight  objectid.ObjectID `json:"rearRight" bson:"rearRight"`
		} `json:"wheels" bson:"wheels"`
	} `json:"mods" bson:"mods"`
}
