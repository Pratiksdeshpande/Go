package main

import (
	"fmt"
	"sync"
	"time"
)

/*
concurrent worker pool pattern for processing multiple type of tasks at a time.
*/

// MultiTask definition
type MultiTask interface {
	Process()
}

// EmailTask definition
type EmailTask struct {
	EmailId string
	Subject string
	Message string
}

// Process way to process the email tasks
func (e *EmailTask) Process() {
	fmt.Println("Sending email to:", e.EmailId)
	time.Sleep(1 * time.Second)
}

// ImageProcessingTask definition
type ImageProcessingTask struct {
	ImageURL string
}

// Process way to process the image processing tasks
func (e *ImageProcessingTask) Process() {
	fmt.Println("Processing image from URL:", e.ImageURL)
	time.Sleep(4 * time.Second)
}

// NewWorkerPool definition
type NewWorkerPool struct {
	MultiTasks    []MultiTask    // MultiTask to be processed
	Concurrency   int            // Number of concurrent workers
	MultiTaskChan chan MultiTask // Channel for distributing multiple tasks to workers
	wg            sync.WaitGroup // WaitGroup to synchronize worker completion
}

// worker continuously processes tasks from the task channel until channel is closed
func (wp *NewWorkerPool) worker() {
	for task := range wp.MultiTaskChan {
		task.Process()
		wp.wg.Done()
	}
}

// Run executes all tasks using the configured number of workers
func (wp *NewWorkerPool) Run() {
	// initialize the task channel
	wp.MultiTaskChan = make(chan MultiTask, len(wp.MultiTasks))

	// start workers
	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	// send tasks to the tasks channel
	wp.wg.Add(len(wp.MultiTasks))
	for _, task := range wp.MultiTasks {
		wp.MultiTaskChan <- task
	}
	// close the task channel after all tasks are sent to the channel to avoid deadlock
	close(wp.MultiTaskChan)

	// wait for all tasks to complete
	wp.wg.Wait()
}
