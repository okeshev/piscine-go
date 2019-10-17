package piscine

func BasicAtoi(s string) int {
	change := []rune(s)
	var mean int = 0
	for _, r := range change {
		mean *= 10
		if r == '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' {
			mean += 0, mean += 1, mean += 2, mean += 3, mean += 4, mean += 5, mean += 6, mean += 7, mean += 8, mean += 9
		}
	}
	return mean
}