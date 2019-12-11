package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dgravesa/WaterLogger-UserServer/model"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	qname := q.Get("name")

	if qname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, found := model.FindUserByName(qname)
	if !found {
		w.WriteHeader(http.StatusNotFound)
	} else {
		enc := json.NewEncoder(w)
		enc.Encode(user)
		w.WriteHeader(http.StatusOK)
	}
}
