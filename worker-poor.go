package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w < 4; w++ {
		go worker2(w, jobs, results)
	}

	for i := 1; i < 10; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i < 10; i++ {
		<-results
	}
}

func worker2(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing jon", job)
		time.Sleep(time.Second)
		result <- job * 2
	}
}
