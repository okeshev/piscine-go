package piscine
import "fmt"


func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
func UltimateDivMod(a *int, b *int) {
	m := *a / *b
	n := *a % *b
	*a = m
	*b = n
}
