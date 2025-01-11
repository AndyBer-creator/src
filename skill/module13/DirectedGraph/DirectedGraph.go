package main

import (
	"fmt"
	"math"
)

// DirectedGraph представляет ориентированный граф
type DirectedGraph struct {
	Vertices map[int]map[int]float64
}

// NewDirectedGraph создает новый ориентированный граф
func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{Vertices: make(map[int]map[int]float64)}
}

// AddEdge добавляет ребро в ориентированный граф
func (g *DirectedGraph) AddEdge(v1, v2 int, weight float64) {
	if g.Vertices[v1] == nil {
		g.Vertices[v1] = make(map[int]float64)
	}
	g.Vertices[v1][v2] = weight
}

// Dijkstra находит кратчайший путь от начальной вершины до всех остальных вершин
func (g *DirectedGraph) Dijkstra(start int) map[int]float64 {
	distances := make(map[int]float64)
	for vertex := range g.Vertices {
		distances[vertex] = math.Inf(1)
	}
	distances[start] = 0

	visited := make(map[int]bool)

	for len(visited) < len(g.Vertices) {
		var current int
		minDistance := math.Inf(1)

		for vertex, distance := range distances {
			if !visited[vertex] && distance < minDistance {
				minDistance = distance
				current = vertex
			}
		}

		visited[current] = true

		for neighbor, weight := range g.Vertices[current] {
			newDistance := distances[current] + weight
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
			}
		}
	}

	return distances
}

// Пример использования ориентированного графа
func main() {
	directedGraph := NewDirectedGraph()
	directedGraph.AddEdge(1, 2, 1.0)
	directedGraph.AddEdge(1, 3, 4.0)
	directedGraph.AddEdge(2, 3, 2.0)
	directedGraph.AddEdge(3, 4, 1.0)

	fmt.Println("Shortest paths from vertex 1:")
	distances := directedGraph.Dijkstra(1)
	for vertex, distance := range distances {
		fmt.Printf("Distance to vertex %d: %.2f\n", vertex, distance)
	}
}
