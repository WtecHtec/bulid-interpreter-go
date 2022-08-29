// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"
// )

// func test1(ch chan string) {
// 	time.Sleep(time.Second * 5)
// 	ch <- "test1"
// }
// func test2(ch chan string) {
// 	time.Sleep(time.Second * 2)
// 	ch <- "test2"
// }
// func bufioDemo() {
// 	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
// 	fmt.Print("请输入内容：")
// 	text, _ := reader.ReadString('\n') // 读到换行
// 	text = strings.TrimSpace(text)
// 	fmt.Printf("%#v\n", text)
// }
// func main() {
// 	// // 2个管道
// 	// output1 := make(chan string)
// 	// output2 := make(chan string)
// 	// // 跑2个子协程，写数据
// 	// go test1(output1)
// 	// go test2(output2)
// 	// // 用select监控
// 	// for {
// 	// 	select {
// 	// 	case s1 := <-output1:
// 	// 		fmt.Println("s1=", s1)
// 	// 	case s2 := <-output2:
// 	// 		fmt.Println("s2=", s2)
// 	// 	}
// 	// }
// 	// var (
// 	// 	name    string
// 	// 	age     int
// 	// 	married bool
// 	// )
// 	// fmt.Scan(&name, &age, &married)
// 	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
// 	// bufioDemo()
// 	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("open log file failed, err:", err)
// 		return
// 	}
// 	log.SetOutput(logFile)
// 	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
// 	log.Println("这是一条很普通的日志。")
// 	log.SetPrefix("[小王子]")
// 	log.Println("这是一条很普通的日志。")
// }

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //注意是字符
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	strconv.ParseInt()
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
