package piscine

func Swap(a *int, b *int) {
	cont := *a
	*a = *b
	*b = cont
}