package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hell world")

	queue := make(chan int, 1000)
	sum := 0

	start := time.Now()
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			ranTime := rand.Intn(10)
			fmt.Printf("#%d Sleeping %d seconds...\n", i, ranTime)
			time.Sleep(time.Second + time.Duration(ranTime)*time.Second)
			fmt.Println("#%d hi", i)
			queue <- ranTime
			wg.Done()
		}(i)
	}

	go func() {
		for randomInt := range queue {
			sum += randomInt
		}
	}()

	wg.Wait()
	close(queue)

	fmt.Println("sum: %d", sum)
	fmt.Println("all time: ", time.Since(start))
}
