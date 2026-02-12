// Package slices some tasks
package slices

import (
	"fmt"
	"reflect"
)

const Count = 50

func GetTasks() map[int]func() {
	return map[int]func(){
		1: a, 2: b, 3: c, 4: d, 5: e, 6: f, 7: g, 8: h, 9: i, 10: j,
		11: k, 12: l, 13: m, 14: n, 15: o, 16: p, 17: q, 18: r, 19: s, 20: t,
		21: u, 22: v, 23: w, 24: x, 25: y, 26: z, 27: aa, 28: ab, 29: ac, 30: ad,
		31: ae, 32: af, 33: ag, 34: ah, 35: ai, 36: aj, 37: ak, 38: al, 39: am, 40: an,
		41: ao, 42: ap, 43: aq, 44: ar, 45: as, 46: at, 47: au, 48: av, 49: aw, 50: ax,
	}
}

// ЗАДАЧА 1 Что выведет?
func a() {
	x := []int{}
	fmt.Println("x = ", x, "len: ", len(x), "cap: ", cap(x)) // [] 0 0
	x = append(x, 0)
	fmt.Println("x = ", x, "len: ", len(x), "cap: ", cap(x)) // [0] 1 1
	x = append(x, 1)
	fmt.Println("x = ", x, "len: ", len(x), "cap: ", cap(x)) // [01] 2 2
	x = append(x, 2)
	fmt.Println("x = ", x, "len: ", len(x), "cap: ", cap(x))             // [012] 3 4
	y := append(x, 3)                                                    // baseArr [0123]; len(x) = 3; cap(x) = 4
	fmt.Println("y = ", y, "len: ", len(y), "cap: ", cap(y))             // [0123] 4 4
	z := append(x, 4)                                                    // baseArr [0124]; len(x) = 3; cap(x) = 4
	fmt.Println("z = ", z, "len: ", len(z), "cap: ", cap(z))             // [0124] 4 4
	fmt.Println(x, len(x), cap(x), y, len(y), cap(y), z, len(y), cap(z)) //[012] 3 4; [0124] 4 4; [0124] 4 4
}

// ЗАДАЧА 2: Что выведет?
func b() {
	var x []int
	y := []int{}
	fmt.Println(x == nil, y == nil)
	fmt.Println(len(x), len(y))
	fmt.Println(cap(x), cap(y))
}

// ЗАДАЧА 3: Что выведет?
func c() {
	x := make([]int, 0, 5)
	fmt.Println(len(x), cap(x))
	x = append(x, 1, 2, 3)
	fmt.Println(len(x), cap(x))
	x = append(x, 4, 5, 6)
	fmt.Println(len(x), cap(x))
}

// ЗАДАЧА 4: Что выведет?
func d() {
	x := []int{1, 2, 3, 4, 5}
	y := x[1:3]
	y[0] = 100
	fmt.Println(x)
	fmt.Println(y)
}

// ЗАДАЧА 5: Что выведет?
func e() {
	x := []int{1, 2, 3, 4, 5}
	y := x[1:3]
	y = append(y, 100)
	fmt.Println(x)
	fmt.Println(y)
}

// ЗАДАЧА 6: Что выведет?
func f() {
	x := []int{1, 2, 3}
	y := x
	y[0] = 100
	fmt.Println(x)
}

// ЗАДАЧА 7: Что выведет?
func g() {
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 2)
	copy(y, x)
	fmt.Println(y)
}

// ЗАДАЧА 8: Что выведет?
func h() {
	x := []int{1, 2, 3}
	copy(x, x[1:])
	fmt.Println(x)
}

// ЗАДАЧА 9: Что выведет?
func i() {
	x := []int{1, 2, 3}
	for i, v := range x {
		x[i] = v * 2
	}
	fmt.Println(x)
}

// ЗАДАЧА 10: Что выведет?
func j() {
	type Person struct{ name string }
	x := []Person{{"Alice"}, {"Bob"}}
	for _, v := range x {
		v.name = "Changed"
	}
	fmt.Println(x)
}

// ЗАДАЧА 11: Что выведет?
func k() {
	x := []int{1, 2, 3, 4, 5}
	y := x[:3]
	z := append(y, 100)
	fmt.Println(x)
	fmt.Println(z)
}

