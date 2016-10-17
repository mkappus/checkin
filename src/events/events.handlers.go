package events

import (
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/mkappus/checkin/src/users"
)

type EventStoreHandler struct {
	*EventStore
}

func (esh EventStoreHandler) GetAllCheckins(w rest.ResponseWriter, r *rest.Request) {
	events, err := esh.AllCheckins()
	if err != nil {
		rest.Error(w, err.Error(), 404)
		return
	}
	w.WriteJson(events)
}

func (esh EventStoreHandler) PostAddCheckin(w rest.ResponseWriter, r *rest.Request) {
	loc := r.PathParam("loc")
	// TODO: Enumerate locations and verify valid loc
	s := new(users.Student)
	if err := r.DecodeJsonPayload(&s); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c, err := esh.AddCheckin(s, loc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteJson(&c)

}

func (esh EventStoreHandler) PutCheckout(w rest.ResponseWriter, r *rest.Request) {
	cID, err := strconv.Atoi(r.PathParam("id"))
	if err = esh.Checkout(cID); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
