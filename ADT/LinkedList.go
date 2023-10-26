package ADT

import "fmt"

type Node[T AlNum] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T AlNum] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
}

func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}

func (l *LinkedList[T]) Remove(value T) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
			if currentNode.Next == nil {
				l.Tail = currentNode
			}
			return
		}
		currentNode = currentNode.Next
	}
}

func (l *LinkedList[T]) Contains(value T) bool {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (l *LinkedList[T]) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Print(currentNode.Value, " ")
		currentNode = currentNode.Next
	}
	fmt.Println()
}
