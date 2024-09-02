package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	jobChan := make(chan int64, 100)
	resultChan := make(chan int64, 100)
	defer close(jobChan)
	defer close(resultChan)
	go f1(jobChan)
	wg.Add(24)
	for range 24 {
		go f2(jobChan, resultChan)
	}
	for sum := range resultChan {
		fmt.Println(sum)
	}
	wg.Wait()
}

func f1(jobChan chan<- int64) {
	for {
		jobChan <- rand.Int63()
		// 1 Second = 1000 * Millisecond
		time.Sleep(1000 * time.Millisecond)
	}
}

func f2(jobChan <-chan int64, resultChan chan<- int64) {
	defer wg.Done()
	for num := range jobChan {
		var sum int64
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		resultChan <- sum
	}
}
