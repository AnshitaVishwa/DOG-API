package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset Not Found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API V"))
}
