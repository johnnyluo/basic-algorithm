package tree

// Node tree node
type Node struct {
	left  *Node
	right *Node
	value interface{}
}
