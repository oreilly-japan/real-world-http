package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type ToDo struct {
	Task string  `json:"task"`
	Time DueDate `json:"due"`
	Done bool    `json:"done"`
}

type DueDate struct {
	time.Time
}

func (d *DueDate) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(d.Unix()))), nil
}

type ToDoList []ToDo

func (l ToDoList) MarshalJSON() ([]byte, error) {
	tmpList := make([]ToDo, 0, len(l))
	for _, todo := range l {
		if !todo.Done {
			tmpList = append(tmpList, todo)
		}
	}
	return json.Marshal(tmpList)
}

func main() {
	todos := []ToDo{
		ToDo{
			Task: "幼稚園登園",
			Time: DueDate{time.Now()},
			Done: true,
		},
		ToDo{
			Task: "エリクソン研究会に行く",
			Time: DueDate{time.Now()},
			Done: false,
		},
	}
	d, _ := json.Marshal(ToDoList(todos))
	log.Println(string(d))
}
