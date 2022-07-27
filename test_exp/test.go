package main

import "fmt"

type Lexer struct {
	input        string
	position     int  // 输入中的当前位置（指向当前字符）
	readPosition int  // 输入中的当前读取位置（在当前字符之后）
	ch           byte // 当前正在检查的字符
}

func test() {
	fmt.Print("hello test")
}
func main() {
	test()
	l := &Lexer{input: "7777"}
	n := l

	// l.ch = 'c'
	n.ch = '6'
	fmt.Print(*l)
	fmt.Print(n)
}
