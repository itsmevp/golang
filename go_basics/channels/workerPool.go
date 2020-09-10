package main

import (
	"fmt"
	"sync"
)

type job struct {
	jobID     int
	randomNum int
}

type result struct {
	job    job
	result int
}

var (
	jobChannel    = make(chan job, 10)
	resultChannel = make(chan result, 10)
)

func createJobs(numOfJobs int) {
	for i := 0; i < numOfJobs; i++ {
		jobChannel <- job{jobID: i, randomNum: i}
	}
	close(jobChannel)
}

func worker(w *sync.WaitGroup) {
	for j := range jobChannel {
		resultChannel <- result{job: j, result: (j.randomNum * j.randomNum)}
	}
	w.Done()
}

func createWorkers(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(resultChannel)
}

func printResults(c chan bool) {
	for r := range resultChannel {
		fmt.Println(r.job.jobID, "---->", r.result)
	}
	c <- true
}

func main() {
	go createJobs(10)
	createWorkers(2)
	c := make(chan bool)
	go printResults(c)
	<-c
}
