package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type User struct {
	conn *websocket.Conn
	msg  chan []byte
}
type Hub struct {
	//用户列表，保存所有用户
	userList map[*User]bool
	//注册chan，用户注册时添加到chan中
	register chan *User
	//注销chan，用户退出时添加到chan中，再从map中删除
	unregister chan *User
	//广播消息，将消息广播给所有连接
	broadcast chan []byte
}

//处理中心处理获取到的信息
func (h *Hub) run() {
	for {
		select {
		//从注册chan中取数据
		case user := <-h.register:
			//取到数据后将数据添加到用户列表中
			h.userList[user] = true
		case user := <-h.unregister:
			//从注销列表中取数据，判断用户列表中是否存在这个用户，存在就删掉
			if _, ok := h.userList[user]; ok {
				delete(h.userList, user)
			}
		case data := <-h.broadcast:
			//从广播chan中取消息，然后遍历给每个用户，发送到用户的msg中
			for u := range h.userList {
				select {
				case u.msg <- data:
				default:
					delete(h.userList, u)
					close(u.msg)
				}
			}
		}
	}
}

//定义一个升级器，将普通的http连接升级为websocket连接
var up = &websocket.Upgrader{
	//定义读写缓冲区大小
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	//校验请求
	CheckOrigin: func(r *http.Request) bool {
		//如果不是get请求，返回错误
		// if r.Method != "GET" {
		// 	fmt.Println("请求方式错误")
		// 	return false
		// }
		//如果路径中不包括chat，返回错误
		// if r.URL.Path != "/ws" {
		// 	fmt.Println("请求路径错误")
		// 	return false
		// }
		//还可以根据其他需求定制校验规则
		return true
	},
}

//初始化处理中心，以便调用
var hub = &Hub{
	userList:   make(map[*User]bool),
	register:   make(chan *User),
	unregister: make(chan *User),
	broadcast:  make(chan []byte),
}

func wsHandle(w http.ResponseWriter, r *http.Request) {
	//通过升级后的升级器得到链接
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("获取连接失败:", err)
		return
	}
	//连接成功后注册用户
	user := &User{
		conn: conn,
		msg:  make(chan []byte),
	}
	hub.register <- user
	defer func() {
		hub.unregister <- user
	}()
	//得到连接后，就可以开始读写数据了
	go read(user)
	write(user)
}
func read(user *User) {

	//从连接中循环读取信息
	for {
		_, msg, err := user.conn.ReadMessage()
		if err != nil {
			fmt.Println("用户退出:", err, user.conn.RemoteAddr().String())
			hub.unregister <- user
			break
		}
		//将读取到的信息传入websocket处理器中的broadcast中，
		fmt.Println("读取数据", msg)
		hub.broadcast <- msg
	}
}
func write(user *User) {
	for data := range user.msg {
		err := user.conn.WriteMessage(1, data)
		fmt.Println("写入成功", string(data[0:]))
		if err != nil {
			fmt.Println("写入错误")
			break
		}
	}
}

func main() {
	fmt.Println("服务启动")
	//后台启动处理器
	go hub.run()
	http.HandleFunc("/ws", wsHandle)           //将chat请求交给wshandle处理
	http.ListenAndServe("127.0.0.1:8888", nil) //开始监听
}
