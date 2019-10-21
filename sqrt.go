package main

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	for i := 1; ; i++ {
		if nb == i*i {
			return i
		}
		if nb < i*i {
			return 0
		}
	}
}
