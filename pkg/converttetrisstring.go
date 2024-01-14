package pkg

func ConvertTetrisString(tetromisRows []string) [][]byte {
	// Initialize the valid rows and columns arrays.
	validRows := []int{}
	validColumns := []int{}

	// Iterate over the rows of the board.
	for i := 0; i < len(tetromisRows); i++ {
		isRowValid := false
		for j := 0; j < len(tetromisRows[0]); j++ {
			if tetromisRows[i][j] == '#' {
				isRowValid = true
				break
			}
		}
		if isRowValid {
			validRows = append(validRows, i)
		}
	}

	// Iterate over the columns of the board.
	for j := 0; j < len(tetromisRows[0]); j++ {
		isColumnValid := false
		for i := 0; i < len(tetromisRows); i++ {
			if tetromisRows[i][j] == '#' {
				isColumnValid = true
				break
			}
		}
		if isColumnValid {
			validColumns = append(validColumns, j)
		}
	}
	// Create a new 3D byte array of size len(validRows) x len(validColumns).
	newTetrominoes := make([][]byte, len(validRows))
	for i := range newTetrominoes {
		newTetrominoes[i] = make([]byte, len(validColumns))
	}

	// Iterate over the valid rows of the original board.
	for i, row := range validRows {
		// Iterate over the valid columns of the original board.
		for j, column := range validColumns {
			// Copy the value from the original board to the new board.
			newTetrominoes[i][j] = tetromisRows[row][column]
		}
	}
	return newTetrominoes
}
