package list

// LinkList linklist
type LinkList struct {
	head *LinkNode
	tail *LinkNode
	len  int
}

// LinkNode node
type LinkNode struct {
	next  *LinkNode
	prev  *LinkNode
	value interface{}
}