// ЗАДАЧА 12: Что выведет?
func l() {
	x := make([]int, 3, 5)
	x[0] = 1
	x = append(x, 2)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
}

// ЗАДАЧА 13: Что выведет?
func m() {
	arr := [5]int{1, 2, 3, 4, 5}
	x := arr[1:3]
	x[0] = 100
	fmt.Println(arr)
	fmt.Println(x)
}

// ЗАДАЧА 14: Что выведет?
func n() {
	var x []int
	x = append(x, 1)
	fmt.Println(x)
}

// ЗАДАЧА 15: Что выведет?
func o() {
	x := []int{1, 2}
	y := append(x, 3)
	z := append(x, 4)
	w := append(x, 5)
	fmt.Println(x, y, z, w)
}

// ЗАДАЧА 16: Что выведет?
func p() {
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	fmt.Println("Нужно написать функцию для сравнения", len(x), len(y))
	fmt.Println(reflect.DeepEqual(x, y))
}

// ЗАДАЧА 17: Что выведет?
func q() {
	x := make([]int, 3, 10)
	y := x[:2:2]
	fmt.Println(len(y), cap(y))
}

// ЗАДАЧА 18: Что выведет?
func r() {
	x := []int{1, 2}
	y := []int{3, 4, 5}
	x = append(x, y...)
	fmt.Println(x, "|", len(x), cap(x))
}

// ЗАДАЧА 19: Что выведет?
func s() {
	x := []int{1, 2, 3, 4, 5}
	i := 2
	x = append(x[:i], x[i+1:]...)
	fmt.Println(x)
}

// ЗАДАЧА 20: Что выведет?
func t() {
	x := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-1-i] = x[len(x)-1-i], x[i]
	}
	fmt.Println(x)
}

// ЗАДАЧА 21: Что выведет?
func u() {
	x := make([]int, 0, 4)
	fmt.Println(cap(x))
	x = append(x, 1, 2, 3, 4)
	fmt.Println(cap(x))
	x = append(x, 5)
	fmt.Println(cap(x))
}

// ЗАДАЧА 22: Что выведет?
func v() {
	x := []int{1, 2, 3, 4, 5}
	y := x[2:]
	z := y[1:]
	fmt.Println(z)
}

// ЗАДАЧА 23: Что выведет?
func w() {
	x := []int{1, 2, 3, 4, 5}
	y := x[1:3]
	z := x[2:4]
	y[1] = 100
	fmt.Println(y)
	fmt.Println(z)
}

// ЗАДАЧА 24: Что выведет?
func x() {
	var result []int
	for i := 0; i < 5; i++ {
		result = append(result, i*2)
	}
	fmt.Println(result)
	fmt.Println(len(result), cap(result))
}

// ЗАДАЧА 25: Что выведет?
func y() {
	x := []int{1, 2, 4, 5}
	i := 2
	x = append(x[:i], append([]int{3}, x[i:]...)...)
	fmt.Println(x)
}

// ЗАДАЧА 26: Что выведет?
func z() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := x[:0]
	fmt.Println(len(result), cap(result))
	for _, v := range x {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	fmt.Println(result)
	fmt.Println(x)
}

// ЗАДАЧА 27: Что выведет?
func aa() {
	x := [][]int{{1, 2}, {3, 4}}
	x[0][0] = 100
	fmt.Println(x)
}

// ЗАДАЧА 28: Что выведет?
func ab() {
	a, b := 1, 2
	x := []*int{&a, &b}
	*x[0] = 100
	fmt.Println(a, b)
	fmt.Println(*x[0], *x[1])
}

// ЗАДАЧА 29: Что выведет?
func ac() {
	var x []int
	for i := 0; i < 3; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
}

// ЗАДАЧА 30: Что выведет?
func ad() {
	x := []int{1, 2, 3, 4, 5}
	y := x[1:1]
	fmt.Println(y)
	fmt.Println(len(y))
}

// ЗАДАЧА 31: Что выведет?
func ae() {
	x := []int{1, 2, 3}
	modifySlice(x)
	fmt.Println(x)
}
func modifySlice(s []int) {
	s[0] = 100
}

// ЗАДАЧА 32: Что выведет?
func af() {
	x := []int{1, 2, 3}
	appendToSlice(x)
	fmt.Println(x, "|", len(x), cap(x))
}
func appendToSlice(s []int) {
	s = append(s, 4)
}

