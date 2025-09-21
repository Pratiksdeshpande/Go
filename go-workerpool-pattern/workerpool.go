package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Concurrent worker pool pattern for processing tasks.
The worker pool manages a fixed number of goroutines that process tasks
from a shared channel, providing controlled concurrency and resource management.
Note: This implementation supports only one Task type at a time.
*/

// Task represents a unit of work to be processed by the worker pool
type Task struct {
	Id int
}

// Process way to process the tasks
func (t *Task) Process() {

	// Simulate task processing time
	fmt.Println("Processing task with ID:", t.Id)
	time.Sleep(5 * time.Second)
}

// WorkerPool definition
type WorkerPool struct {
	Tasks       []Task         // Tasks to be processed
	Concurrency int            // Number of concurrent workers
	TaskChan    chan Task      // Channel for distributing tasks to workers
	wg          sync.WaitGroup // WaitGroup to synchronize worker completion
}

// worker continuously processes tasks from the task channel until channel is closed
func (wp *WorkerPool) worker() {
	for task := range wp.TaskChan {
		task.Process()
		wp.wg.Done()
	}
}

// Run executes all tasks using the configured number of workers
func (wp *WorkerPool) Run() {
	// initialize the task channel
	wp.TaskChan = make(chan Task, len(wp.Tasks))

	// start workers
	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	// send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.TaskChan <- task
	}
	// close the task channel after all tasks are sent to the channel to avoid deadlock
	close(wp.TaskChan)

	// wait for all tasks to complete
	wp.wg.Wait()
}
