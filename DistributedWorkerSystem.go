package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID int
}

func (t Task) Process() {
	fmt.Printf("task %d start\n", t.ID)
	time.Sleep(time.Second * 2)
	fmt.Printf("task %d end\n", t.ID)
}
func worker(workerId int, tasks <-chan Task, done chan<- bool) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d \n", workerId, task.ID)
		task.Process()
		fmt.Printf("Worker %d ended task %d \n", workerId, task.ID)
	}
	done <- true
}

func main() {
	tasks := make(chan Task, 10)
	done := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go worker(i, tasks, done)
	}

	for i := 0; i < 10; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)
	for i := 0; i < 3; i++ {
		<-done
	}
}
