package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "18628888888"
	matched, err := regexp.MatchString("[0-9-()（）]{7,18}", str)
	fmt.Println(matched, err)
}
