package farms

import (
	req "delos/app/api/farms/requests"
	res "delos/app/api/farms/responses"
	service "delos/app/service/farms"
	model "delos/app/service/model"
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
)

// GET /farm - Get all farms
func getFarm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetFarm(&id))
}

// POST /farm - Create a new farm
func createFarm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var farm req.CreateFarmRequest
	err := json.NewDecoder(r.Body).Decode(&farm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := *service.CreateFarm(&model.Farm{FarmName: farm.FarmName})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res.FarmResponse{ID: result.ID, FarmName: result.FarmName})
}

// PUT /farm - Update a farm
func updateFarm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var farm req.UpdateFarmRequest
	err := json.NewDecoder(r.Body).Decode(&farm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id != farm.ID {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := *service.UpdateFarm(&model.Farm{ID: farm.ID, FarmName: farm.FarmName})

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&res.FarmResponse{ID: result.ID, FarmName: result.FarmName})
}

// DELETE /farm - Delete a farm
func deleteFarm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	result := service.DeleteFarm(&id)
	if result == 0 {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}

func initFarm(router *mux.Router) *mux.Router {
	router.HandleFunc("/farm/{id}", getFarm).Methods("GET")
	router.HandleFunc("/farm", createFarm).Methods("POST")
	router.HandleFunc("/farm/{id}", updateFarm).Methods("PUT")
	router.HandleFunc("/farm/{id}", deleteFarm).Methods("DELETE")
	return router
}
