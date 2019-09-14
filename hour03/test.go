package main

import "fmt"

//方法传入数组，修改元素
func printArr(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	//定义数组
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	printArr(arr1)
	//不同长度，不能传参
	//printArr(arr2)
	printArr(arr3)
	fmt.Println()
	fmt.Println(arr1, arr3)
}
