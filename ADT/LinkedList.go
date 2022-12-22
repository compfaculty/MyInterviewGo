package ADT

type Node struct {
	key  string
	next *Node
}
type LinkedList struct {
	head  *Node
	count int
}

func (l *LinkedList) Prepend(n *Node) {
	currentHead := l.head
	n.next, l.head = currentHead, n
	//l.head.next = currentHead
	l.count++

}
func (l *LinkedList) Insert(v string) {
	if !l.Contains(v) {
		node := &Node{key: v}
		node.next = l.head
		l.head = node
		l.count++
	}

}
func (l *LinkedList) Delete(v string) *Node {
	var ret *Node
	if l.head.key == v {
		ret = l.head
		l.head = l.head.next
		l.count--
		return ret
	}

	prev := l.head
	for prev.next != nil {
		if prev.next.key == v {
			ret = prev.next
			prev.next = prev.next.next
			l.count--
			return ret
		}
		prev = prev.next
	}
	return nil
}

func (l *LinkedList) Get(v string) *Node {
	current := l.head
	for current != nil {
		if current.key == v {
			return current
		}
		current = current.next
	}
	return nil
}

func (l *LinkedList) Contains(v string) bool {
	ret := l.Get(v)
	if ret == nil || ret.key != v {
		return false
	}
	return true
}
