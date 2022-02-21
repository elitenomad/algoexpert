package main

import "fmt"

type Vertex struct {
	Vertices []*Vertex
	Value    int
}

type Graph struct {
	Vertices []*Vertex
	Directed bool
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		Vertices: []*Vertex{},
		Directed: directed,
	}
}

func NewVertex(v int) *Vertex {
	return &Vertex{Value: v, Vertices: []*Vertex{}}
}

func (g *Graph) AddVertex(v int) {
	if g.ContainsVertex(v) {
		return
	}

	vertex := &Vertex{Value: v, Vertices: []*Vertex{}}
	g.Vertices = append(g.Vertices, vertex)
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		return
	}

	if g.ContainsEdgeBetween(fromVertex, toVertex) {
		return
	}

	fromVertex.Vertices = append(fromVertex.Vertices, toVertex)
	if !g.Directed {
		toVertex.Vertices = append(toVertex.Vertices, fromVertex)
	}
}

func (g *Graph) ContainsEdgeBetween(from, to *Vertex) bool {
	var edgeOne bool
	for _, vertex := range from.Vertices {
		if vertex.Value == to.Value {
			edgeOne = true
		}
	}

	var edgeTwo bool
	for _, vertex := range to.Vertices {
		if vertex.Value == from.Value {
			edgeTwo = true
		}
	}

	return edgeOne && edgeTwo
}

func (g *Graph) GetVertex(v int) *Vertex {
	var result *Vertex

	for _, vertex := range g.Vertices {
		if vertex.Value == v {
			result = vertex
		}
	}

	return result
}

func (g *Graph) ContainsVertex(v int) bool {
	for _, vertex := range g.Vertices {
		if v == vertex.Value {
			return true
		}
	}

	return false
}

/*
	For BFS disconnected:
		Call the below function with Visited Array
		Have count and incremenet for each vertext passing - connected Islands
*/
func (g *Graph) BFS(start int) []int {
	s := g.GetVertex(start)
	visted := make([]bool, len(g.Vertices)+1)
	queue := []*Vertex{s}
	visted[start] = true
	result := []int{}

	for len(queue) > 0 {
		current := queue[0]
		fmt.Println(current.Value)
		queue = queue[1:]
		result = append(result, current.Value)
		for _, child := range current.Vertices {
			if visted[child.Value] == false {
				visted[child.Value] = true
				queue = append(queue, child)
			}
		}
	}

	return result
}

func (g *Graph) DFS(start int) []int {
	result := []int{}
	startVertex := g.GetVertex(start)

	visted := make([]bool, len(g.Vertices)+1)
	visted[start] = true
	g.DFSHelper(startVertex, &result, visted)

	return result
}

func (g *Graph) DFSHelper(start *Vertex, result *[]int, visited []bool) {
	*result = append(*result, start.Value)

	for _, vertex := range g.Vertices {
		if visited[vertex.Value] == false {
			visited[vertex.Value] = true
			g.DFSHelper(vertex, result, visited)
		}

	}
}

// Goal is to use BFS to find shortest paths
// From a given vertex to all the other vertices
func (g *Graph) ShortestPaths(start int) []int {
	// Visited vertices
	visitedVertices := make([]bool, len(g.Vertices)+1)
	distanceVertices := make([]int, len(g.Vertices)+1)

	// Get vertex
	startVertex := g.GetVertex(start)
	//queue
	queue := []*Vertex{startVertex}

	// set bool for start
	// Initial distance from start
	visitedVertices[start] = true
	distanceVertices[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, vertex := range current.Vertices {
			if visitedVertices[vertex.Value] == false {
				visitedVertices[vertex.Value] = true
				distanceVertices[vertex.Value] = distanceVertices[current.Value] + 1
				queue = append(queue, vertex)
			}
		}
	}

	return distanceVertices
}

// Detect a cycle in the graph
func (g *Graph) DetectCycleHelper(visited map[int]bool, vertex *Vertex, parent *Vertex) bool {
	visited[vertex.Value] = true
	for _, v := range vertex.Vertices {
		if g.DetectCycleHelper(visited, v, vertex) == true {
			return true
		} else if v != parent {
			return true
		}
	}

	return false
}

func (g *Graph) DetectCycle() bool {
	visited := map[int]bool{}

	for _, vertex := range g.Vertices {
		if _, found := visited[vertex.Value]; !found {
			if g.DetectCycleHelper(visited, vertex, nil) == true {
				return true
			}
		}
	}

	return false
}

func main() {
	graph := NewGraph(false)
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 6)

	fmt.Println(graph.BFS(1))
	fmt.Println(graph.DFS(0))
}
