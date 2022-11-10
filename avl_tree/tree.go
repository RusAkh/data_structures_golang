package avl

import "fmt"

// Always pass 0 to i, will increase with recursion.
// It Always passes same n, because computing len() on each iteration is costly
func Build_tree(values []int, i, n int) *Node {
	var node *Node
	if i < n {
		node = &Node{Value: values[i]}
		fmt.Println(node)
		node.Left = Build_tree(values, 2*i+1, n)

		node.Right = Build_tree(values, 2*i+2, n)
	}
	setHeights(node)
	return node
}

func setHeights(root *Node) {
	fmt.Println("root ", root)
	root.Height = Max(nodeHeight(root.Left), nodeHeight(root.Right)) + 1
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func nodeHeight(root *Node) int64 {
	if root == nil {
		return -1
	} else if root.Left == nil && root.Right == nil {
		return 0
	} else {
		return Max(nodeHeight(root.Left), nodeHeight(root.Right)) + 1
	}
}
