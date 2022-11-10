package avl

// always pass 0 to i, will increase with recursion
func Build_tree(values []int, i int) *Node {
	var node *Node
	n := len(values)
	if i < n {
		node = &Node{value: values[i]}

		node.left = Build_tree(values, 2*i+1)

		node.right = Build_tree(values, 2*i+2)
	}
	setHeights(node)
	return node
}

func setHeights(root *Node) {
	root.height = Max(nodeHeight(root.left), nodeHeight(root.right)) + 1
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
	} else if root.left == nil && root.right == nil {
		return 0
	} else {
		return Max(nodeHeight(root.left), nodeHeight(root.right)) + 1
	}
}
