package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgravesa/WaterLogger-UserServer/model"
)

func postUser(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	qname := q.Get("name")

	if qname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, existing := model.FindUserByName(qname); existing {
		w.WriteHeader(http.StatusConflict)
		return
	}

	model.AddUser(model.User{Name: qname})

	user, inserted := model.FindUserByName(qname)
	if !inserted {
		log.Printf("failed to add user: %s", qname)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO is this the best way to construct location URI
	location := fmt.Sprintf("%s://%s/user?id=%d", r.URL.Scheme, r.URL.Host, user.ID)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}
