package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

// Create a new Linked List
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Create a new Node
func NewNode(value int) *Node {
	return &Node{value: value, next: nil}
}

// Append a value into last position (TAIL)
func (ll *LinkedList) Append(value int) {
	newNode := NewNode(value)

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	ll.tail.next = newNode
	ll.tail = newNode
}

// Prepend a value (HEAD)
func (ll *LinkedList) Prepend(value int) {
	newNode := NewNode(value)

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode
}

// Print entire linked list
func (ll *LinkedList) Traverse() {
	if ll.head == nil {
		return
	}

	currentNode := ll.head
	fmt.Println("HEAD: ", ll.head.value)
	fmt.Println("TAIL: ", ll.tail.value)
	fmt.Println("-------")

	for currentNode != nil {
		if currentNode.next != nil {
			fmt.Println("Value: ", currentNode.value, "| Next: ", currentNode.next.value)
		} else {
			fmt.Println("Value: ", currentNode.value, "Next: ", nil)
		}
		currentNode = currentNode.next
	}
	fmt.Println("-------")
}

// Add a value in the middle of the list, after a choosed value
func (ll *LinkedList) AddInTheMiddle(value int, afterValue int) {
	newNode := NewNode(value)

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	if ll.tail.value == afterValue {
		ll.tail.next = newNode
		ll.tail = newNode
		return
	}

	currentNode := ll.head

	for currentNode != nil {
		if currentNode.value == afterValue {
			newNode.next = currentNode.next
			currentNode.next = newNode
			return
		}
		currentNode = currentNode.next
	}
}

// Find an choosed element from the linked list
func (ll *LinkedList) FindElement(value int) bool {
	if ll.head == nil {
		return false
	}

	currentNode := ll.head

	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

// Remove head of the linked list
func (ll *LinkedList) DeleteHead() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.next
}

// Remover tail of the linked list
func (ll *LinkedList) DeleteTail() {
	if ll.head == nil {
		return
	}

	currentNode := ll.head
	nextNode := ll.head.next

	for currentNode != nil {
		if nextNode.next == nil {
			ll.tail = currentNode
			ll.tail.next = nil
			return
		}
		currentNode = currentNode.next
		nextNode = nextNode.next
	}
}

// Remove a choosed element from the middle of the linked list
func (ll *LinkedList) RemoveFromTheMiddle(value int) {
	if ll.head == nil {
		return
	}

	if ll.head.value == value {
		ll.DeleteHead()
		return
	}

	if ll.tail.value == value {
		ll.DeleteTail()
		return
	}

	currentNode := ll.head
	nextNode := ll.head.next

	for currentNode != nil {
		if nextNode.value == value {
			currentNode.next = nextNode.next
			nextNode.next = nil
			return
		}
		currentNode = currentNode.next
		nextNode = nextNode.next
	}
}
