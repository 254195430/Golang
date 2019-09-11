package main

import (
	"fmt"
)

func SUM(max int) int {
	sum := 0
	for i := 1; i <= max; i += 1 {
		sum += i
	}
	return (sum)
}
func test(num int) int {
	if num == 1 {
		return 1
	}
	return num + test(num-1)
}

func main() {
	fmt.Println(SUM(10))
	fmt.Println(test(100))
}
