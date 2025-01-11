package main

import (
	"fmt"
)

// Graph представляет неориентированный граф
type Graph struct {
	Vertices map[int][]int
}

// NewGraph создает новый граф
func NewGraph() *Graph {
	return &Graph{Vertices: make(map[int][]int)}
}

// AddEdge добавляет ребро в граф
func (g *Graph) AddEdge(v1, v2 int) {
	g.Vertices[v1] = append(g.Vertices[v1], v2)
	g.Vertices[v2] = append(g.Vertices[v2], v1)
}

// BFS выполняет поиск в ширину
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Print(vertex, " ")

		for _, neighbor := range g.Vertices[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

// Пример использования неориентированного графа
func main() {
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 5)

	fmt.Print("BFS starting from vertex 1: ")
	graph.BFS(1) // 1 2 3 4 5
}
