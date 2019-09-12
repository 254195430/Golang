package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}
func swap2(c, d int) {
	c, d = d, c
}
func main() {
	a, b := 3, 4
	c, d := 3, 4

	swap(&a, &b)
	fmt.Println(a, b)
	swap2(c, d)
	fmt.Println(c, d)
}
