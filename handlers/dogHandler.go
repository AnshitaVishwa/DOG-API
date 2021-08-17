package handlers

import (
	"encoding/json"
	"net/http"

	"../dog"
)

func GetAllDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := dog.All()
	if err != nil {
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"dogs": dogs})
}