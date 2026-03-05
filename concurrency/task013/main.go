// ЗАДАЧА 13: ConcurrentSortHead
// Напишите функцию ConcurrentSortHead, которая из files ридеров, которые содержат упорядоченные по возрастанию
// строки, вернет m первых строк. Чтение из ридеров files должно быть конкурентным.

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	f1 := `bbb
eee
`
	f2 := `aaa
fff
`
	f3 := `ccc
ddd
`

	files := []io.Reader{
		strings.NewReader(f1),
		strings.NewReader(f2),
		strings.NewReader(f3),
	}

	rows, err := ConcurrentSortHead(4, files...)
	fmt.Println(rows)
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(rows, []string{"aaa", "bbb", "ccc", "ddd"}) {
		panic("wrong code")
	}
}

func ConcurrentSortHead(m int, files ...io.Reader) ([]string, error) {
	if m <= 0 || len(files) == 0 {
		return []string{}, nil
	}

	n := len(files)
	channels := make([]chan string, n)
	errCh := make(chan error, n)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// создаём буферизированные каналы
	for i := 0; i < n; i++ {
		channels[i] = make(chan string, 1)
	}

	// запускаем читателей
	for i := 0; i < n; i++ {
		go func(i int) {
			defer close(channels[i])

			scanner := bufio.NewScanner(files[i])
			for scanner.Scan() {
				select {
				case channels[i] <- scanner.Text():
				case <-ctx.Done():
					return
				}
			}

			if err := scanner.Err(); err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}(i)
	}

	result := make([]string, 0, m)
	curr := make([]string, n)
	active := make([]bool, n)

	// читаем первую строку из каждого канала
	for i := 0; i < n; i++ {
		select {
		case err := <-errCh:
			return nil, err
		case s, ok := <-channels[i]:
			if ok {
				curr[i] = s
				active[i] = true
			}
		}
	}

	for len(result) < m {
		// проверяем ошибки
		select {
		case err := <-errCh:
			cancel()
			return nil, err
		default:
		}

		minIdx := -1
		var minVal string

		for i := 0; i < n; i++ {
			if !active[i] {
				continue
			}
			if minIdx == -1 || curr[i] < minVal {
				minIdx = i
				minVal = curr[i]
			}
		}

		if minIdx == -1 {
			break // больше строк нет
		}

		result = append(result, minVal)

		// читаем следующую строку только из выбранного канала
		select {
		case err := <-errCh:
			cancel()
			return nil, err
		case s, ok := <-channels[minIdx]:
			if ok {
				curr[minIdx] = s
			} else {
				active[minIdx] = false
			}
		case <-ctx.Done():
			return result, nil
		}
	}

	// досрочно отменяем оставшиеся горутины
	cancel()
	return result, nil
}
