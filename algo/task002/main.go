package main

import (
	"fmt"
	"slices"
)

//import "slices"

/*
// Необходимо проверить 2 строки, являются ли они а‍награммами.
// Если это так, т‍о вернуть true, иначе false.
// В строках могут быть т‍олько буквы из латиницы и кириллицы

// Пример1
// in
s = "anagram"
t = "nagaram"
// out
true

// Пример2
// in
s = "кит"
t = "ток"
// out
false
*/

func main() {
	fmt.Println(anagram2("кит", "тик"))
}

func anagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	chars := make(map[rune]int)

	for _, c := range sr {
		chars[c]++
	}

	for _, c := range tr {
		chars[c]--
		if chars[c] < 0 {
			return false
		}
	}

	return true
}

func anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	for _, sRune := range sr {
		i := slices.Index(tr, sRune)
		notContain := i < 0
		if notContain {
			return false
		}

		tr = append(tr[:i], tr[i+1:]...)
	}

	return true
}
