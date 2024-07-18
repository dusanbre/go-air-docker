package router

import (
	"dusanbre/test-1/handlers"
	"net/http"
)

func LoadRoutes(r *http.ServeMux) {
	commonHandler := &handlers.Handler{}

	r.HandleFunc("GET /", commonHandler.Hello)
}
