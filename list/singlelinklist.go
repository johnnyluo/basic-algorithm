package list

// SingleLinkList it is a single linked list
type SingleLinkList struct {
	head *SingleLinkNode
}

// SingleLinkNode a node
type SingleLinkNode struct {
	Next  *SingleLinkNode
	Value interface{}
}

// NewSingleLinkList create a new single linked list
func NewSingleLinkList() *SingleLinkList {
	return &SingleLinkList{}
}

// AddNode add a node to the list
// need to traverse the linked list to the end
func (sll *SingleLinkList) AddNode(value interface{}) error {
	if sll.head == nil {
		node := &SingleLinkNode{
			Value: value,
			Next:  nil,
		}
		sll.head = node
		return nil
	}
	tmpnode := sll.head
	for {
		if tmpnode.Next == nil {
			break
		}
		tmpnode = tmpnode.Next
	}
	newNode := &SingleLinkNode{
		Value: value,
		Next:  nil,
	}
	tmpnode.Next = newNode
	return nil
}

// Traverse the nodes
func (sll *SingleLinkList) Traverse(c func(node *SingleLinkNode)) error {
	tmpnode := sll.head
	for {
		if tmpnode == nil {
			return nil
		}
		c(tmpnode)
		tmpnode = tmpnode.Next
	}
}

// Reverse the list
func (sll *SingleLinkList) Reverse() error {
	current := sll.head
	// empty or one node
	if current == nil || current.Next == nil {
		return nil
	}
	var first *SingleLinkNode
	var next *SingleLinkNode
	for {
		next = current.Next
		current.Next = first
		first = current
		current = next
		if current == nil {
			sll.head = first
			break
		}
	}
	return nil
}
