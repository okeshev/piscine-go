package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	a := []rune(toFind)
	b := 0
	c := 0
	for range str {
		b++
	}
	for range a {
		c++
	}
	for i := 0; i <= b-c; i++ {
		if toFind == s[i:i+c] {
			return (i)
		}
	}
	return -1
}
