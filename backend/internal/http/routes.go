package http

import (
	"net/http"

	"github.com/sebas3261/calculator/backend/internal/calculator"
)

func NewRouter(service *calculator.Service) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthHandler)
	mux.HandleFunc("POST /calculate", NewCalculateHandler(service))

	return mux
}