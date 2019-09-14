package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println(s1)
	s2 := s1[3:6]
	fmt.Println(s2)
	s3 := append(s2, 10)
	fmt.Println(s3)
	fmt.Println(arr)
	s4 := append(s3, 666)
	fmt.Println(s4)
	fmt.Println(arr)

}
