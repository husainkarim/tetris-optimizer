package main

import (
	"fmt"
	"os"
	"tetris-optimizer/pkg"
)

var BOARD [][]byte

type Tetris struct {
	Form [][]byte
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error: unexpected number of arguments, it should be one only!")
		os.Exit(0)
	}
	List, err := pkg.TetrominoInput(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for _, v := range List {
		if !pkg.CheckFormat(v) {
			fmt.Println("ERROR")
			os.Exit(0)
		}
	}
	var TetrisList []Tetris
	var temp Tetris
	for _, v := range List {
		temp.Form = v
		TetrisList = append(TetrisList, temp)
	}

	lenght := len(TetrisList)
	BOARD = pkg.BorderSize(lenght)
	PlacePosition(0, lenght, TetrisList)
}

func PlacePosition(psc, size int, tetrominoes []Tetris) {
	for y := 0; y < len(BOARD); y++ {
		for x := 0; x < len(BOARD); x++ {
			if CheckPosition(y, x, psc, tetrominoes) {
				if y == len(BOARD)-1 || psc == len(tetrominoes)-1 {
					pkg.PrintBoard(BOARD)
					os.Exit(0)
				} else {
					PlacePosition(psc+1, size, tetrominoes)
				}
				RemovePosition(y, x, psc, tetrominoes)
			}
		}
	}

	if psc == 0 {
		increaseSize := size + 1
		BOARD = pkg.BorderSize(increaseSize)
		PlacePosition(0, increaseSize, tetrominoes)
	}
}

func CheckPosition(y, x, psc int, tetrominoesList []Tetris) bool {
	for i := 0; i < len(tetrominoesList[psc].Form); i++ {
		if len(tetrominoesList[psc].Form)+y > len(BOARD) || len(tetrominoesList[psc].Form[i])+x > len(BOARD) {
			return false
		}
	}

	for a := y; a < (len(tetrominoesList[psc].Form) + y); a++ {
		for b := x; b < (len(tetrominoesList[psc].Form[a-y]) + x); b++ {
			if tetrominoesList[psc].Form[a-y][b-x] == 46 {
				continue
			}
			if BOARD[a][b] == 0 {
				BOARD[a][b] = byte(65 + psc)
			} else {
				RemovePosition(y, x, psc, tetrominoesList)
				return false
			}
		}
	}

	return true
}

func RemovePosition(y, x, psc int, tetrominoesList []Tetris) {
	for a := y; a < (len(tetrominoesList[psc].Form) + y); a++ {
		for b := x; b < (len(tetrominoesList[psc].Form[a-y]) + x); b++ {
			if (tetrominoesList[psc].Form[a-y][b-x]) == 46 {
				continue
			}
			if BOARD[a][b] == byte(65+psc) {
				BOARD[a][b] = 0
			}
		}
	}
}
