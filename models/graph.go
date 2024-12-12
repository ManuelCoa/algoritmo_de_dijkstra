package models

type Node struct {
	ID    string
	Edges map[string]int // Nodo destino -> Peso de la arista
}

type Graph struct {
	Nodes map[string]Node
}
