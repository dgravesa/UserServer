package controller

import (
	"github.com/dgravesa/WaterLogger-UserServer/model"
	"net/http"
	"strconv"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	qid, err := strconv.ParseUint(q.Get("id"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, found := model.FindUserByID(qid); !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	model.DeleteUser(qid)

	w.WriteHeader(http.StatusOK)
	return
}
