package main

import (
	"algoritmo_dijkstra/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shortest-path", handlers.ShortestPathHandler)
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}