package pointers

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Counter struct {
	count int
}

type Point struct {
	x, y int
}

type Node struct {
	value int
	next  *Node
}

type Box struct {
	value int
}

type Inner struct {
	value int
}

type Outer struct {
	Inner
	name string
}

type Data struct {
	values []int
}

func (c Counter) increment() {
	c.count++
}

func (c *Counter) incrementPtr() {
	c.count++
}

func (c *Counter) doublePtr() {
	c.count *= 2
}

func (b *Box) setValue(v int) {
	b.value = v
}

func GetTasks() map[int]func() {
	return map[int]func(){
		1: task1, 2: task2, 3: task3, 4: task4, 5: task5, 6: task6, 7: task7, 8: task8, 9: task9, 10: task10,
		11: task11, 12: task12, 13: task13, 14: task14, 15: task15, 16: task16, 17: task17, 18: task18, 19: task19, 20: task20,
		21: task21, 22: task22, 23: task23, 24: task24, 25: task25, 26: task26, 27: task27, 28: task28, 29: task29, 30: task30,
		31: task31, 32: task32, 33: task33, 34: task34, 35: task35, 36: task36, 37: task37, 38: task38, 39: task39, 40: task40,
		41: task41, 42: task42, 43: task43, 44: task44, 45: task45, 46: task46, 47: task47, 48: task48, 49: task49, 50: task50,
	}
}

func Count() int {
	return 50
}

// ЗАДАЧА 1: Что выведет? (основы указателей)
func task1() {
	x := 10
	p := &x
	fmt.Println(x, *p)
	*p = 20
	fmt.Println(x, *p)
}

// ЗАДАЧА 2: Что выведет? (nil pointer)
func task2() {
	var p *int
	fmt.Println(p)
	fmt.Println(p == nil)
}

// ЗАДАЧА 3: Что выведет? (new function)
func task3() {
	p := new(int)
	fmt.Println(*p)
	*p = 42
	fmt.Println(*p)
}

// ЗАДАЧА 4: Что выведет? (указатели на структуры)
func task4() {
	p := &Person{name: "Alice", age: 30}
	fmt.Println(p.name, p.age)
	p.age = 31
	fmt.Println(p.name, p.age)
}

// ЗАДАЧА 5: Что выведет? (передача по значению vs по указателю)
func task5() {
	x := 10
	changeValue(x)
	fmt.Println(x)
	changePointer(&x)
	fmt.Println(x)
}
func changeValue(val int) {
	val = 20
}
func changePointer(ptr *int) {
	*ptr = 20
}

// ЗАДАЧА 6: Что выведет? (указатели на одно значение)
func task6() {
	x := 10
	p1 := &x
	p2 := &x
	*p1 = 20
	fmt.Println(*p2)
}

// ЗАДАЧА 7: Что выведет? (адреса)
func task7() {
	x := 10
	p := &x
	fmt.Println(&x == p)
	fmt.Println(&x == &x)
}

// ЗАДАЧА 8: Что выведет? (указатель на указатель)
func task8() {
	x := 10
	p := &x
	pp := &p
	fmt.Println(x, *p, **pp)
	**pp = 20
	fmt.Println(x)
}

// ЗАДАЧА 9: Что выведет? (zero value структуры)
func task9() {
	p := new(Point)
	fmt.Println(*p)
}

// ЗАДАЧА 10: Что выведет? (сравнение указателей)
func task10() {
	x, y := 10, 10
	px := &x
	py := &y
	fmt.Println(px == py)
	fmt.Println(*px == *py)
}

// ЗАДАЧА 11: Что выведет? (указатель в цикле)
func task11() {
	arr := [3]int{1, 2, 3}
	var ptrs []*int
	for i := range arr {
		ptrs = append(ptrs, &arr[i])
	}
	for _, p := range ptrs {
		fmt.Print(*p, " ")
	}
	fmt.Println()
}

// ЗАДАЧА 12: Что выведет? (указатель на локальную переменную)
func task12() {
	p := getPointer()
	fmt.Println(*p)
}
func getPointer() *int {
	x := 42
	return &x
}

// ЗАДАЧА 13: Что выведет? (методы с pointer receiver)
func task13() {
	c := Counter{count: 0}
	c.increment()
	fmt.Println(c.count)
	c.incrementPtr()
	fmt.Println(c.count)
}

// ЗАДАЧА 14: Что выведет? (nil pointer dereference - panic!)
func task14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	var p *int
	fmt.Println("Before panic")
	fmt.Println(*p)
	fmt.Println("After panic")
}

