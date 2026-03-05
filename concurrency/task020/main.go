// ЗАДАЧА 20: Or функция
// Напишите функцию or, которая возвращает канал, который закрывается либо как только какой-либо из каналов
// channels стал доступен для чтения, либо как только какой-либо из каналов channels закрыли
// В функции main написан код, который при корректной работе функции or выведет число очень близкое к секунде.

package main

import (
	"fmt"
	"time"
)

func orGPT(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	out := make(chan interface{})

	go func() {
		defer close(out)

		select {
		case <-channels[0]:
		case <-channels[1]:
		case <-or(channels[2:]...):
		}
	}()

	return out
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}

	out := make(chan interface{})
	done := make(chan struct{}, 1)
	done <- struct{}{}

	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			select {
			case <-ch:
			case <-out:
				return
			}

			select {
			case <-done:
				close(out)
			default:
			}
		}(ch)
	}

	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
