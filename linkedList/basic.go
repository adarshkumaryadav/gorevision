// creation of linked list ds

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type MyLinkedList struct {
	head  *Node
	count int
}

func (list *MyLinkedList) insert(data int) {
	// create the node first in the heap memory
	// below commented code creates the Node in the stack memory although GC later moves this to heap
	// node := Node{data: data, next: nil}
	node := &Node{data: data}
	current := list.head
	list.count++
	if current == nil {
		// let it store the very first entry
		list.head = node
		fmt.Println("Insertion completed, list was empty")
		return
	}
	// if there is already any existing node/entry
	// traverse till the last element is not found i.e. last element next == nil

	for current.next != nil {
		current = current.next
	}
	current.next = node

	fmt.Println("Insertion completed")
}

func (list *MyLinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println("Reading Done of ", list.count, " element")
}

func main() {
	fmt.Println("Linked list Data structure")
	list := MyLinkedList{}

	list.Print()
	list.insert(1)
	list.insert(2)
	list.Print()
	list.insert(3)
	list.insert(4)
	list.Print()

}
