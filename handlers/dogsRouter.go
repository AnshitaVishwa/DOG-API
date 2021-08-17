package handlers

import (
	"net/http"
	"strings"
)

// It will handle different type of Requests of dogs

func DogsRouter(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSuffix(path, "/")

	// Checking the type of request (Either GET/POST)

	if path == "/dogs" {
		switch r.Method {
		case http.MethodGet:
			return
		case http.MethodPost:
			return
		default:
			return
		}
	}

}
