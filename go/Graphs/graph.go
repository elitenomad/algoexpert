/*
	Underlying structure is Vertex and graph where the vertices are stored
	In terms of Hash
*/
package main

import "fmt"

type Vertex struct {
	Key int // Key is the unique identifier of the vertex
	// Vertices will describe vertices connected to this one
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
}

// We then create a constructor function for the Vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

type Graph struct {
	// Vertices describes all vertices contained in the graph
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
	// This will decide if it's a directed or undirected graph
	directed bool
}

// Define constructor functions that create
// a new directed and a undirected graph
func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

// AddVertex creates a new vertex with the given
// key and adds it to the graph
func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	// Add the vertices to the graph's vertex map
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func visitCb(n int) {
	fmt.Println("Vertex :::> ", n)
}

func BFS(g *Graph, startVertexKey int, visitCb func(int)) {
	// fmt.Println(startVertex.Vertices)

	// Create a queue
	vertexQueue := &Queue{}
	// Maintain the hash of visited vertices
	visitedVertices := map[int]bool{}

	// Start with the vertext mentioned
	currentVertex := g.Vertices[startVertexKey]

	for {
		// Visit the current Node
		visitCb(currentVertex.Key)
		// Maintain the keys
		visitedVertices[currentVertex.Key] = true

		// fmt.Println("VERTICES ::", currentVertex.Vertices)
		for _, v := range currentVertex.Vertices {
			if !visitedVertices[v.Key] {
				vertexQueue.enqueue(v)
			}
		}

		// Change the current vertext
		currentVertex = vertexQueue.dequeue()
		if currentVertex == nil {
			break
		}
	}
}

func DFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	visited := map[int]bool{}

	currentVertex := startVertex
	if currentVertex == nil {
		return
	}

	visited[currentVertex.Key] = true
	visitCb(currentVertex.Key)

	for _, v := range currentVertex.Vertices {
		if visited[v.Key] {
			continue
		}

		DFS(g, v, visitCb)
	}
}

// QUEUE Implementation
type Node struct {
	v    Vertex
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) enqueue(v *Vertex) {
	n := &Node{v: *v}

	// If queue is empty
	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *Queue) dequeue() *Vertex {
	n := q.head

	if n == nil {
		return nil
	}

	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return &n.v
}

// END OF QUEUE Implementation

func main() {
	// graph := NewDirectedGraph()
	g := NewDirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 1)

	// v := &Vertex{Key: 3, Vertices: g.Vertices[3].Vertices}
	BFS(g, 2, visitCb)
	fmt.Println("***")
	v := &Vertex{Key: 2}
	DFS(g, v, visitCb)
}
