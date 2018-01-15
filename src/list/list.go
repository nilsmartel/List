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
}

// New
// Create a new (Linked) List from input parameter
// @param values Variable Amount of values to be made into a linked list
func New(values ...interface{}) List {
	list := List{nil}

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
}

// Pop
// Remove and Return first Element of a List.
// if list is empty this will return nil
func (list *List) Pop() interface{} {
	if list.node == nil {
		return nil
	}

	node := list.node.next
	value := list.node.value
	list.node = node
	return value
}

// IsEmpty
// returns true if a List is empty
func (list *List) IsEmpty() bool {
	return list.node == nil
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

// Array
// Create an Array from a List
func (list *List) Array() []interface{} {
	arr := []interface{}
	current := list.node

	for current != nil {
		append(arr, current.value)
		current = current.next
	}
}


// Len
// returns Length (Size) of a List
func (list *List) Len() int {
	length := 0
	current := list.node

	for current != nil {
		length += 1
		current = current.next
	}

	return length
}

func newNode(value interface{}) *Node {
	node := Node{value, nil}
	return &node
}
