package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Task struct {
	Id        string
	Status    int
	Content   string
	Task_time string
}

type ToDoData struct {
	List []Task
}
type Handle struct {
	List []Task
}

func (h *Handle) seachAll() []Task {
	// 打开json文件
	jsonFile, err := os.Open("data.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 要记得关闭
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	todoValue := string(byteValue)
	if todoValue == "" {
		return nil
	}
	var toDoData ToDoData
	// 解析字符串为Json
	json.Unmarshal([]byte(todoValue), &toDoData)
	h.List = toDoData.List
	return toDoData.List
}

func (h *Handle) updateJson() bool {
	var toDoData ToDoData
	toDoData.List = h.List
	// 从新生成json字符串查看真实结构
	jsons, _ := json.Marshal(toDoData)
	todoValue := string(jsons)
	err := ioutil.WriteFile("data.json", []byte(todoValue), 0755)
	if err != nil {
		log.Println("读取数据失败")
		fmt.Println("WriteFile failed")
		return false
	}
	return true
}
func (h *Handle) getTaskOne(id string) Task {
	for _, item := range h.List {
		if item.Id == id {
			return item
		}
	}
	return Task{}
}

func (h *Handle) modify(task *Task, tasks []Task, opt string) []Task {
	var index int
	for i, item := range tasks {
		if item.Id == task.Id {
			index = i
			break
		}
	}
	if opt == "del" {
		log.Println("删除一条记录", task.Id)
		tasks = append(tasks[:index], tasks[index+1:]...)
	}

	if opt == "update" {
		tasks[index] = *task
		log.Println("更新一条记录", task.Id, "-----", *task)
	}

	if opt == "add" {
		tasks = append(tasks, *task)
		log.Println("添加一条记录", task.Id, "-----", *task)
	}
	h.List = tasks
	return tasks
}

// func main() {
// var tasks []Task
// tasks := seachAll()
// if tasks == nil {
// 	return
// }
// updateJson(tasks)
// task := &Task{
// 	Id:      "1",
// 	Content: "什么都没做",
// 	Status:  0,
// }
// tasks = modify(task, tasks, "add")
// fmt.Print(tasks)
// var handle Handle
// handle.seachAll()
// fmt.Print(handle.List)
// }
