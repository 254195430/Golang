package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:38888")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ClientConn(conn)
	}
}

func ClientConn(conn net.Conn) {
	defer conn.Close()
	ipAddr := conn.RemoteAddr().String()
	fmt.Println("连接成功", ipAddr)
	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := bytes[:n]
		fmt.Printf("接收到的数据来自[%s]==>[%d]:%s\n", ipAddr, n, string(result))
		if "exit" == string(result) {
			fmt.Println("程序退出")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(result))))
	}
}
