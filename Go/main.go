package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan string, results *[]string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started processing %s\n", id, task)
		time.Sleep(500 * time.Millisecond) // simulate work
		result := fmt.Sprintf("Processed: %s", task)

		mu.Lock()
		*results = append(*results, result)
		mu.Unlock()

		fmt.Printf("Worker %d finished processing %s\n", id, task)
	}
}

func main() {
	tasks := make(chan string, 10)
	var results []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Add tasks
	for i := 1; i <= 10; i++ {
		tasks <- fmt.Sprintf("Task %d", i)
	}
	close(tasks)

	// Start workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &results, &mu, &wg)
	}

	wg.Wait()

	fmt.Println("Final Results:")
	for _, res := range results {
		fmt.Println(res)
	}
}
