package main

//Copied from the internet
import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
	tail *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{head: nil, tail: nil}
}

func (cll *CircularLinkedList) IsEmpty() bool {
	return cll.head == nil
}

func (cll *CircularLinkedList) AddNode(data int) {
	newNode := &Node{data: data, next: nil}

	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = cll.head
	} else {
		cll.tail.next = newNode
		cll.tail = newNode
		newNode.next = cll.head
	}
}

func (cll *CircularLinkedList) Traverse() {
	if cll.IsEmpty() {
		fmt.Println("Circular Linked List is empty")
	} else {
		current := cll.head
		for {
			fmt.Printf("%d -> ", current.data)
			current = current.next
			if current == cll.head {
				break
			}
		}
		fmt.Printf("%d\n", cll.head.data)
	}
}

func main() {
	cll := NewCircularLinkedList()
	fmt.Println("The nodes of circular linked list are:")
	cll.AddNode(1)
	cll.AddNode(2)
	cll.AddNode(3)
	cll.AddNode(4)
	cll.AddNode(5)
	cll.Traverse()
}
