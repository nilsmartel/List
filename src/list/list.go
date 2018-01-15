package list

//	Node
//	Don't use this outside of list.go
type Node struct {
	value interface{}
	next  *Node
}

// List
// Standart Linked List
type List struct {
	node *Node
	len  uint
}

// New
// Create a new (Linked) List from input parameter
// @param values Variable Amount of values to be made into a linked list
func New(values ...interface{}) List {
	list := List{nil, 0}

	for _, v := range values {
		list.Push(v)
	}

	return list
}

// Push
// push Value at the end of a list
func (list *List) Push(value interface{}) {
	if list.node == nil {
		list.node = newNode(value)
	} else {
		current := list.node

		for current.next != nil {
			current = current.next
		}

		current.next = newNode(value)
	}

	list.len += 1
}

// Pop
// Remove and Return first Element of a List.
// if list is empty this will return nil
func (list *List) Pop() interface{} {
	if list.len == 0 {
		return nil
	}

	node := list.node.next
	value := list.node.value
	list.node = node
	list.len -= 1
	return value
}

// IsEmpty
// returns true if a List is empty
func (list *List) IsEmpty() bool {
	return list.len == 0
}

// Contains
// Check if a List contains a certain Element
func (list *List) Contains(value interface{}) bool {
	current := list.node
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}

	return false
}

func newNode(value interface{}) *Node {
	node := Node{value, nil}
	return &node
}
