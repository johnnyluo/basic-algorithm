package list

// LinkList linklist
type LinkList struct {
	head *LinkNode
	tail *LinkNode
	len  int
}

// LinkNode node
type LinkNode struct {
	Next  *LinkNode
	Prev  *LinkNode
	Value interface{}
}

// NewLinkList create a new double link list
func NewLinkList() (*LinkList, error) {
	return &LinkList{}, nil
}

// AddNode Add a node into the list
// O(1) , insert is constant
func (ll *LinkList) AddNode(value interface{}) error {
	if ll.head == nil && ll.tail == nil { // when the whole list is empty
		node := &LinkNode{
			Next:  nil,
			Prev:  nil,
			Value: value,
		}
		ll.head = node
		ll.tail = node
		return nil
	}
	newNode := &LinkNode{
		Next:  nil,
		Prev:  ll.tail,
		Value: value,
	}

	ll.tail.Next = newNode
	ll.tail = newNode
	return nil
}

// Remove a node
func (ll *LinkList) Remove(value interface{}) error {
	return nil
}

// Traverse LinkList, O(N)
func (ll *LinkList) Traverse(c func(node *LinkNode)) error {
	node := ll.head
	for {
		if nil == node {
			break
		}
		c(node)
		node = node.Next
	}
	return nil
}

// Reverse link list
func (ll *LinkList) Reverse() error {
	node := ll.tail
	ll.head = node
	for {
		node.Prev, node.Next = node.Next, node.Prev
		if node.Next == nil {
			ll.tail = node
			break
		}
		node = node.Next
	}
	// swap the head and tail
	return nil
}
