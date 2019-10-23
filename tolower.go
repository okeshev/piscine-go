package piscine

func ToLower(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			res = res + string(r+('a'-'A'))
		} else {
			res = res + string(r)
		}
	}
	return res
}
