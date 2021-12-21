package main

import "fmt"

// Kahn's Algorithm - Print Topological sort values

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

func (q *Queue) IsEmpty() bool {
	return q.head == nil && q.tail == nil
}

// END OF QUEUE Implementation

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
	if !g.Directed {
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
	for _, u := range g.Vertices {
		fmt.Printf("\n Vertex %v ", u.Key)

		if len(u.Vertices) <= 0 {
			continue
		}
		fmt.Printf("\t Adjacent")
		for _, v := range u.Vertices {
			fmt.Printf("\t %v", v.Key)
		}
	}
}

func (g *Graph) IngressOfVertices() []int {
	ingress := make([]int, len(g.Vertices))

	for _, v := range g.Vertices {
		for _, u := range v.Vertices {
			ingress[u.Key] += 1
		}
	}

	return ingress
}

func visitCb(key int) {
	fmt.Println(key)
}

// For a graph there are multiple versions of topological sort

// Algorithm
// Store Indegree of every vertex
// Create a Queue
// All all 0 degree to the Queue (They are the ones you need to display first)
// while queue is not empty:
// 		u = queue.pop()
// 		print u
//		for every adjacent v of u:
//			reduce indegree of v by 1
//			if indegree becomes 0:
//				add v to the Queue

func (g *Graph) TopologicalSort() {
	indegrees := g.IngressOfVertices()
	queue := &Queue{}

	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			queue.enqueue(g.Vertices[i])
		}
	}

	for !queue.IsEmpty() {
		u := queue.dequeue()
		fmt.Println(u.Key)
		for _, v := range u.Vertices {
			indegrees[v.Key] -= 1
			if indegrees[v.Key] == 0 {
				queue.enqueue(v)
			}
		}
	}
}

func (g *Graph) DetectCycleUsingKahnAlgorithm() bool {
	indegrees := g.IngressOfVertices()
	queue := &Queue{}

	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			queue.enqueue(g.Vertices[i])
		}
	}

	count := 0
	for !queue.IsEmpty() {
		u := queue.dequeue()
		count += 1

		for _, v := range u.Vertices {
			indegrees[v.Key] -= 1
			if indegrees[v.Key] == 0 {
				queue.enqueue(v)
			}
		}
	}

	fmt.Println(count, len(g.Vertices))
	return count != len(g.Vertices)
}

/*
	Shortest path in Direct acyclic Graph
	Initialize distances[v] = [INF,INF...INF]
	distances[s] = 0
	Find topological sort of a graph
	for every vertex in topological sort array:
	 loop through adjacent v of u:
	     if distance[v] > distances[u] + weight(u, v):
		 	distances[v] = distances[u] + weight(u, v)
*/

func main() {
	g := NewGraph(true)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)

	g.AddEdge(0, 1)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	// g.AddEdge(5, 3)

	fmt.Println(g.IngressOfVertices())
	g.TopologicalSort()
	fmt.Println(g.DetectCycleUsingKahnAlgorithm())
}
