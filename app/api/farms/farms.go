package farms

import (
	service "delos/app/service/farms"
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
)

// GET /farms - Get all farms
func getFarms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetFarms())
}

func Init(router *mux.Router) *mux.Router {
	router = initFarm(router)
	router.HandleFunc("/farms", getFarms).Methods("GET")
	return router
}
