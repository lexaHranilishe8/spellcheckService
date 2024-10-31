package main

import (
	"log"
	"net/http"
	"spellcheck-service/api"
)

func main() {
	http.HandleFunc("/api/start-check", api.StartCheck)
	http.HandleFunc("/api/stats", api.GetStats)
	http.HandleFunc("/api/errors", api.GetErrors)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
