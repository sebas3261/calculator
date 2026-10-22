package main

import (
	"log"
	"net/http"

	"github.com/sebas3261/calculator/backend/internal/calculator"
	"github.com/sebas3261/calculator/backend/internal/calculator/evaluator"
	"github.com/sebas3261/calculator/backend/internal/calculator/parser"
	apphttp "github.com/sebas3261/calculator/backend/internal/http"
)

func main() {
	port := ":8080"

	service := calculator.NewService(parser.NewExpressionParser(), evaluator.NewEvaluator())
	router := apphttp.NewRouter(service)

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}