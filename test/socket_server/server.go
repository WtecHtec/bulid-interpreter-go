package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		revStr := string(buf[:n])
		fmt.Println("客户端发送：", revStr)
		conn.Write([]byte("server"))
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:4004")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务器启动，等待连接")
	for {
		// 等待
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Println("连接ing....")
		go process(conn) // 启动一个goroutine处理连接
	}
}
