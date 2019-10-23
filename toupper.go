package piscine

func ToUpper(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			res = res + string(r +('A'-'a'))
		} else {
			res = res + string(r)
		}
	}
	return res
}
