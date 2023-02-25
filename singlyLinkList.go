// Define a method to remove the front node from the linked list
package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) AddAtFirst(value int) {
	newNode := &Node{Data: value}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) RemoveFront() (int, error) {
	if ll.Head == nil {
		return 0, errors.New("linked list is empty")
	}

	data := ll.Head.Data
	ll.Head = ll.Head.Next
	return data, nil

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
	list.RemoveFront()
	list.RemoveFront()

	list.Display()
}
