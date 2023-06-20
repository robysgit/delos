package app

import (
	"log"
	"net/http"

	api_histories "delos/app/api/api_histories"
	farms "delos/app/api/farms"
	middleware "delos/app/api/middleware"
	ponds "delos/app/api/ponds"
	db "delos/app/db"

	"github.com/gorilla/mux"
)

func Start() {
	db.Init()
	// Initialize the router
	router := mux.NewRouter()

	// Register the route handlers
	router = api_histories.Init(router)
	router = farms.Init(router)
	router = ponds.Init(router)
	router = middleware.Init(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
