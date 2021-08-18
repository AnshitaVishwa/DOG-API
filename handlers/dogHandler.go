package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"../dog"
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func getAllDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := dog.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"dogs": dogs})
}

func bodyToDog(r *http.Request, d *dog.Dog) error {
	if r.Body == nil {
		return errors.New("Requests body is empty")
	}
	if d == nil {
		return errors.New("a dog is required")
	}
	bd, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}
	return json.Unmarshal(bd, d)
}

func dogsPostOne(w http.ResponseWriter, r *http.Request) {
	d := new(dog.Dog)
	err := bodyToDog(r, d)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	//  Assigning an ID to the dog
	d.ID = bson.NewObjectId()
	err = d.Save()
	if err != nil {
		if err == dog.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/dogs/"+d.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

func dogsGetOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	d, err := dog.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"dogs": d})
}

func dogsPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	d := new(dog.Dog)
	err := bodyToDog(r, d)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	d.ID = id
	err = d.Save()
	if err != nil {
		if err == dog.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": d})
}

func dogsPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	d, err := dog.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToDog(r, d)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	d.ID = id
	err = d.Save()
	if err != nil {
		if err == dog.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": d})
}
