// ЗАДАЧА 13: ConcurrentSortHead
// Напишите функцию ConcurrentSortHead, которая из files ридеров, которые содержат упорядоченные по возрастанию
// строки, вернет m первых строк. Чтение из ридеров files должно быть конкурентным.

package main

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	f1 := `aaa
ddd
`
	f2 := `bbb
eee
`
	f3 := `ccc
fff
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
	fLen := len(files)
	strChans := make([]chan string, 0, fLen)
	strs := make([]string, 0, fLen)

	for i := range fLen {
		go func(i int) {
			scanner := bufio.NewScanner(files[i])
			ch := make(chan string)
			strChans[i] = ch

			for scanner.Scan() {
				strChans[i] <- scanner.Text()
			}
		}(i)
	}

	for i := range fLen {
		s := <-strChans[i]
		strs = append(strs, s)

	}

	for range m {
	}

	return nil, nil
}
