package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/mkappus/checkin/src/events"
	"github.com/mkappus/checkin/src/users"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "data/madison.db")
	api := rest.NewApi()

	// User objects
	ss := &users.StudentStore{db}
	ssh := users.StudentStoreHandler{ss}

	// Event objects
	es := &events.EventStore{db}
	esh := events.EventStoreHandler{es}

	// Web API
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: true,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true

		},
		// AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// AllowedHeaders: []string{
		// 	"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		// AccessControlAllowCredentials: true,
		// AccessControlMaxAge:           3600,
	})

	router, err := rest.MakeRouter(
		// Student Handlers
		rest.Get("/students", ssh.GetAll),

		// Event handlers
		rest.Get("/checkins", esh.GetAllCheckins),
		rest.Put("/checkout/:id", esh.PutCheckout),
		rest.Post("/checkins/:loc", esh.PostAddCheckin),
	)
	if err != nil {
		panic(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))

}
