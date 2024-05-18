package List

import (
	"errors"
	"fmt"
)

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.data)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.data)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}

func (d *DoublyLinkedList[T]) AddElements(index int, value T) error {
	// Check the current size of the list
	currentSize := d.size

	if index > currentSize {
		return errors.New("index out of range")
	}

	// Create a new Node
	newNode := &Node[T]{
		data: value,
	}

	// Check if list is empty
	if currentSize == 0 {
		d.head = newNode
		d.tail = newNode

		d.size += 1

		return nil
	}

	// Check if the head needs to be changed
	if index == 0 {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode

		d.size += 1

		return nil
	}

	// Check if the tail needs to be changed.
	if index == d.size {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode

		d.size += 1

		return nil
	}

	// Insert the node in the middle
	temp := d.head

	for i := 1; i < index; i++ {
		temp = temp.next
	}

	temp.next.prev = newNode
	newNode.next = temp.next
	newNode.prev = temp
	temp.next = newNode

	d.size += 1

	return nil
}
