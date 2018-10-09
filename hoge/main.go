package main

import (
	"fmt"
)

// ID id
type ID int

// Priority priority
type Priority int

// ProcessTask show process task
func ProcessTask(id ID, priority Priority) {
	fmt.Printf("id = %d, priority = %d \n", id, priority)
}

// Task task
type Task struct {
	ID     int
	Detail string
	done   bool
}

// NewTask constructor
func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

// Finish finish task
func (task *Task) Finish() {
	task.done = true
}

func main() {
	var id ID = 3
	var priority Priority = 10
	ProcessTask(id, priority)

	var task = Task{
		ID:     3,
		Detail: "hello",
		done:   true,
	}
	var task2 = new(Task)

	fmt.Println(task, *task2)

	task3 := NewTask(1, "buy a milk")
	fmt.Println(task3, task3.done)
	task3.Finish()
	fmt.Println(task3, task3.done)

}
