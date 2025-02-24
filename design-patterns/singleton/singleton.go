package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		defer lock.Unlock()
		lock.Lock()
		if singleInstance == nil {
			fmt.Println("Creating single instance")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created 1")
		}
	} else {
		fmt.Println("Single instance already created 2")
	}
	return singleInstance
}

func getInstance2() *single {
	if singleInstance == nil {
		defer lock.Unlock()
		lock.Lock()
		singleInstance = &single{}
		fmt.Println("new instance")
	} else {
		fmt.Println("Single instance already created 2")
	}
	return singleInstance
}

func main() {
	for i := 0; i <= 30; i++ {
		go getInstance2()
	}
}
