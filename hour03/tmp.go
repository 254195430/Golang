package main

import (
	"fmt"
	"time"
)

func recevier(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	message := make(chan string, 3)
	message <- "hello"
	message <- "China"
	message <- "world!!!!!!!!!!/n !!!!!!!!!!"
	close(message)
	fmt.Println("请等1秒钟")
	time.Sleep(time.Second * 1)
	recevier(message)

}
