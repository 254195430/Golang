package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myint int
	var i myint = 100
	fmt.Println(reflect.TypeOf(i))
}
