package piscine

func UltimateDivMod(a *int, b *int) {
	m := *a / *b
	n := *a % *b
	*a = m
	*b = n
}
