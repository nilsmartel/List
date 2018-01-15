package list

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	node *Node
	len  uint
}

func New(values ...interface{}) List {
	list := List{nil, 0}

	for _, v := range values {
		list.Push(v)
	}

	return list
}

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

func newNode(value interface{}) *Node {
	node := Node{value, nil}
	return &node
}
