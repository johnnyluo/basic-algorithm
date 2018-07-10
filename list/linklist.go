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

// NewLinkList create a new double link list
func NewLinkList() (*LinkList, error) {
	return &LinkList{}, nil
}

// AddNode Add a node into the list
// O(1) , insert is constant
func (ll *LinkList) AddNode(value interface{}) error {
	if ll.head == nil && ll.tail == nil { // when the whole list is empty
		node := &LinkNode{
			next:  nil,
			prev:  nil,
			value: value,
		}
		ll.head = node
		ll.tail = node
		return nil
	}
	newNode := &LinkNode{
		next:  nil,
		prev:  ll.tail,
		value: value,
	}

	ll.tail.next = newNode
	ll.tail = newNode
	return nil
}

// Remove a node
func (ll *LinkList) Remove(value interface{}) error {
	return nil
}

// traverse LinkList, O(N)
func (ll *LinkList) traverse(c func(node *LinkNode)) error {
	node := ll.head
	for {
		if nil == node {
			break
		}
		c(node)
		node = node.next
	}
	return nil
}

func (ll *LinkList) reverse() error {
	node := ll.tail
	for {
		if node.prev == nil {
			break
		}
		node.prev, node.next = node.next, node.prev
		node = node.next
	}
	// swap the head and tail
	ll.head, ll.tail = ll.tail, ll.head
	return nil
}