// ЗАДАЧА 15: Что выведет? (указатель на массив vs слайс)
func task15() {
	arr := [3]int{1, 2, 3}
	p := &arr
	p[0] = 100
	fmt.Println(arr)
}

// ЗАДАЧА 16: Что выведет? (swap через указатели)
func task16() {
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println(x, y)
}
func swap(a, b *int) {
	*a, *b = *b, *a
}

// ЗАДАЧА 17: Что выведет? (указатель на элемент слайса)
func task17() {
	s := []int{1, 2, 3}
	p := &s[1]
	*p = 100
	fmt.Println(s)
}

// ЗАДАЧА 18: Что выведет? (указатели в структуре)
func task18() {
	n1 := Node{value: 1}
	n2 := Node{value: 2}
	n1.next = &n2
	fmt.Println(n1.value, n1.next.value)
}

// ЗАДАЧА 19: Что выведет? (модификация через функцию)
func task19() {
	p := Person{name: "Bob", age: 25}
	updatePerson19(&p)
	fmt.Println(p)
}
func updatePerson19(p *Person) {
	p.name = "Alice"
	p.age = 30
}

// ЗАДАЧА 20: Что выведет? (указатель на bool)
func task20() {
	b := true
	p := &b
	*p = false
	fmt.Println(b, *p)
}

// ЗАДАЧА 21: Что выведет? (указатель на string)
func task21() {
	s := "hello"
	p := &s
	*p = "world"
	fmt.Println(s, *p)
}

// ЗАДАЧА 22: Что выведет? (slice of pointers)
func task22() {
	a, b, c := 1, 2, 3
	ptrs := []*int{&a, &b, &c}
	for _, p := range ptrs {
		*p = *p * 2
	}
	fmt.Println(a, b, c)
}

// ЗАДАЧА 23: Что выведет? (указатель на interface{})
func task23() {
	var x interface{} = 10
	p := &x
	*p = "hello"
	fmt.Println(x)
}

// ЗАДАЧА 24: Что выведет? (копирование структуры с указателями)
func task24() {
	type Data24 struct {
		value *int
	}
	x := 10
	d1 := Data24{value: &x}
	d2 := d1
	*d2.value = 20
	fmt.Println(*d1.value)
}

// ЗАДАЧА 25: Что выведет? (указатель в map)
func task25() {
	m := make(map[string]*int)
	x := 10
	m["a"] = &x
	*m["a"] = 20
	fmt.Println(x)
}

// ЗАДАЧА 26: Что выведет? (указатель на map)
func task26() {
	m := map[string]int{"a": 1}
	p := &m
	(*p)["a"] = 2
	fmt.Println(m["a"])
}

// ЗАДАЧА 27: Что выведет? (указатель на функцию - не совсем указатель)
func task27() {
	f := func(x int) int { return x * 2 }
	result := f(5)
	fmt.Println(result)
}

// ЗАДАЧА 28: Что выведет? (pointer receiver автоматическая разыменовка)
func task28() {
	c := Counter{count: 5}
	c.doublePtr()
	fmt.Println(c.count)

	pc := &c
	pc.doublePtr()
	fmt.Println(c.count)
}

// ЗАДАЧА 29: Что выведет? (указатель на канал)
func task29() {
	ch := make(chan int, 1)
	p := &ch
	*p <- 42
	fmt.Println(<-*p)
}

// ЗАДАЧА 30: Что выведет? (new vs make)
func task30() {
	p1 := new([]int)
	fmt.Println(*p1 == nil)

	p2 := new(map[string]int)
	fmt.Println(*p2 == nil)
}

// ЗАДАЧА 31: Что выведет? (указатели в range)
func task31() {
	nums := []int{1, 2, 3}
	var ptrs []*int
	for _, v := range nums {
		ptrs = append(ptrs, &v)
	}
	for _, p := range ptrs {
		fmt.Print(*p, " ")
	}
	fmt.Println()
}

// ЗАДАЧА 32: Что выведет? (исправленная версия задачи 31)
func task32() {
	nums := []int{1, 2, 3}
	var ptrs []*int
	for i := range nums {
		ptrs = append(ptrs, &nums[i])
	}
	for _, p := range ptrs {
		fmt.Print(*p, " ")
	}
	fmt.Println()
}

// ЗАДАЧА 33: Что выведет? (указатель на slice - изменение length)
func task33() {
	s := []int{1, 2, 3}
	appendToSlicePtr(&s, 4)
	fmt.Println(s)
}
func appendToSlicePtr(s *[]int, val int) {
	*s = append(*s, val)
}

