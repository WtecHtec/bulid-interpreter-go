package main

import (
	"fmt"
	"net/rpc"
)

type Params struct {
	Height, Width int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8088")
	if err != nil {
		fmt.Println("连接失败")
	}
	ret := 0
	err1 := conn.Call("Rect.Area", Params{50, 10}, &ret)
	if err1 != nil {
		fmt.Println("调用失败")
	}
	fmt.Println("面积", ret)
}
