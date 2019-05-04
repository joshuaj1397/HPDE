package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type Brakes struct {
	ID           objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Manufacturer string            `json:"manufacturer" bson:"manufacturer"`
	Model        string            `json:"model" bson:"model"`
	Rotors       struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
		Slotted      bool   `json:"slotted" bson:"slotted"`
		Drilled      bool   `json:"drilled" bson:"drilled"`
		Vented       bool   `json:"vented" bson:"vented"`
	} `json:"rotors" bson:"rotors"`
	Pads struct {
		Manufacturer string `json:"manufacturer" bson:"manufacturer"`
		Model        string `json:"model" bson:"model"`
	} `json:"pads" bson:"pads"`
}
