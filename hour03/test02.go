package main

import "fmt"

func main() {
	//1.go语言切片是对原数组的映射，并没有创建一个真正的切片
	// 定义数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//取切片
	s1 := arr[2:]
	//修改值
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println()

	// go语言切片没有取到的位置，可以向后延申，不可向前延申
	s3 := arr[2:6]
	fmt.Println(s3)
	s4 := s3[3:5]
	fmt.Println(s4)

	//容量和大小
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n", s3, len(s3), cap(s3))
}
