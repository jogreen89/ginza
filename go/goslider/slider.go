package main

import (
	"fmt"
)

type Board struct {
	x,y int 
	board [3][3]int 
}

func buildBoard() {

	var (
		i = 0
		j = 0
   		k = 0
	    puzzle = new(Board)
	)

	for i < 3 {
		for j < 3 {
			puzzle.board[j][i] = k 
			k++
	
			fmt.Printf("%v ", puzzle.board[j][i])

			j++
		}
		i++
		j = 0
		fmt.Println()
	}
}

func main() {
	// puzzle := new(Board)
	// Check out the output of this object!
	// fmt.Println(puzzle)
	buildBoard()
}

