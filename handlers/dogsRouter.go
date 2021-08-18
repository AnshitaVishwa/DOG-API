package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// It will handle different type of Requests of dogs

func DogsRouter(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	path = strings.TrimSuffix(path, "/")

	// Checking the type of request (Either GET/POST)

	if path == "/dogs" {
		switch r.Method {
		case http.MethodGet:
			getAllDogs(w, r)
			return
		case http.MethodPost:
			dogsPostOne(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	// Adding CRUD functionalities to Route

	path = strings.TrimPrefix(path, "/dogs/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusMethodNotAllowed)
		return
	}

	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		dogsGetOne(w, r, id)
		return
	case http.MethodPut:
		dogsPutOne(w, r, id)
		return
	case http.MethodPatch:
		// usersPatchOne(w, r, id)
		return
	case http.MethodDelete:
		// usersDeleteOne(w, r, id)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}

}
