// ЗАДАЧА 14: Run конкурентных функций
// Напишите функцию Run, которая запускает конкурентное выполнение функций fs и дожидается их окончания. Если одна или
// несколько функций из fs завершились с ошибкой, Run возвращает любую из них.

package main

import (
	"errors"
	"sync"
)

type fn func() error

func main() {
	expErr := errors.New("error")

	funcs := []fn{
		func() error { return nil },
		func() error { return nil },
		func() error { return expErr },
		func() error { return nil },
	}

	if err := Run(funcs...); !errors.Is(err, expErr) {
		panic("wrong code")
	}
}

func Run(fs ...fn) error {
	var wg sync.WaitGroup
	var once sync.Once
	var err error

	wg.Add(len(fs))
	for _, f := range fs {
		go func(f fn) {
			defer wg.Done()

			if e := f(); e != nil {
				once.Do(func() {
					err = e
				})
			}
		}(f)
	}

	wg.Wait()
	return err
}
