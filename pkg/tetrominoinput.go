package pkg

import (
	"errors"
	"os"
	"strings"
)

func TetrominoInput(file string) ([][][]byte, error) {
	contain, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.New("Invalid Tetromino input file: " + file)
	}
	var primeList []string
	var list [][]string
	lines := strings.Split(string(contain), "\n")
	for i, r := range lines {
		if len(r) != 4 && r != "" {
			return nil, errors.New("ERROR")
		}
		if r != "" {
			primeList = append(primeList, r)
		}
		if r == "" || i == len(lines)-1 {
			list = append(list, primeList)
			primeList = nil
		}
	}
	for _, l := range list {
		if !CheckTetromino(l) {
			return nil, errors.New("ERROR")
		}
	}
	var TetrominoInput [][][]byte
	for _, n := range list {
		TetrominoInput = append(TetrominoInput, ConvertTetrisString(n))
	}
	return TetrominoInput, nil
}
