package piscine

func StrLen(str string) int {
	aString := []rune(str)
	ch := 0
	for range aString {
		ch++
	}
	return ch
}
