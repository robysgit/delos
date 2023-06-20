package middleware

import (
	service "delos/app/service/trackers"

	"net/http"

	mux "github.com/gorilla/mux"
)

func trackingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_agent := r.Header.Get("User-Agent")
		service.TrackApi(&r.Method, &r.RequestURI, &user_agent)
		next.ServeHTTP(w, r)
	})
}

func Init(router *mux.Router) *mux.Router {
	router.Use(trackingMiddleware)
	return router
}
