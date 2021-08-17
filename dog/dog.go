package dog

// go get github.com/asdine/storm
// Storm is an advanced query system, storm provides indexes,
// a wide range of methods to store and fetch data
// `storm:"index"` // this field will be indexed

//go get gopkg.in/mgo.v2/bson
// bson.ObjectId // ObjectId is a unique ID identifying a BSON value.
// It must be exactly 12 bytes long.
// MongoDB objects by default have such a property set in their "_id" property.

import "gopkg.in/mgo.v2/bson"

// Creating the data structure of our dogs info

type Dog struct {

	// First Letter is Capital as it has to be exported

	ID    bson.ObjectId `json: "id" storm: "id"` // Fields appear in json will be like these
	Name  string        `json:"name"`
	Breed string        `json:"breed"`
}
