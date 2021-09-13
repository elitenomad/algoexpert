package main

import "fmt"

type Vertex struct {
	Key      int
	Vertices []*Vertex
}

type Graph struct {
	Directed bool
	Vertices []*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: []*Vertex{},
	}
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		Directed: directed,
		Vertices: []*Vertex{},
	}
}

// Create add Vertex
func (g *Graph) AddVertex(key int) {
	// Verify if the Vertex has already contains the key in vertices
	if g.ContainsVertex(g.Vertices, key) {
		fmt.Printf("Vertex already exists for %v ::", key)
		return
	}

	// Else create new  vertex
	v := NewVertex(key)
	g.Vertices = append(g.Vertices, v)
}

// Create add Edge
func (g *Graph) AddEdge(from, to int) {
	// Get vertexes of from and to
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	// If one of them is nil do not procede
	if fromVertex == nil || toVertex == nil {
		fmt.Printf("Need from and to vertices %v", from)
		return
	}

	// Validate if fromVertex Adjacent already has toVertex
	if g.ContainsVertex(fromVertex.Vertices, toVertex.Key) {
		fmt.Printf("\nEdge already exists for %v to %v", fromVertex.Key, toVertex.Key)
		return
	}

	fromVertex.Vertices = append(fromVertex.Vertices, toVertex)
	if g.Directed {
		toVertex.Vertices = append(toVertex.Vertices, fromVertex)
	}
}

// Get vertex
func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}
	return nil
}

// Contains the vertex already
func (g *Graph) ContainsVertex(vertices []*Vertex, k int) bool {
	for _, v := range vertices {
		if v.Key == k {
			return true
		}
	}

	return false
}

// Print adjacent list for each vertex
func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\n Vertex %v ", v.Key)

		if len(v.Vertices) <= 0 {
			continue
		}
		fmt.Printf("\t Adjacent")
		for _, v := range v.Vertices {
			fmt.Printf("\t %v", v.Key)
		}
	}
}

func visitCb(key int) {
	fmt.Println(key)
}

func (g *Graph) DFSRecursive(startVertexKey int, visitedVertices []bool, parent int) bool {
	visitedVertices[startVertexKey] = true
	currentVertex := g.GetVertex(startVertexKey)

	for _, v := range currentVertex.Vertices {
		if !visitedVertices[v.Key] {
			if g.DFSRecursive(v.Key, visitedVertices, currentVertex.Key) == true {
				return true
			}
		} else if v.Key != parent {
			return true
		}
	}

	return false
}

func (g *Graph) CycleExistsUsingDFS() bool {
	visitedVertices := make([]bool, len(g.Vertices))

	for _, v := range g.Vertices {
		if !visitedVertices[v.Key] {
			if g.DFSRecursive(v.Key, visitedVertices, -1) == true {
				return true
			}
		}
	}

	return false
}

func main() {
	g := NewGraph(false)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)

	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	g.AddEdge(4, 5)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	// fmt.Println(g.GetVertex(1).Vertices[1].Vertices[0].Key)

	// startVertex := &Vertex{Key: 3}
	fmt.Println(g.CycleExistsUsingDFS())
}
