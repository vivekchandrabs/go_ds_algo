package main

import (
	"fmt"

	db "./List"
)

func main() {
	// Create new Doubly LinkedList
	dl := &db.DoublyLinkedList[string]{}

	err := dl.AddElements(0, "A")
	if err != nil {
		fmt.Println(err)
	}

	err = dl.AddElements(0, "D")
	if err != nil {
		fmt.Println(err)
	}

	err = dl.AddElements(1, "B")
	if err != nil {
		fmt.Println(err)
	}

	err = dl.AddElements(3, "C")
	if err != nil {
		fmt.Println(err)
	}

	// Print the list
	fmt.Println(dl.PrintForward())

	// Print the list in reverse
	fmt.Println(dl.PrintReverse())

	err = dl.RemoveElement(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("After removing element at index 1")

	// Print the list
	fmt.Println(dl.PrintForward())

	fmt.Println("After removing element at index 0")

	err = dl.RemoveElement(0)
	if err != nil {
		fmt.Println(err)
	}

	// Print the list
	fmt.Println(dl.PrintForward())

	fmt.Println("After removing element at index 1")

	err = dl.RemoveElement(1)
	if err != nil {
		fmt.Println(err)
	}

	// Print the list
	fmt.Println(dl.PrintForward())
}
