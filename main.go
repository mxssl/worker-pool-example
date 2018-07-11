package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, ch, wg)
	}

	go func() {
		for taskID := 0; taskID < 100; taskID++ {
			ch <- taskID
		}
		close(ch)
	}()

	wg.Wait()
}

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range ch {
		fmt.Printf("Worker ID: %v\nTask: %v\n\n", id, task)
	}
}
