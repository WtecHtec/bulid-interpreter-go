package main

import (
	"fmt"
)

type Lexer struct {
	input        string
	position     int  // 输入中的当前位置（指向当前字符）
	readPosition int  // 输入中的当前读取位置（在当前字符之后）
	ch           byte // 当前正在检查的字符
}

func test() {
	fmt.Print("hello test")
}

//定义函数类型
type handler func(name string) int

//实现函数类型方法
func (h handler) add(name string) int {
	return h(name) + 10
}

type Retriever interface {
	Get(url string) string
	Add(num int) int
}

type Retrievers struct {
	name string
}

func (r *Retrievers) Get(url string) string {
	return string(url)
}

func (r *Retrievers) Add(num int) int {
	return num
}

func main() {
	test()
	l := &Lexer{input: "7777"}
	n := l

	// l.ch = 'c'
	n.ch = '6'
	fmt.Print(l)
	fmt.Print(n)

	//注意要成为函数对象必须显式定义handler类型
	var my handler = func(name string) int {
		return len(name)
	}
	fmt.Print(my("777"))
	fmt.Println(my.add("taozs")) //调用函数对象的方法

	var r Retriever
	r = &Retrievers{name: "测试"}
	// fmt.Println(download(r))

	// Type Assertion
	if retriever, ok := r.(*Retrievers); ok {
		fmt.Printf("%T %v\n", retriever, retriever.name)
	}
}
