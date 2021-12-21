package main

import (
	"fmt"
)

// *************************************
type Queue struct {
	Elements []int
}

func NewQueue() *Queue {
	return &Queue{}

}

func (q *Queue) Enqueue(element int) {
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Dequeue() int { // O(N)
	removedElement := q.Elements[0]
	q.Elements = q.Elements[1:]
	return removedElement
}

func (q *Queue) Front() int {
	return q.Elements[0]
}

func (q *Queue) Rear() int {
	return q.Elements[len(q.Elements)-1]
}

func (q *Queue) Size() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) <= 0
}

// ******************************************

type Graph struct {
	v        int
	e        int
	adj      [][]int
	directed bool
}

func New(v int, e int, directed bool) *Graph {
	adj := make([][]int, v)

	// By default you have zeroes everywhere
	for i := 0; i < v; i++ {
		adj[i] = make([]int, v)
	}

	return &Graph{
		v:   v,
		e:   e,
		adj: adj,
	}
}

func (g *Graph) addEdge(from, to int) {
	if from < 0 && from > g.v {
		return
	}

	if to < 0 && to > g.v {
		return
	}

	g.adj[from][to] = 1
	if !g.directed {
		g.adj[to][from] = 1
	}
}

func (g *Graph) BFS(start int) {
	visited := make([]bool, g.v)
	queue := NewQueue()

	queue.Enqueue(start)
	visited[start] = true

	for !queue.IsEmpty() {
		vis := queue.Front()
		fmt.Printf("%d \t", vis)
		queue.Dequeue()

		for i := 0; i < g.v; i++ {
			if g.adj[vis][i] == 1 && !visited[i] {
				queue.Enqueue(i)
				visited[i] = true
			}
		}
	}
}

func main() {
	v := 5
	e := 4

	g := New(v, e, false)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 3)
	g.BFS(0)
}
