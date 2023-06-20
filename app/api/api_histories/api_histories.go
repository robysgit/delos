package histories

import (
	service "delos/app/service/api_histories"
	"encoding/json"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

// GET /stats - Get all api stats
func getStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := service.GetApiHistorySummaries()
	var response map[string]map[string]string
	response = make(map[string]map[string]string)
	for _, r := range result {
		response[r.Method+" /"+r.Url] = make(map[string]string)
		response[r.Method+" /"+r.Url]["count"] = strconv.Itoa(r.Total)
		response[r.Method+" /"+r.Url]["unique_user_agent"] = strconv.Itoa(r.UniqueUser)
	}
	json.NewEncoder(w).Encode(response)
}

func Init(router *mux.Router) *mux.Router {
	router.HandleFunc("/stats", getStats).Methods("GET")
	return router
}