// ЗАДАЧА 34: Что выведет? (указатель внутри goroutine)
func task34() {
	x := 0
	done := make(chan bool)
	go func(p *int) {
		*p = 42
		done <- true
	}(&x)
	<-done
	fmt.Println(x)
}

// ЗАДАЧА 35: Что выведет? (множественные указатели на структуру)
func task35() {
	type Box struct{ value int }
	b := Box{value: 10}
	p1, p2 := &b, &b
	p1.value = 20
	fmt.Println(p2.value)
}

// ЗАДАЧА 36: Что выведет? (указатель на анонимную структуру)
func task36() {
	p := &Point{x: 10, y: 20}
	fmt.Println(p.x, p.y)
}

// ЗАДАЧА 37: Что выведет? (возврат указателя на литерал)
func task37() {
	p := getIntPtr(100)
	fmt.Println(*p)
}
func getIntPtr(x int) *int {
	return &x
}

// ЗАДАЧА 38: Что выведет? (указатель на composite literal)
func task38() {
	p := &Point{x: 1, y: 2}
	fmt.Println(*p)
}

// ЗАДАЧА 39: Что выведет? (изменение через метод)
func task39() {
	b := Box{value: 5}
	b.setValue(10)
	fmt.Println(b.value)
	(&b).setValue(15)
	fmt.Println(b.value)
}

// ЗАДАЧА 40: Что выведет? (указатель на элемент массива)
func task40() {
	arr := [3]int{1, 2, 3}
	p := &arr[1]
	*p = 100
	fmt.Println(arr)
}

// ЗАДАЧА 41: Что выведет? (nil pointer в структуре)
func task41() {
	n := Node{value: 1, next: nil}
	fmt.Println(n.next == nil)
}

// ЗАДАЧА 42: Что выведет? (defer и указатели)
func task42() {
	x := 1
	defer fmt.Println(*getPtr(&x))
	x = 2
}
func getPtr(p *int) *int {
	return p
}

// ЗАДАЧА 43: Что выведет? (указатель на результат функции - невозможно!)
func task43() {
	x := getValue()
	p := &x
	fmt.Println(*p)
}
func getValue() int {
	return 42
}

// ЗАДАЧА 44: Что выведет? (цепочка указателей)
func task44() {
	n3 := &Node{value: 3, next: nil}
	n2 := &Node{value: 2, next: n3}
	n1 := &Node{value: 1, next: n2}
	fmt.Println(n1.value, n1.next.value, n1.next.next.value)
}

// ЗАДАЧА 45: Что выведет? (указатель меняет адрес)
func task45() {
	x, y := 10, 20
	p := &x
	fmt.Println(*p)
	p = &y
	fmt.Println(*p)
}

// ЗАДАЧА 46: Что выведет? (встроенная структура и указатели)
func task46() {
	o := Outer{Inner: Inner{value: 42}, name: "test"}
	p := &o.Inner
	p.value = 100
	fmt.Println(o.value)
}

// ЗАДАЧА 47: Что выведет? (map значений-указателей)
func task47() {
	type Person47 struct{ name string }
	m := map[int]*Person47{
		1: {name: "Alice"},
		2: {name: "Bob"},
	}
	m[1].name = "Charlie"
	fmt.Println(m[1].name)
}

// ЗАДАЧА 48: Что выведет? (указатель после append)
func task48() {
	s := []int{1, 2, 3}
	p := &s[1]
	*p = 100
	s = append(s, 4, 5, 6, 7, 8, 9, 10)
	*p = 200
	fmt.Println(s)
}

// ЗАДАЧА 49: Что выведет? (копирование vs указатель в функции)
func task49() {
	d := Data{values: []int{1, 2, 3}}
	modifyDataValue49(d)
	fmt.Println(d.values)
	modifyDataPointer49(&d)
	fmt.Println(d.values)
}
func modifyDataValue49(d Data) {
	d.values[0] = 100
}
func modifyDataPointer49(d *Data) {
	d.values[1] = 200
}

// ЗАДАЧА 50: Что выведет? (указатели и escape analysis)
func task50() {
	p1 := createOnStack()
	p2 := createOnHeap()
	fmt.Println(*p1, *p2)
}
func createOnStack() *int {
	x := 10
	return &x // escape to heap
}
func createOnHeap() *int {
	return new(int) // allocated on heap
}
