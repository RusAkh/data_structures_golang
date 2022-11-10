package main

import (
	"fmt"
)

func main() {
	var values = []int{1, 2, 3, 4, 5}
	root := Build_tree(values, 0)
	fmt.Println(root.height)
}
