package linkedlist

import (
	"fmt"
)


type Node struct {
	value int
	next *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type LinkedList struct {
	head *Node
	len int
}

func (l *LinkedList) Add(value int) {
	newNode := Node{value: value}
	l.len ++
	if l.head == nil {
		l.head = &newNode
		return
	}
	newNode.next = l.head
	l.head = &newNode

}

func (l * LinkedList) String() {
	if l.head == nil {
		return
	}
	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (l *LinkedList) Get(value int) *Node{
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.value == value {
			return curr
		}
	}
	return nil
}

func (l *LinkedList) Delete(value int) {
	if l.len == 0 {
		return
	}
	l.len --
	if l.head.value == value {
		l.head = l.head.next
		return
	}

	prev := l.head
	for prev.next.value != value {
		if prev.next.next == nil {
			return
		}
		prev = prev.next
	}
	prev.next = prev.next.next
}