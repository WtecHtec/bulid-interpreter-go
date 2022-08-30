package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Params struct {
	Height, Width int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

func (r *Rect) Permieter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	rect := new(Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server")
}
