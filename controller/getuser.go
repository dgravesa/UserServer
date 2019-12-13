package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgravesa/WaterLogger-UserServer/model"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	qname := q.Get("name")
	qidstr := q.Get("id")
	qid, qiderr := strconv.ParseUint(qidstr, 10, 64)

	var userByID, userByName *model.User

	// find user by id
	if qidstr != "" {
		if qiderr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, found := model.FindUserByID(qid)
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		userByID = &user
	}

	// find user by name
	if qname != "" {
		user, found := model.FindUserByName(qname)
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		userByName = &user
	}

	// no query provided
	if userByID == nil && userByName == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userToWrite *model.User

	if userByID != nil && userByName != nil {
		// compare users from both query parameters for differences
		if userByID.ID != userByName.ID || userByID.Name != userByName.Name {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		userToWrite = userByID
	} else if userByID != nil {
		userToWrite = userByID
	} else {
		userToWrite = userByName
	}

	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.Encode(userToWrite)
}
