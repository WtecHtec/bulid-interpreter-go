package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func TCPClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:4004")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("发送失败")
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接收服务器失败")
		}
		fmt.Println("服务器：", string(buf[:n]))
	}
}

func UDPClient() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 4004,
	})
	if err != nil {
		fmt.Println("连接服务失败")
		return
	}
	defer socket.Close()
	sendData := []byte("你好 udp")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("发送数据失败")
		return
	}
	data := make([]byte, 4096)
	n, addr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("接收数据失败")
		return
	}
	fmt.Println("接收数据", string(data[n:]), addr)

}
func main() {
	UDPClient()
}
