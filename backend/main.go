package main

import (
	"log"
	"net/http"

	apphttp "github.com/sebas3261/calculator/backend/internal/http"
)

func main(){
	port := ":8080"

	router := apphttp.NewRouer()

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}