package pkg

func CheckTetromino(input []string) bool {
	count := 0
	allowed := map[rune]bool{'.': true, '#': true}
	for _, character := range input {
		for _, r := range character {
			if !allowed[r] {
				return false
			}
		}
		count++
	}
	return count == 4
}
