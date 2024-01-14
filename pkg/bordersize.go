package pkg

import "math"

func BorderSize(lenght int) [][]byte {
	hashs := lenght * 4
	lenside := math.Sqrt(float64(hashs))
	board := make([][]byte, int(math.Ceil(lenside)))
	for i := range board {
		board[i] = make([]byte, int(math.Ceil(lenside)))
	}
	return board
}
