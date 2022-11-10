package avl

type Node struct {
	Right   *Node
	Left    *Node
	Value   int
	Height  int64
	Balance uint32
}
