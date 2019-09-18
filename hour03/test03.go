package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

func main() {
	//1.方式一
	p := Person{"Mr.Sun", "女"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))

}
