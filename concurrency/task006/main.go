// ЗАДАЧА 6: MergeSorted
// Напишите функцию mergeSorted, которая принимает на вход два "отсортированных" канала и возвращает результирующий
// отсортированный канал.

package main

import (
	"fmt"
	"math"
)

func mergeSorted(cs ...<-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		vs := make(map[int]int)

		for i, c := range cs {
			if v, ok := <-c; ok {
				vs[i] = v
			}
		}

		for len(vs) > 0 {
			minV := math.MaxInt
			key := -1
			for k, v := range vs {
				if v <= minV {
					key = k
					minV = v
				}
			}

			ch <- minV

			if v, ok := <-cs[key]; ok {
				vs[key] = v
			} else {
				delete(vs, key)
			}
		}
	}()

	return ch
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 9223372036854775807
	close(c)
}

func main() {
	for range 10 {
		a, b := make(chan int), make(chan int)
		go fillChanA(a)
		go fillChanB(b)

		c := mergeSorted(a, b)

		for val := range c {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// Вывод: -1 1 2 4 4 5

//func mergeSorted(cs ...<-chan int) <-chan int {
//	var mu sync.Mutex
//	ch := make(chan int)
//	vs := make(map[int]int, len(cs))
//	var wg sync.WaitGroup
//	sigChan := make(chan struct{})
//
//	wg.Add(len(cs))
//	for i, c := range cs {
//		go func() {
//			defer wg.Done()
//			for v := range c {
//				mu.Lock()
//				vs[i] = v
//				mu.Unlock()
//				sigChan <- struct{}{}
//			}
//		}()
//	}
//
//	go func() {
//		for range sigChan {
//			if len(vs) < len(cs) {
//				continue
//			}
//
//			var minV *int
//			mu.Lock()
//			for k, v := range vs {
//				if minV == nil || *minV > v {
//					minV = &v
//					delete(vs, k)
//				}
//			}
//			mu.Unlock()
//
//			ch <- *minV
//			minV = nil
//		}
//
//		mu.Lock()
//		for len(vs) > 0 {
//			var minV *int
//			for k, v := range vs {
//				if minV == nil || *minV > v {
//					minV = &v
//					delete(vs, k)
//				}
//			}
//
//			ch <- *minV
//			minV = nil
//		}
//		mu.Unlock()
//
//		close(ch)
//	}()
//
//	go func() {
//		wg.Wait()
//		close(sigChan)
//	}()
//
//	return ch
//}
