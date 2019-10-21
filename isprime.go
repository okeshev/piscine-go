package piscine

func IsPrime(nb int) bool {
	if nb < 0 || nb == 1 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	} else {
		if nb%2 == 0 {
			return false
		} else {
			for i := 3; i <nb; i = i + 1 {
				if nb%i == 0 {
					return false
				}
			}
			return true
		}
	}
}
