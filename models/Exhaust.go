package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type Exhaust struct {
	ID      objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Headers struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
	} `json:"exhaust" bson:"exhaust"`
	Catback struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
	} `json:"catback" bson:"catback"`
	Overpipe struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
	} `json:"overpipe" bson:"overpipe"`
	Frontpipe struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
	} `json:"frontpipe" bson:"frontpipe"`
	Midpipe struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
	} `json:"midpipe" bson:"midpipe"`
	Muffler struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
	} `json:"muffler" bson:"muffler"`
}
