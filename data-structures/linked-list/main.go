package main

import (
	"fmt"
	linkedlist "linked-list/package"
)

func main() {
	ll := linkedlist.NewLinkedList()

	ll.Append(1)
	ll.Append(3)
	ll.Append(4)

	ll.Prepend(0)
	ll.Prepend(6)

	ll.AddInTheMiddle(5, 0)
	ll.AddInTheMiddle(9, 4)
	ll.AddInTheMiddle(10, 9)

	ll.RemoveFromTheMiddle(3)
	ll.RemoveFromTheMiddle(4)

	ll.Traverse()

	elementToFind := 4

	fmt.Println("O elemento", elementToFind, "existe: ", ll.FindElement(elementToFind))
}
