package ponds

import (
	req "delos/app/api/ponds/requests"
	res "delos/app/api/ponds/responses"
	model "delos/app/service/model"
	service "delos/app/service/ponds"
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
)

// GET /pond - Get pond
func getPond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	json.NewEncoder(w).Encode(service.GetPond(&id))
}

// POST /pond - Create a new pond
func createPond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var pond req.CreatePondRequest
	err := json.NewDecoder(r.Body).Decode(&pond)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := *service.CreatePond(&model.Pond{PondName: pond.PondName, FarmId: pond.FarmId})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res.PondResponse{ID: result.ID, PondName: result.PondName, FarmId: result.FarmId})
}

// PUT /pond - Update a pond
func updatePond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var pond req.UpdatePondRequest
	err := json.NewDecoder(r.Body).Decode(&pond)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id != pond.ID {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := *service.UpdatePond(&model.Pond{ID: pond.ID, PondName: pond.PondName, FarmId: pond.FarmId})

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&res.PondResponse{ID: result.ID, PondName: result.PondName, FarmId: result.FarmId})
}

// DELETE /pond - Delete a pond
func deletePond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	result := service.DeletePond(&id)
	if result == 0 {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}

func initPond(router *mux.Router) *mux.Router {
	router.HandleFunc("/pond/{id}", getPond).Methods("GET")
	router.HandleFunc("/pond", createPond).Methods("POST")
	router.HandleFunc("/pond/{id}", updatePond).Methods("PUT")
	router.HandleFunc("/pond/{id}", deletePond).Methods("DELETE")
	return router
}
