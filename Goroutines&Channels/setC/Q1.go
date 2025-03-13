// Go Program to Implement Checkpoint Synchronization Problem
package main

import (
	"fmt"
	"sync"
	"time"
)

// Number of workers
const numWorkers = 5

// Create a barrier using a sync.WaitGroup
var wg sync.WaitGroup

func worker(id int, barrier chan bool) {
	defer wg.Done()

	// Simulate work time
	fmt.Printf("Worker %d is working...\n", id)
	time.Sleep(time.Duration(id) * time.Second)

	fmt.Printf("Worker %d finished work and waiting at the checkpoint\n", id)

	// Notify that this worker has reached the barrier
	barrier <- true

	// Wait for all workers at the checkpoint
	fmt.Printf("Worker %d is assembling with others at checkpoint\n", id)

	// Simulate assembling time
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d has completed the task\n", id)
}

func main() {
	barrier := make(chan bool, numWorkers)

	wg.Add(numWorkers)

	// Start all workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, barrier)
	}

	// Wait until all workers reach the checkpoint
	for i := 0; i < numWorkers; i++ {
		<-barrier
	}

	fmt.Println("All workers reached the checkpoint. Assembling details...")

	// Wait for all workers to finish assembling
	wg.Wait()

	fmt.Println("All tasks are completed!")
}
