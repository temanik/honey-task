package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()

	start := time.Now()
	processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(start))
}

// операция должна выполняться не более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int) {

}

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}
