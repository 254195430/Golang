package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 16, 32)
	return m.Name + "," + r
}
func main() {
	m := Movie{
		"haha",
		3.1415926,
	}
	fmt.Println(m.summary())
}
