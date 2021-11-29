package main

import "fmt"

/*
  Queue Implementation
*/
type Node struct {
	Val  *Vertex
	Next *Node
}
type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(v *Vertex) {
	n := &Node{Val: v}

	if q.Tail == nil {
		q.Head = n
		q.Tail = n

		return
	}

	q.Tail.Next = n
	q.Tail = n
}

func (q *Queue) Dequeue() *Vertex {
	h := q.Head

	if q.Tail == nil {
		return nil
	}

	q.Head = q.Head.Next

	if q.Head == nil {
		q.Tail = nil
	}

	return h.Val
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil && q.Tail == nil
}

type Vertex struct {
	Key      int
	Vertices []*Vertex
}

type Graph struct {
	Vertices   []*Vertex
	IsDirected bool
}

func NewVertex(key int) *Vertex {
	return &Vertex{Key: key, Vertices: []*Vertex{}}
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		IsDirected: directed,
		Vertices:   []*Vertex{},
	}
}

func (g *Graph) AddVertex(key int) {
	if g.ContainsVertex(key) {
		return
	}

	v := NewVertex(key)
	g.Vertices = append(g.Vertices, v)
}

func (g *Graph) ContainsVertex(key int) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}

	return false
}

func (g *Graph) ContainsVertexIn(v *Vertex, key int) bool {
	for _, t := range v.Vertices {
		if t.Key == key {
			return true
		}
	}

	return false
}

func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}

	return nil
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		return // Goal is we can't create an edge with unknown vertices
	}

	if g.ContainsVertexIn(fromVertex, toVertex.Key) {
		return // Edge already exists b/w from and to
	}

	fromVertex.Vertices = append(fromVertex.Vertices, toVertex)
	if !g.IsDirected {
		toVertex.Vertices = append(toVertex.Vertices, fromVertex)
	}
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

func (g *Graph) BFS(s *Vertex, h *map[int]bool) {
	q := &Queue{}

	(*h)[s.Key] = true
	// fmt.Println(s.Key)
	q.Enqueue(s)
	for !q.IsEmpty() {
		d := q.Dequeue()
		for _, v := range d.Vertices {
			if _, found := (*h)[v.Key]; found {
				continue
			}

			// fmt.Println(v.Key)
			(*h)[v.Key] = true
			q.Enqueue(v)
		}
	}
}

func (g *Graph) PrintIslandsUsingBFSDisconnected() int {
	h := map[int]bool{}

	count := 0
	for _, v := range g.Vertices {
		if _, found := h[v.Key]; found {
			continue
		}

		g.BFS(v, &h)
		count += 1
	}

	return count
}

func main() {
	g := NewGraph(false)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 7)
	g.AddEdge(7, 8)

	fmt.Println(g.PrintIslandsUsingBFSDisconnected())
}
