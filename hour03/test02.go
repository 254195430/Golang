package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Oh I do like to beside the seaside"
	fmt.Println(strings.ToUpper(s))
	s1 := strings.Replace(s, "beside", "haha", -1)
	fmt.Println(strings.ToUpper(s1))
	s2 := strings.Index(s, "the")
	fmt.Println(s2)
}
