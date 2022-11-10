package avl

type Node struct {
	right   *Node
	left    *Node
	value   int
	height  int64
	balance uint32
}
