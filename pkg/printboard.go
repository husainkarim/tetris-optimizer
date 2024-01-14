package pkg

import "fmt"

func PrintBoard(board [][]byte) {
	for i := range board {
		for _, c := range board[i] {
			if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}
