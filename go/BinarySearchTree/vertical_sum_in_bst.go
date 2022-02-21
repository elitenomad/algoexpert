package main

import (
	"container/list"
	"fmt"
	"math"

	"github.com/emirpasic/gods/maps/treemap"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func VerticalSum(tree *BST) []int {
	m := treemap.NewWithIntComparator()
	result := []int{}
	VerticalSumHelper(tree, 0, m)

	m.Each(func(key, value interface{}) {
		result = append(result, value.(int))
	})

	return result
}

func VerticalSumHelper(tree *BST, hd int, m *treemap.Map) {
	if tree == nil {
		return
	}
	VerticalSumHelper(tree.Left, hd-1, m)

	val, found := m.Get(hd)
	if found {
		m.Put(hd, val.(int)+tree.Value)
	} else {
		m.Put(hd, tree.Value)
	}
	// fmt.Println(m)
	VerticalSumHelper(tree.Right, hd+1, m)
}

func VerticalOrderTraversal(tree *BST) [][]int {
	m := map[int][]int{}
	min := math.MaxInt32
	max := math.MinInt32
	VerticalOrderTraversalHelper(tree, &min, &max, 0, m)

	result := [][]int{}
	for i := min; i <= max; i++ {
		result = append(result, m[i])
	}

	return result
}

func VerticalOrderTraversalHelper(tree *BST, min *int, max *int, hd int, m map[int][]int) {
	if tree == nil {
		return
	}
	VerticalOrderTraversalHelper(tree.Left, min, max, hd-1, m)
	*min = minimum(*min, hd)
	*max = maximum(*max, hd)
	if _, found := m[hd]; found {
		m[hd] = append(m[hd], tree.Value)
	} else {
		m[hd] = []int{tree.Value}
	}
	VerticalOrderTraversalHelper(tree.Right, min, max, hd+1, m)
}

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Pair struct {
	node     *BST
	distance int
}

func VerticalOrderWithLeveOrderTraversal(tree *BST) [][]int {
	queue := list.New()
	result := [][]int{}
	if tree == nil {
		return result
	}

	queue.PushBack(Pair{node: tree, distance: 0})

	m := map[int][]int{}
	minColumn := math.MaxInt32
	maxColumn := math.MinInt32

	for queue.Len() != 0 {
		d := queue.Front()
		queue.Remove(d)

		p := d.Value.(Pair)
		root := p.node
		dist := p.distance

		if root != nil {
			// fmt.Println(root.Value)
			m[dist] = append(m[dist], root.Value)
			minColumn = minimum(minColumn, dist)
			maxColumn = maximum(maxColumn, dist)

			queue.PushBack(Pair{node: root.Left, distance: dist - 1})
			queue.PushBack(Pair{node: root.Right, distance: dist + 1})
		}
	}

	fmt.Println(minColumn, maxColumn)
	for i := minColumn; i <= maxColumn; i++ {
		result = append(result, m[i])
	}
	return result
}

func main() {
	root := &BST{}
	root.Value = 15
	root.Left = &BST{Value: 5}
	root.Left.Left = &BST{Value: 3}
	root.Left.Right = &BST{Value: 30}
	root.Right = &BST{Value: 20}
	root.Right.Left = &BST{Value: 18}
	root.Right.Left.Left = &BST{Value: 16}
	root.Right.Right = &BST{Value: 80}

	fmt.Println(VerticalSum(root))
	fmt.Println(VerticalOrderTraversal(root))
	fmt.Println(VerticalOrderWithLeveOrderTraversal(root))
}
