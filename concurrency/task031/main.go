package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// напишите функции writer, double, reader.
// writer - пишет значения от 1 до 100 в возвращаемый канал
// double - возвращает канал и читает из принимаемого канала значения и умножает их
// reader - выводит получаемые значения из канала
// По сути это реализация Pipeline

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	reader(double(writer(ctx)))
}

func writer(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		count := 0

		for {
			select {
			case <-ctx.Done():
				time.Sleep(3 * time.Second)
				return
			default:
			}

			count++
			select {
			case <-ctx.Done():
				time.Sleep(3 * time.Second)
				return
			case out <- count:
			}
		}
	}()

	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			time.Sleep(500 * time.Millisecond)
			out <- v * v
		}
	}()

	return out
}

func reader(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
