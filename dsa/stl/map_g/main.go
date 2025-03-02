package main

import (
	"fmt"
	"sort"
)

/*
c++ stl sets and maps...


*/

func main() {
	m := map[int]bool{}
	m[12] = true
	m[14] = true
	m[15] = true
	m[1] = true

	keys := []int{}
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Println("Sorted-> ", key, m[key])
	}

	// if val, ok := m[2]; ok {
	// 	fmt.Print("preset -> ", val)
	// } else {

	// }

}
