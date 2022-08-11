package main

import (
	"fmt"
	"reflect"
	"testing"
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

func mpSort() {
	dataList := []int{2, 4, 6, 1, 8, 9, 7}
	arrLen := len(dataList)
	for i := 0; i < arrLen; i++ {
		for j, num := range dataList[:arrLen-1-i] {
			if num > dataList[j+1] {
				temp := num
				dataList[j] = dataList[j+1]
				dataList[j+1] = temp
			}
		}
	}
	fmt.Println(dataList)
}

type Stack struct {
	value    []int
	stackLen int
}

// func (s *Stack) refresh() {
// 	newValue := make([]int, s.stackLen)
// 	for i, num := range s.value {
// 		if i < s.stackLen {
// 			newValue[i] = num
// 		}
// 	}
// 	s.value = newValue
// }

func (s *Stack) push(value int) {
	s.stackLen = s.stackLen + 1
	s.value = append(s.value, value)
	s.value[s.stackLen-1] = value
}

func (s *Stack) pop() int {
	value := s.value[s.stackLen-1]
	s.stackLen = s.stackLen - 1
	s.value = s.value[:s.stackLen]
	return value
}

type DiyStack []int

func (s *DiyStack) push(newValue int) {
	*s = append(*s, newValue)
}
func (s *DiyStack) pop() (int, bool) {
	sLen := len(*s)
	fmt.Println(sLen)
	if sLen <= 0 {
		return -1, false
	}
	num := (*s)[sLen-1]
	(*s) = (*s)[:sLen-1]
	fmt.Println(*s)
	return num, true
}

type Node struct {
	value       int
	left, right *Node
}

func main() {
	// test()
	// l := &Lexer{input: "7777"}
	// n := l

	// l.ch = 'c'
	// n.ch = '6'
	// fmt.Print(l)
	// fmt.Print(n)

	//注意要成为函数对象必须显式定义handler类型
	// var my handler = func(name string) int {
	// 	return len(name)
	// }
	// fmt.Print(my("777"))
	// fmt.Println(my.add("taozs")) //调用函数对象的方法

	// var r Retriever
	// r = &Retrievers{name: "测试"}
	// fmt.Println(download(r))

	// Type Assertion
	// if retriever, ok := r.(*Retrievers); ok {
	// 	fmt.Printf("%T %v\n", retriever, retriever.name)
	// }

	mpSort()
	diyStack := &DiyStack{}
	diyStack.push(2)
	diyStack.push(8)
	diyStack.pop()
	// ds := *diyStack
	fmt.Println(diyStack)
	testStack := &Stack{}
	testStack.push(9)
	testStack.push(78)
	testStack.push(78)
	fmt.Println(testStack)
	num := testStack.pop()
	fmt.Println(num)
	// fmt.Println(testStack)
	rootNode := &Node{
		value: 1,
	}
	leftNode := &Node{
		value: 2,
	}
	rightNode := &Node{
		value: 3,
	}
	leftNode.right = rightNode
	rootNode.left = leftNode
	fmt.Println(rootNode.left.right)

}

func TestIF(t *testing.T) {
	x := 3
	if n := "abc"; x > 0 { // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
		t.Logf("====%v", n[2])
	} else if x < 0 { // 注意 else if 和 else 左大括号位置。
		println(n[1])
	} else {
		println(n[0])
	}
	m_type(x, t)
	typeStr := reflect.TypeOf(x)
	t.Log(typeStr)
}

func m_type(i interface{}, t *testing.T) {
	switch i.(type) {
	case string:
		//...
		t.Log("string")
	case int:
		//...
		t.Log("int")

	}
	return
}

func TestRange(t *testing.T) {
	// str := "abc"
	// for i, s := range str {
	// 	t.Logf("i===%v,value===%v", i, s)
	// }

	// arry := [...]int{1, 2, 4, 5, 6}
	// for i, s := range arry {
	// 	t.Logf("i===%v,value===%v", i, s)
	// }

	// mapData := make(map[string]int)
	// mapData["a"] = 97
	// mapData["b"] = 98
	// mapData["c"] = 99
	// for i, s := range mapData {
	// 	t.Logf("i===%v,value===%v", i, s)
	// }

	grader := 90
	switch grader {
	case 90:
		t.Log("90分")
		t.Log("优秀")
	default:
		t.Log("一般")
	}

	var a int
	b := &a
	(*b) = 8
	t.Log(a)
}

func param(value interface{}, t *testing.T) {
	t.Log(value)
}
func TestParams(t *testing.T) {
	value := 5
	param(value, t)
}

type Animia interface {
	say()
}

type Dog struct {
	name string
}

func (d *Dog) say() {
	fmt.Println(d.name)
}

func (d *Dog) run() {
	fmt.Println(d.name, "running")
}

type Cat struct {
	name string
}

func TestInterface(t *testing.T) {
	var d Animia
	dog := &Dog{name: "狗子"}
	d = dog
	d.say()
	dog.run()
	td, ok := d.(*Dog)
	fmt.Println(td.name, ok)
}
