package controller

import "net/http"

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		postUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	case http.MethodPut:
		fallthrough // TODO implement
		// putUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
