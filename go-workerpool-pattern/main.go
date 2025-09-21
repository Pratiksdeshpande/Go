package main

import (
	"fmt"
)

func main() {
	// comment out one of the following function calls to test either implementation
	WorkerPoolWithOneTypeOfTask()
	WorkerPoolWithMultipleTypeOfTasks()
}

func WorkerPoolWithOneTypeOfTask() {

	//create 20 tasks of one type
	tasks := make([]Task, 20)

	for i := 0; i < 20; i++ {
		tasks[i] = Task{Id: i + 1}
	}

	//create a worker pool with 5 concurrent workers
	wp := WorkerPool{
		Tasks:       tasks,
		Concurrency: 6,
	}

	wp.Run()
	fmt.Println("All tasks completed.")
}

func WorkerPoolWithMultipleTypeOfTasks() {

	//create multiple tasks of type EmailTask and ImageProcessing
	multiTask := []MultiTask{
		&EmailTask{EmailId: "abc", Subject: "hello abc", Message: "message 1"},
		&ImageProcessingTask{"ABC"},
		&EmailTask{EmailId: "def", Subject: "hello def", Message: "message 2"},
		&ImageProcessingTask{"DEF"},
		&EmailTask{EmailId: "ghi", Subject: "hello ghi", Message: "message 3"},
		&ImageProcessingTask{"GHI"},
		&EmailTask{EmailId: "jkl", Subject: "hello jkl", Message: "message 4"},
		&ImageProcessingTask{"JKL"},
		&EmailTask{EmailId: "mno", Subject: "hello mno", Message: "message 5"},
		&ImageProcessingTask{"MNO"},
		&ImageProcessingTask{"PQR"},
		&ImageProcessingTask{"STU"},
		&EmailTask{EmailId: "VWX", Subject: "hello vwx", Message: "message 6"},
	}

	//create a worker pool with 5 concurrent workers
	wp := NewWorkerPool{
		MultiTasks:  multiTask,
		Concurrency: 3,
	}

	wp.Run()
	fmt.Println("All tasks completed.")
}
