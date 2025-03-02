package main

import (
	"container/list"
	"fmt"
)

/*
Golang STL supports container / list ( Implementation of circular double linked list )

- Stack
- Queue
- Dequeue
- c++ vector ~= slice
*/

func main() {

	// first list :
	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)

	l := list.New()
	mark1 := l.PushBack(12)
	l.PushBack(13)
	mark2 := l.PushBack(15)
	l.PushBack(16)
	l.PushFront(14)
	fmt.Println(l.Back())
	fmt.Println(l.Front())
	fmt.Println(l.Len())
	l.InsertAfter(100, mark2)
	l.InsertBefore(100, mark2)
	//l.Remove(mark)
	l.MoveAfter(mark1, mark2) // move 12 after 15
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ") // Print element value
	}
	fmt.Println()
	l.MoveBefore(mark1, mark2)
	l.MoveToBack(mark1)
	l.MoveToFront(mark1)
	l.PushBackList(l1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ") // Print element value
	}

	// remove all
	//l.Init()

	// remove some starting elements.
	for e := l.Front(); e != nil; {
		nxt := e.Next()
		l.Remove(e)
		e = nxt
	}
}
