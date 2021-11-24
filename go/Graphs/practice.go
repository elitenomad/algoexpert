package main

import "fmt"

type Vertex struct {
	Vertices []*Vertex
	Key      int
}

type Graph struct {
	Vertices []*Vertex
	Directed bool
}

func NewVertex(key int) *Vertex {
	return &Vertex{Key: key, Vertices: []*Vertex{}}
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		Directed: directed,
		Vertices: []*Vertex{},
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

func (g *Graph) AddVertex(key int) {
	if g.ContainsVertex(key) {
		return
	}

	vertex := NewVertex(key)
	g.Vertices = append(g.Vertices, vertex)
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
	if !g.Directed {
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

func (g *Graph) DFS(startVertex *Vertex) {
	h := map[int]bool{}

	g.DFSRec(startVertex, &h)
}

func (g *Graph) DFSRec(startVertex *Vertex, h *map[int]bool) {
	v := g.GetVertex(startVertex.Key)

	if v == nil {
		return
	}

	(*h)[v.Key] = true
	fmt.Println(v.Key)
	for _, v := range v.Vertices {
		if _, found := (*h)[v.Key]; found {
			continue
		}

		g.DFSRec(v, h)
	}
}

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

// Algorithm BFS(G, v)
//     Q ← new empty FIFO queue
//     Mark v as visited.
//     Q.enqueue(v)
//     while Q is not empty
//         a ← Q.dequeue()
//         // Perform some operation on a.
//         for all unvisited neighbors x of a
//             Mark x as visited.
//             Q.enqueue(x)
func (g *Graph) BFS(s *Vertex) {
	q := &Queue{}
	h := map[int]bool{}

	h[s.Key] = true
	fmt.Println(s.Key)
	q.Enqueue(s)
	for !q.IsEmpty() {
		d := q.Dequeue()
		for _, v := range d.Vertices {
			if _, found := h[v.Key]; found {
				continue
			}

			fmt.Println(v.Key)
			h[v.Key] = true
			q.Enqueue(v)
		}
	}
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
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 4)

	// g.Print()

	v := g.GetVertex(0)
	g.BFS(v)
}
