package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// worker processes tasks from a channel and stores results in a shared slice
// It uses a mutex to safely append results and a WaitGroup to track completion
func worker(id int, tasks <-chan string, results *[]string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done() // Make sure to mark this worker as done when it exits

	for task := range tasks {
		fmt.Printf("Worker %d started processing %s\n", id, task)
		time.Sleep(500 * time.Millisecond) // Simulate some actual work being done
		result := fmt.Sprintf("Processed: %s", task)

		// Safely append the result to our shared slice
		mu.Lock()
		*results = append(*results, result)
		mu.Unlock()

		fmt.Printf("Worker %d finished processing %s\n", id, task)
	}
}

func main() {
	// Create a buffered channel to hold our tasks
	tasks := make(chan string, 10)

	// Shared resources that need synchronization
	var results []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Queue up 10 tasks to be processed
	for i := 1; i <= 10; i++ {
		tasks <- fmt.Sprintf("Task %d", i)
	}
	close(tasks) // Signal that no more tasks will be added

	// Launch 3 worker goroutines to process the tasks
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, tasks, &results, &mu, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Write results to file
	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write each result to the file
	for _, res := range results {
		_, err := file.WriteString(res + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}

	fmt.Println("All tasks processed and results saved to results.txt")
}
