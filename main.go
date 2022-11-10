package main

import (
	avl "data_structures_go/avl_tree"
	"fmt"
)

func main() {
	var values = []int{1, 2, 3, 4, 5}
	root := avl.Build_tree(values, 0, len(values))
	fmt.Println(root.Height)
}
