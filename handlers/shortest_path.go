package handlers

import (
	"algoritmo_dijkstra/algoritmo"
	"algoritmo_dijkstra/models"
	"encoding/json"
	"net/http"
)

type Request struct {
	Graph models.Graph `json:"graph"`
	Start string       `json:"start"`
}

type Response struct {
	Distances map[string]int `json:"distances"`
}

func ShortestPathHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := algoritmo.Dijkstra(req.Graph, req.Start)
	resp := Response{Distances: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
