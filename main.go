package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk depth-first tree traversal
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value // process node

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// FullWalk the full tree, close output channel when done
func FullWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same checks whether two binary trees store the same sequence
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go FullWalk(t1, ch1)
	ch2 := make(chan int)
	go FullWalk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 { // no more nodes in trees
			break
		}
		if ok1 != ok2 { // trees with different number of nodes
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Trees equivalent?", Same(tree.New(1), tree.New(2)))
}
