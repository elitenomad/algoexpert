package main

import "fmt"

// Queue Implemenatation start
type VNode struct {
	v    *Vertex
	Next *VNode
}

type Queue struct {
	head *VNode
	tail *VNode
}

func (q *Queue) enqueue(v *Vertex) {
	current := &VNode{v: v}

	if q.tail == nil {
		q.head = current
		q.tail = current
		return
	}

	// Ensure current Next points to current
	q.tail.Next = current
	// Then set current tail to current
	q.tail = current
}

func (q *Queue) dequeue() *Vertex {
	h := q.head

	if h == nil {
		return nil
	}

	q.head = h.Next

	if q.head == nil {
		q.tail = nil
	}

	return h.v
}

func (q *Queue) isEmpty() bool {
	return q.tail == nil && q.head == nil
}

// Queue Implemenation end

// Graph Implementation start
type Graph struct {
	Vertices []*Vertex
	Directed bool
}

type Vertex struct {
	Key      int
	Vertices []*Vertex
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		Directed: directed,
		Vertices: []*Vertex{},
	}
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: []*Vertex{},
	}
}

func (g *Graph) AddVertex(key int) {
	v := &Vertex{
		Key: key,
	}

	g.Vertices = append(g.Vertices, v)
}

func (g *Graph) AddEdge(from, to int) {
	// from and to represent vertices.Goal is to create an edge b/w then
	// Fetch from and to
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	// Ensure you proceed only when vertices are present in the graph
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Both vertices must exist on the graph")
		return
	}

	// Ensure you proceed only if the edges are not present b/w them
	if g.ContainsVertex(fromVertex.Vertices, to) || g.ContainsVertex(toVertex.Vertices, from) {
		fmt.Println("Edge already exists b/w them")
		return
	}

	fromVertex.Vertices = append(fromVertex.Vertices, toVertex)
	if !g.Directed {
		toVertex.Vertices = append(toVertex.Vertices, fromVertex)
	}
}

func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}

	return nil
}

func (g *Graph) ContainsVertex(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if v.Key == key {
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

// BFS with given StartVertex
func (g *Graph) ShortestPaths(startVertexKey int) []int {
	// Create a queue
	vertexQueue := &Queue{}

	// Use Hash or Array for visited Vertices
	// Length of Array should not exceed graphs number of vertices
	visitedVertices := make([]bool, len(g.Vertices))
	distanceVertices := make([]int, len(g.Vertices))

	// Start with the vertex mentioned
	currentVertex := g.GetVertex(startVertexKey)
	// Maintain the keys for traversal
	visitedVertices[currentVertex.Key] = true
	// Maintain the keys for distance from startVertexKey
	distanceVertices[startVertexKey] = 0

	// Start the processes. Insert the startVertexKey vertex
	vertexQueue.enqueue(currentVertex)

	for !vertexQueue.isEmpty() {
		// Change the current vertex
		u := vertexQueue.dequeue()

		for _, v := range u.Vertices { // These represent adjacents for
			if !visitedVertices[v.Key] {
				distanceVertices[v.Key] = distanceVertices[u.Key] + 1

				// Once you push it to Queue try and
				// Maintain the keys
				visitedVertices[v.Key] = true
				vertexQueue.enqueue(v)
			}
		}
	}

	return distanceVertices
}

func (g *Graph) ShortestPathPractice(startVertexKey int) []int {
	// Handle a use case when the vertex is not present in the Graph
	// That way you are informing interviewer that you are thinking of all
	// the use cases.

	// Fetch the vertex
	startVertex := g.GetVertex(startVertexKey)

	// Initialize Queue
	queue := &Queue{}

	// Have hash or array to know the visited vertices
	visitedVertices := make([]bool, len(g.Vertices))
	distanceVertices := make([]int, len(g.Vertices))

	// Start with Initial value, push it to queue
	queue.enqueue(startVertex)

	// Visited the first vertex
	visitedVertices[startVertex.Key] = true
	distanceVertices[startVertex.Key] = 0

	for !queue.isEmpty() {
		u := queue.dequeue()

		for _, v := range u.Vertices {
			if !visitedVertices[v.Key] {
				distanceVertices[v.Key] = distanceVertices[u.Key] + 1

				visitedVertices[v.Key] = true
				queue.enqueue(v)
			}
		}
	}

	return distanceVertices
}

// Graph Implementation end
func main() {
	g := NewGraph(false)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	fmt.Println(g.ShortestPaths(0))
	fmt.Println(g.ShortestPathPractice(0))
}
