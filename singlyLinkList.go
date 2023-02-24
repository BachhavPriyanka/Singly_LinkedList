// Define a method to add a node to the front of the linked list
package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) AddAtFirst(value int) {
	newNode := &Node{Data: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *LinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{}
	list.AddAtFirst(1)
	list.AddAtFirst(2)
	list.AddAtFirst(3)
	list.Display()
}
