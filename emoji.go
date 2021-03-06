package main

import (
	"gopkg.in/mgo.v2/bson"
	//"io/ioutil"
	//"fmt"
	//"strings"
)

type Emoji struct {
	Id    bson.ObjectId `bson:"_id"`
	Emoji string        `bson:"emoji,omitempty"`
	Count int           `bson:"count"`
}