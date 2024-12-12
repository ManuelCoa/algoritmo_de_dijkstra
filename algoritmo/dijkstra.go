package algoritmo

import "algoritmo_dijkstra/models"

// Dijkstra calcula los caminos más cortos desde un nodo origen.
func Dijkstra(graph models.Graph, start string) map[string]int {
	distances := make(map[string]int)
	visited := make(map[string]bool)

	// Inicializamos distancias
	for node := range graph.Nodes {
		distances[node] = int(^uint(0) >> 1) // Infinito
	}
	distances[start] = 0

	for {
		// Encuentra el nodo no visitado con la menor distancia
		var currentNode string
		shortestDistance := int(^uint(0) >> 1) // Infinito
		for node, distance := range distances {
			if !visited[node] && distance < shortestDistance {
				currentNode = node
				shortestDistance = distance
			}
		}

		if currentNode == "" {
			break // No quedan nodos por visitar
		}

		// Marcar nodo como visitado
		visited[currentNode] = true

		// Actualizar distancias a los vecinos
		for neighbor, weight := range graph.Nodes[currentNode].Edges {
			if newDistance := distances[currentNode] + weight; newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
			}
		}
	}

	return distances
}
