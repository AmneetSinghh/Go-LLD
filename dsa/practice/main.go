package main

import "fmt"

func main() {

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	// Print characters as runes (characters)
	for _, row := range board {
		for _, ch := range row {
			cur := ch
			fmt.Printf("%c ", cur) // Output: A B C E \n S F C S \n A D E E
		}
		fmt.Println()
	}

	// Print ASCII values (since byte stores ASCII)
	fmt.Println(board[0][0]) // Output: 65 (ASCII for 'A')

	val := []int{1, 2, 3, 3, 4, 5}
	fmt.Println(val[1:])

}
