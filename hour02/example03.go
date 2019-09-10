package main

import (
	"fmt"
)

func addition(x string, y string) string {
	return x + y
}

func main() {
	var s string = "three"
	fmt.Println(addition("1", s))
}