// ЗАДАЧА 33: Что выведет?
func ag() {
	s := "hello"
	x := []byte(s)
	x[0] = 'H'
	fmt.Println(s)
	fmt.Println(string(x))
}

// ЗАДАЧА 34: Что выведет?
func ah() {
	s := "привет"
	x := []rune(s)
	fmt.Println(len(s))
	fmt.Println(len(x))
}

// ЗАДАЧА 35: Что выведет?
func ai() {
	x := []int{}

	fmt.Println("len: ", len(x), "cap: ", cap(x))
	for i := 0; i < 5; i++ {
		x = append(x, i)
		fmt.Println("len:", len(x), "cap:", cap(x))
	}
}

// ЗАДАЧА 36: Что выведет?
func aj() {
	x := []int{2, 3, 4}
	x = append([]int{1}, x...)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
}

// ЗАДАЧА 37: Что выведет?
func ak() {
	x := []int{1, 2, 3, 4, 5}
	x = x[:0]
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
}

// ЗАДАЧА 38: Что выведет?
func al() {
	x := []int{1, 2, 3, 4, 5}
	copy(x[2:], x[:3])
	fmt.Println(x)
}

// ЗАДАЧА 39: Что выведет?
func am() {
	x := make([]int, 3, 10)
	x[0], x[1], x[2] = 1, 2, 3
	y := x[1:2]
	fmt.Println(len(y), cap(y))
}

// ЗАДАЧА 40: Что выведет?
func an() {
	x := make([]int, 2, 4)
	x[0], x[1] = 1, 2
	y := x[0:2:2]
	y = append(y, 3)
	x = append(x, 4)
	fmt.Println("x: ", x, "|", "len: ", len(x), "|", "cap: ", cap(x))
	fmt.Println("y: ", y, "|", "len: ", len(y), "|", "cap: ", cap(y))
}

// ЗАДАЧА 41: Что выведет?
func ao() {
	x := []int{1, 2, 3}
	y := x
	z := x
	y = append(y, 4)
	z = append(z, 5)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// ЗАДАЧА 42: Что выведет?
func ap() {
	var x []int
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("Done")
}

// ЗАДАЧА 43: Что выведет?
func aq() {
	var x []int
	y := []int{}
	fmt.Println(x == nil)
	fmt.Println(y == nil)
	fmt.Println(len(x) == 0)
	fmt.Println(len(y) == 0)
}

// ЗАДАЧА 44: Что выведет?
func ar() {
	x := []int{1, 2}
	x = append(x, append([]int{3}, 4)...)
	fmt.Println(x)
}

// ЗАДАЧА 45: Что выведет?
func as() {
	x := []int{1, 2, 3}
	fmt.Println("len: ", len(x), "cap: ", cap(x))
	for i := range x {
		x = append(x, i)
		if i > 5 {
			break
		}
	}
	fmt.Println(x, len(x), cap(x))
}

// ЗАДАЧА 46: Что выведет?
func at() {
	x := make([]int, 3, 10)
	x[0], x[1], x[2] = 1, 2, 3
	y := x[:6]
	fmt.Println(y)
}

// ЗАДАЧА 47: Что выведет?
func au() {
	x := []int{1, 2, 3}
	p := &x[1]
	x = append(x, 4, 5, 6, 7, 8)
	*p = 100
	fmt.Println(x)
	fmt.Println(p)
}

// ЗАДАЧА 48: Что выведет?
func av() {
	x := []int{1, 2, 3, 4, 5}
	y := x[2:4]
	z := x[2:3]
	z[0] = 100
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// ЗАДАЧА 49: Что выведет?
func aw() {
	x := []int{1, 1, 2, 2, 3, 3}
	j := 0
	for i := 1; i < len(x); i++ {
		if x[i] != x[j] {
			j++
			x[j] = x[i]
		}
	}
	result := x[:j+1]
	fmt.Println(result)
}

// ЗАДАЧА 50: Что выведет?
func ax() {
	x := make([]int, 5, 10)
	for i := range x {
		x[i] = i + 1
	}
	fmt.Println(x, len(x))
	y := x[1:3:4]
	fmt.Println(y)
	fmt.Println(len(y), cap(y))
	y = append(y, 100)
	fmt.Println(x)
	fmt.Println(y)
}
