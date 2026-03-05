// ЗАДАЧА 16: Inc цепочка каналов

// Напишите функцию inc, которая принимает на вход канал, читает из него значения и записывает эти значения, увеличенные
// на единицу, в возвращаемый канал.
// Дополните функцию main созданием цепочки каналов, используя inc, так, чтобы программа завершалась без паники.

package main

func main() {
	last := make(<-chan int)
	n := 10

	start := make(chan int, 1)
	start <- 0
	close(start)

	last = start

	for range n {
		out := inc(last)
		last = out
	}

	result := <-last
	if n != result {
		panic("wrong code")
	}
}

func inc(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			out <- v + 1
		}
	}()

	return out
}
