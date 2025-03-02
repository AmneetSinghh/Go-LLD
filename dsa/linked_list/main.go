package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

type SingleLinkedList struct {
	head *Node // it takes addres of node, where node is located in memory block
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (sl *SingleLinkedList) insertAtBegin(val int) {
	node := NewNode(val)
	node.next = sl.head
	sl.head = node
}

func (sl *SingleLinkedList) deleteAtEnd() {

	if sl.head == nil {
		return
	}
	if sl.head.next == nil {
		sl.head = nil
		return
	}
	temp := sl.head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
}

func (sl *SingleLinkedList) insertAtEnd(val int) {

	if sl.head == nil {
		sl.head = NewNode(val)
		return
	}
	temp := sl.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = NewNode(val)
}

func (sl *SingleLinkedList) printList() {
	temp := sl.head
	for temp != nil {
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println("")
}

func main() {
	linkedList := NewSingleLinkedList()
	fmt.Printf("sll: %p\n", linkedList) // Address of SingleLinkedList
	linkedList.insertAtBegin(10)
	linkedList.insertAtBegin(30)
	linkedList.insertAtBegin(40)
	linkedList.printList()
	linkedList.deleteAtEnd()
	linkedList.printList()
	linkedList.insertAtEnd(190)
	linkedList.insertAtEnd(200)
	linkedList.printList()
}
