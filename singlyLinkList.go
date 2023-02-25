// Define a method to add a node to the end of the linked list
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

func (ll *LinkedList) AddAtLast(value int) {
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

func (ll *LinkedList) AddAtFirst(value int) {
	newNode := &Node{Data: value}

	if ll.Head == nil {
		ll.Head = newNode
	}
	newNode.Next = ll.Head
	ll.Head = newNode
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
	list.AddAtLast(1)
	list.AddAtLast(2)
	list.AddAtLast(3)
	//list.RemoveFront()
	list.AddAtFirst(11)
	list.Display()
}
