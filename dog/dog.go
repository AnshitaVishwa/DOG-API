package dog

// go get github.com/asdine/storm
// Storm is an advanced query system, storm provides indexes,
// a wide range of methods to store and fetch data
// `storm:"index"` // this field will be indexed

//go get gopkg.in/mgo.v2/bson
// bson.ObjectId // ObjectId is a unique ID identifying a BSON value.
// It must be exactly 12 bytes long.
// MongoDB objects by default have such a property set in their "_id" property.

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

// Creating the data structure of our dogs info

type Dog struct {

	// First Letter is Capital as it has to be exported

	ID    bson.ObjectId `json:"id" storm:"id"` // Fields appear in json will be like these
	Name  string        `json:"name"`
	Breed string        `json:"breed"`
}

// connect the server with the database
// these are global variables hence they are to be initialized as follows

// this will create dogs.db file if not present or if present access the data from there
const (
	dbPath = "dogs.db"
)

var (
	ErrRecordInvalid = errors.New("record is in valid")
)

// Writing Functions to retrieve data from the database DOG

func All() ([]Dog, error) {
	// open the file / database
	db, err := storm.Open(dbPath)

	if err != nil {
		return nil, err
	}
	defer db.Close()
	dogs := []Dog{}
	err = db.All(&dogs)
	if err != nil {
		return nil, err
	}
	return dogs, nil
}

func One(id bson.ObjectId) (*Dog, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	d := new(Dog)
	err = db.One("ID", id, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// validate makes sure that the record contains v

func (d *Dog) validate() error {
	if d.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}

// save updates or creates a given record in the database

func (d *Dog) Save() error {
	if err := d.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(d)
}

// Delete removes a given record from the database
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	d := new(Dog)
	err = db.One("ID", id, d)
	if err != nil {
		return err
	}
	return db.DeleteStruct(d)
}
