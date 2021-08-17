package handlers

import (
	"net/http"

	"../dog"
)

func GetAllDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := dog.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"dogs": dogs})
}
