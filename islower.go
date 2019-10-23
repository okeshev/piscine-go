package piscine

func IsNumeric(str string) bool {
	for _, r := range str {
		if r >= 'a' && r <= 'z' {
		} else {
			return false
		}
	}
	return true
}
