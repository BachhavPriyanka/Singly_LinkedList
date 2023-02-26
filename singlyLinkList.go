// Define a method to insert a node at a specific index
package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) AddAtMiddle(value int, pos int) {
	newNode := &Node{Data: value}

	if pos == 1 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}
	current := ll.Head
	for i := 0; i < pos-2; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
}

func (ll *LinkedList) RemoveLast() {
	current := ll.Head

	if ll.Head == nil {
		return
	}

	if ll.Head.Next == nil {
		ll.Head = nil
		return
	}

	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
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

func (ll *LinkedList) RemoveFront() {
	if ll.Head == nil {
		fmt.Println("Empty")
	}
	ll.Head = ll.Head.Next

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
	//Adding Data into list
	list.AddAtLast(1)
	list.AddAtLast(2)
	list.AddAtLast(3)
	//Adding Data at first
	//list.AddAtFirst(11)
	//Removing Data from from
	//list.RemoveFront()
	//Removing Data from last
	//list.RemoveLast()
	//Add at Middle
	list.AddAtMiddle(22, 1)
	list.Display()
}
