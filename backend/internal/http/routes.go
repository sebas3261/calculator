package http

import (
	"net/http"
)

func NewRouer() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthHandler)
	
	return mux
}