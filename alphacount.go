package piscine

func AlphaCount(str string) int {
	count := 0
	for _, r := range str {
		if (r >= 'A' && r <= 'z') {
			count++
		}
	}
	return count
}
