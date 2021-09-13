package main

import "fmt"

// Create a graph structure
type Graph struct {
	Vertices []*Vertex
	Directed bool
}

// Create a vertext Structure
type Vertex struct {
	Key      int
	Vertices []*Vertex
}

// Add a new Graph
func NewGraph(directed bool) *Graph {
	return &Graph{
		Vertices: []*Vertex{},
		Directed: directed,
	}
}

// Add a new Vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: []*Vertex{},
	}
}

// Create add Vertex
func (g *Graph) AddVertex(key int) {
	if g.ContainsVertex(g.Vertices, key) {
		fmt.Printf("Vertex already exists for %v ::", key)
		return
	}
	v := NewVertex(key)
	g.Vertices = append(g.Vertices, v)
}

// Create add Edge
func (g *Graph) AddEdge(from, to int) {
	// Get vertext
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	// If one of them is nil donot procede
	if fromVertex == nil || toVertex == nil {
		fmt.Printf("Need from and to vertices %v", from)
		return
	}

	// Validate if fromVertexAdjacent already has toVertex
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

// Queue Implementation start
type Node struct {
	v    *Vertex
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) enqueue(v *Vertex) {
	n := &Node{
		v: v,
	}

	//Queue is empty
	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	// Current next must be node formed by our vertex
	q.tail.next = n
	// Current tail must be node formed by our vertex
	q.tail = n
}

func (q *Queue) dequeue() *Vertex {
	h := q.head

	if q.tail == nil {
		return nil
	}

	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return h.v
}

// Queue Implementation end

// BFS Disconnected version of the problem

func (g *Graph) BFSII(startVertexKey int, visitedVertices []bool) {
	// create a queue
	vertexQueue := &Queue{}

	// Start with the vertex mentioned
	currentVertex := g.GetVertex(startVertexKey)

	// Visit the current Node
	visitCb(currentVertex.Key)

	for currentVertex != nil {
		// Maintain the keys
		visitedVertices[currentVertex.Key] = true
		for _, v := range currentVertex.Vertices { // These represent adjacents for
			if !visitedVertices[v.Key] {
				visitCb(v.Key)
				vertexQueue.enqueue(v)

				// Once you push it to Queue try and
				// Maintain the keys
				visitedVertices[v.Key] = true
			}
		}

		// Change the current vertex
		currentVertex = vertexQueue.dequeue()
	}
}

func (g *Graph) BFSDisconnected(verticesCount int) {
	// Create a boolean array of vertices length
	visitedVertices := make([]bool, verticesCount)

	for i := 0; i < verticesCount; i++ {
		if !visitedVertices[i] {
			g.BFSII(i, visitedVertices)
		}
	}
}

// BFS Disconnected version ends

// BFS with given StartVertex
func (g *Graph) BFS(startVertexKey int, visitCb func(int)) {
	// Create a queue
	vertexQueue := &Queue{}

	// Use Hash or Array for visited Vertices
	// Length of Array should not exceed graphs number of vertices
	visitedVertices := make([]bool, len(g.Vertices))

	// Start with the vertex mentioned
	currentVertex := g.GetVertex(startVertexKey)

	// Visit the current Node
	visitCb(currentVertex.Key)

	for currentVertex != nil {
		// Maintain the keys
		visitedVertices[currentVertex.Key] = true
		for _, v := range currentVertex.Vertices { // These represent adjacents for
			if !visitedVertices[v.Key] {
				visitCb(v.Key)
				vertexQueue.enqueue(v)

				// Once you push it to Queue try and
				// Maintain the keys
				visitedVertices[v.Key] = true
			}
		}

		// Change the current vertex
		currentVertex = vertexQueue.dequeue()
	}
}

// DFS
func visitCb(k int) {
	fmt.Println(k)
}
func DFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	visited := map[int]bool{}

	currentVertex := g.GetVertex(startVertex.Key)
	// fmt.Println(currentVertex)
	if currentVertex == nil {
		return
	}

	visited[startVertex.Key] = true
	visitCb(startVertex.Key)

	for _, v := range currentVertex.Vertices {
		if visited[v.Key] {
			continue
		}

		DFS(g, v, visitCb)
	}
}

func main() {
	g := NewGraph(true)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(1, 3)
	// g.AddEdge(3, 4)
	// g.AddEdge(2, 4)

	// g.Print()

	// v := &Vertex{Key: 1}
	g.BFS(4, visitCb)
	// g.BFSDisconnected(len(g.Vertices))
}
