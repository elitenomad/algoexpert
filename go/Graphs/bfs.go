package main

import "fmt"

type Vertex struct {
	Name     string
	Vertices []*Vertex
}

func NewVertex(name string) *Vertex {
	return &Vertex{
		Name:     name,
		Vertices: []*Vertex{},
	}
}

func (n *Vertex) AddVertices(names ...string) *Vertex {
	for _, name := range names {
		child := Vertex{Name: name}
		n.Vertices = append(n.Vertices, &child)
	}

	return n
}

// O(V + E) and O(V) space
func (n *Vertex) BreadthFirstSearch(input []string) []string {
	// Create a Queue
	queue := []*Vertex{n}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		input = append(input, current.Name)
		for _, child := range current.Vertices {
			queue = append(queue, child)
		}
	}

	return input
}

func main() {
	var graph = NewVertex("A").AddVertices("B", "C", "D")
	graph.Vertices[0].AddVertices("E").AddVertices("F")
	graph.Vertices[2].AddVertices("G").AddVertices("H")
	graph.Vertices[0].Vertices[1].AddVertices("I").AddVertices("J")
	graph.Vertices[2].Vertices[0].AddVertices("K")
	fmt.Println(graph.BreadthFirstSearch([]string{}))
}
