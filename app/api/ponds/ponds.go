package ponds

import (
	res "delos/app/api/ponds/responses"
	service "delos/app/service/ponds"
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
)

// GET /ponds - Get all ponds
func getPonds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := service.GetPonds()
	var response []res.PondResponse
	for _, r := range result {
		response = append(response, res.PondResponse{ID: r.ID, PondName: r.PondName, FarmId: r.FarmId})
	}
	json.NewEncoder(w).Encode(response)
}

func Init(router *mux.Router) *mux.Router {
	router = initPond(router)
	router.HandleFunc("/ponds", getPonds).Methods("GET")
	return router
}
