package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	err := predictableTimeWork(randomTimeWork)

	fmt.Println("Duration: ", time.Since(start), "\nError: ", err)
}

// есть функция, которая работает неопределённо долго (до 100 секунд)
func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

// необходимо написать обёртку над этой функцией, которая будет прерывать выполнение,
// если функция будет работать дольше 3х секунд и возвращать ошибку
func predictableTimeWork(fn func()) error {
	return nil
}
