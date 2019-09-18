package main

import (
	"fmt"
	"time"
)

func Pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C

	}

}

func main() {
	message := make(chan string)
	go Pinger(message)
	for i := 1; i <= 4; i++ {
		msg := <-message
		fmt.Println(msg)
		if i == 3 {
			fmt.Println("第三秒")
			break
		}
	}
}
