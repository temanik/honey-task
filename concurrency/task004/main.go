// ЗАДАЧА 4: RepeatFn и Take
// Напишите функции repeatFn и take.
// Функция repeatFn бесконечно вызывает функцию fn и пишет результат ее работы в возвращаемый канал.
// Прекращает работу раньше, если контекст отменен.
// Функция take читает не более чем num из канала in, пока in открыт, и пишет значение в возвращаемый канал.
// Прекращает работу раньше, если контекст отменен.

package main

import (
	"context"
	"fmt"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fn():
			}
		}
	}()

	return ch
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		for range num {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				ch <- v
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rnd := func() interface{} { return rand.Int() }

	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, rnd), 3) {
		res = append(res, num)
	}

	fmt.Println(res)

	if len(res) != 3 {
		panic("wrong code")
	}
}
