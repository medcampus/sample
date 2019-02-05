package db

import "gopkg.in/mgo.v2/bson"

type (
	Clap struct {
		ClapId bson.ObjectId `json:"clap_id"`
		Claps int64 `json:"claps"`
	}
)
