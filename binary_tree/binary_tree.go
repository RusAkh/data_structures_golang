package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	data  int
	left  *node
	right *node
}

func Build_tree(arr []int, i, n int) *node {
	var root *node
	if i < n {
		root = &node{data: arr[i]}

		root.left = Build_tree(arr, 2*i+1, n)

		root.right = Build_tree(arr, 2*i+2, n)

	}
	return root

}

func main() {
	var array []int
	var num_nodes = 10
	for i := 0; i < num_nodes; i++ {
		array = append(array, rand.Intn(100))
	}
	fmt.Println(Build_tree(array, 0, num_nodes))
}
