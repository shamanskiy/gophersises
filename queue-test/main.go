package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Id        int
	SleepTime float32
}

var wg sync.WaitGroup

func worker(sleepJobs <-chan Job, id int) {
	defer wg.Done()
	jobsDone := 0
	for job := range sleepJobs {
		fmt.Printf("Worker %d: job %d, going to sleep for %fs...\n", id, job.Id, job.SleepTime)
		time.Sleep(time.Duration(job.SleepTime) * time.Second)
		jobsDone++
	}
	fmt.Printf("Worker %d: finished %d jobs\n", id, jobsDone)
}

func main() {
	sleepJobs := make(chan Job, 10)

	for id := 1; id < 3; id++ {
		wg.Add(1)
		go worker(sleepJobs, id)
	}

	fmt.Println("Main: enqueuing jobs...")
	for i := 1; i < 31; i++ {
		sleepTime := rand.Float32() * 3.
		fmt.Printf("Main: waiting to enqueue job %d...\n", i)
		sleepJobs <- Job{i, sleepTime}
		fmt.Printf("Main: enqueued job %d\n", i)
	}

	fmt.Println("Main: closing queue...")
	close(sleepJobs)

	fmt.Println("Main: waiting for workers...")
	wg.Wait()
}
