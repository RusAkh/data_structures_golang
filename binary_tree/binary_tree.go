package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Build_tree(arr []int, i, n int) *Node {
	var root *Node
	if i < n {
		root = &Node{data: arr[i]}

		root.left = Build_tree(arr, 2*i+1, n)

		root.right = Build_tree(arr, 2*i+2, n)

	}
	return root
}

func main() {
	var array []int
	var num_Nodes = 10
	for i := 0; i < num_Nodes; i++ {
		array = append(array, rand.Intn(100))
	}
	fmt.Println(Build_tree(array, 0, num_Nodes))
}
