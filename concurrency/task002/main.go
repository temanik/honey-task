// ЗАДАЧА 2: Параллельное скачивание с обработкой ошибок
// Напишите функцию download.
//
// Функция download:
// - на вход получает адреса для скачивания - urls
// - конкурентно скачивает информацию из каждого url (для скачивания используйте функцию fakeDownload)
// - если вызовы fakeDownload возвращают ошибки, то нужно вернуть их все (см. errors.Join)

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// timeoutLimit - вероятность, с которой не будет возвращаться ошибка от fakeDownload():
// timeoutLimit = 100 - ошибок не будет;
// timeoutLimit = 0 - всегда будет возвращаться ошибка.
// Можете "поиграть" с этим параметром, для проверки случаев с возвращением ошибки.
const timeoutLimit = 90

type Result struct {
	msg string
	err error
}

// fakeDownload - имитирует разное время скачивания для разных адресов
func fakeDownload(url string) Result {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	if r > timeoutLimit {
		return Result{
			err: fmt.Errorf("failed to download data from %s: timeout", url),
		}
	}
	return Result{
		msg: fmt.Sprintf("downloaded data from %s\n", url),
	}
}

// download - параллельно скачивает данные из urls
func download(urls []string) ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]string, len(urls))
	var err error

	wg.Add(len(urls))
	for i, url := range urls {
		// Передаём i и url как аргументы горутины, чтобы избежать захвата переменных цикла
		go func(i int, url string) {
			defer wg.Done()

			res := fakeDownload(url)

			// Мьютекс защищает как запись в results, так и обновление err
			mu.Lock()
			defer mu.Unlock()

			if res.err != nil {
				// Добавляем ошибку в общий список
				err = errors.Join(err, res.err)
			} else {
				// Сохраняем результат по нужному индексу
				results[i] = res.msg
			}
		}(i, url)
	}

	wg.Wait()
	return results, err
}

func main() {
	msgs, err := download([]string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(msgs)
}
