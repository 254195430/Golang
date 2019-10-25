package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:38888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	bytes := make([]byte, 1024)
	for {
		fmt.Println("请输入内容:")
		fmt.Scan(&bytes)
		fmt.Printf("您要发送的内容是%s\n", string(bytes))

		conn.Write(bytes)
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := bytes[:n]
		fmt.Printf("接收到数据[%d]:%s\n", n, string(result))
	}
}
