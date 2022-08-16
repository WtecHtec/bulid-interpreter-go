package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var handle Handle

func pageViews(w http.ResponseWriter, r *http.Request) {
	tasks := handle.seachAll()
	// 解析指定文件生成模板对象
	htmlByte, err := ioutil.ReadFile("./view/index.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 自定义一个夸人的模板函数
	formatStatus := func(arg int) (string, error) {
		if arg == 0 {
			return "进行中", nil
		}
		return "已完成", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的formatStatus函数
	tmpl, _ := template.New("views").Funcs(template.FuncMap{"formatStatus": formatStatus}).Parse(string(htmlByte))

	tmpl.Execute(w, tasks)
}
func finishTask(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q) != 0 {
		task := handle.getTaskOne(string(q["id"][0]))
		task.Status = 1
		handle.modify(&task, handle.List, "update")
		handle.updateJson()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func delTask(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q) != 0 {
		task := handle.getTaskOne(string(q["id"][0]))
		task.Status = 1
		handle.modify(&task, handle.List, "del")
		handle.updateJson()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func addTask(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q) != 0 {
		task := handle.getTaskOne(string(q["id"][0]))
		task.Status = 1
		handle.modify(&task, handle.List, "del")
		handle.updateJson()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func editTask(w http.ResponseWriter, r *http.Request) {
	met := r.Method
	if met == "GET" {
		q := r.URL.Query()
		tmpl, err := template.ParseFiles("./view/edit.html")
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}
		var task Task
		if len(q) != 0 {
			task = handle.getTaskOne(string(q["id"][0]))
		}
		tmpl.Execute(w, task)
		return
	} else if met == "POST" {

		r.ParseForm()
		// var jsonValue string = "{"
		// for k, v := range r.Form {
		// 	jsonValue = jsonValue + "\"" + k + "\":"
		// 	fmt.Println("key is: ", k)
		// 	fmt.Println("val is: ", v)
		// }
		// // 通过json解析器解析参数
		// json.NewDecoder([]byte(r.FormValue)).Decode(param)
		value, err := strconv.Atoi(r.PostFormValue("Status"))
		if err != nil {
			value = 0
		}
		var opt string = "update"
		task := &Task{
			Id:        r.PostFormValue("Id"),
			Content:   r.PostFormValue("Content"),
			Status:    value,
			Task_time: r.PostFormValue("Task_time"),
		}
		if task.Id == "" {
			opt = "add"
			task.Id = strconv.FormatInt(time.Now().Unix(), 10)
		}
		handle.modify(task, handle.List, opt)
		handle.updateJson()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func main() {
	fmt.Println("127.0.0.1:8900服务启动")
	handle = Handle{}
	http.HandleFunc("/", pageViews)
	http.HandleFunc("/finish", finishTask)
	http.HandleFunc("/del", delTask)
	http.HandleFunc("/edit", editTask)
	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}

}
