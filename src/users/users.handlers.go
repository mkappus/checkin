package users

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type StudentStoreHandler struct {
	*StudentStore
}

func (ssh StudentStoreHandler) GetAll(w rest.ResponseWriter, r *rest.Request) {
	students, err := ssh.All()
	if err != nil {
		rest.Error(w, err.Error(), 404)
		return
	}
	w.WriteJson(students)
}
