// Package structs some tasks
package structs

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type Person struct {
	Name string
	age  int
}

type Employee struct {
	name   string
	salary int
}

type Counter struct {
	value int
}

type Product struct {
	ID    int
	name  string
	Price float64
}

type Data struct {
	Value int    `json: "value"`
	Name  string `json:"name"`
}

type Config struct {
	Host string `json:"host"`
	port int    `json:"port"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	data    string `json:"data"`
}

type Address struct {
	City    string
	Country string
}

type Contact struct {
	Email string
	Phone string
}

type User struct {
	Address
	Contact
	Name string
}

type Base struct {
	ID   int
	Name string
}

type Extended struct {
	Base
	Name  string
	Extra string
}

type Inner struct {
	Value int
}

type Outer struct {
	*Inner
	Label string
}

type Node struct {
	Value int
	Next  *Node
}

type Point struct {
	X, Y int
}

type Metrics struct {
	Count int
}

type Stats struct {
	Metrics
	Total int
}

const Count = 50

func GetTasks() map[int]func() {
	return map[int]func(){
		1: task1, 2: task2, 3: task3, 4: task4, 5: task5, 6: task6, 7: task7, 8: task8, 9: task9, 10: task10,
		11: task11, 12: task12, 13: task13, 14: task14, 15: task15, 16: task16, 17: task17, 18: task18, 19: task19, 20: task20,
		21: task21, 22: task22, 23: task23, 24: task24, 25: task25, 26: task26, 27: task27, 28: task28, 29: task29, 30: task30,
		31: task31, 32: task32, 33: task33, 34: task34, 35: task35, 36: task36, 37: task37, 38: task38, 39: task39, 40: task40,
		41: task41, 42: task42, 43: task43, 44: task44, 45: task45, 46: task46, 47: task47, 48: task48, 49: task49, 50: task50,
	}
}

// ЗАДАЧА 1: Что выведет?
func task1() {
	p := &Person{Name: "Bob", age: 30}
	fmt.Println(p.Name)
	changePersonPtr(p)
	fmt.Println(p.Name)
}
func changePersonPtr(p *Person) {
	p = &Person{Name: "Alice", age: 25}
}

// ЗАДАЧА 2: Что выведет?
func task2() {
	p := &Person{Name: "Bob", age: 30}
	fmt.Println(p.Name, p.age)
	modifyPerson(p)
	fmt.Println(p.Name, p.age)
}
func modifyPerson(p *Person) {
	p.Name = "Alice"
	p.age = 25
}

// ЗАДАЧА 3: Что выведет?
func task3() {
	e := Employee{name: "John", salary: 5000}
	b, _ := json.Marshal(e)
	fmt.Println(string(b))
}

// ЗАДАЧА 4: Что выведет?
func task4() {
	d := Data{Value: 42, Name: "test"}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

// ЗАДАЧА 5: Что выведет?
func task5() {
	c := Config{Host: "localhost", port: 8080}
	b, _ := json.Marshal(c)
	fmt.Println(string(b))
	var out Config
	json.Unmarshal(b, &out)
	fmt.Printf("%+v\n", out)
}

// ЗАДАЧА 6: Что выведет?
func task6() {
	p := Product{ID: 1, name: "Laptop", Price: 999.99}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}

// ЗАДАЧА 7: Что выведет?
func task7() {
	p1 := Person{Name: "Alice", age: 30}
	p2 := p1
	p2.Name = "Bob"
	p2.age = 25
	fmt.Println(p1.Name, p1.age)
	fmt.Println(p2.Name, p2.age)
}

// ЗАДАЧА 8: Что выведет?
func task8() {
	p1 := &Person{Name: "Alice", age: 30}
	p2 := p1
	p2.Name = "Bob"
	p2.age = 25
	fmt.Println(p1.Name, p1.age)
}

// ЗАДАЧА 9: Что выведет?
func task9() {
	e := Extended{
		Base:  Base{ID: 1, Name: "Base"},
		Name:  "Extended",
		Extra: "Data",
	}
	fmt.Println(e.Name)
	fmt.Println(e.Base.Name)
}

// ЗАДАЧА 10: Что выведет?
func task10() {
	u := User{
		Address: Address{City: "Moscow", Country: "Russia"},
		Contact: Contact{Email: "test@mail.com", Phone: "123"},
		Name:    "Alice",
	}
	fmt.Println(u.Name)
	fmt.Println(u.City)
	fmt.Println(u.Email)
}

// ЗАДАЧА 11: Что выведет?
func task11() {
	o := Outer{Label: "test"}
	fmt.Println(o.Inner == nil)
	fmt.Println(o.Label)
	fmt.Println(o.Value)
}

// ЗАДАЧА 12: Что выведет?
func task12() {
	o := Outer{
		Inner: &Inner{Value: 42},
		Label: "test",
	}
	o.Value = 100
	fmt.Println(o.Inner.Value)
	fmt.Println(o.Value)
}

// ЗАДАЧА 13: Что выведет?
func task13() {
	persons := []Person{
		{Name: "Alice", age: 30},
		{Name: "Bob", age: 25},
	}
	for _, p := range persons {
		p.Name = "Changed"
	}
	fmt.Println(persons[0].Name)
}

// ЗАДАЧА 14: Что выведет?
func task14() {
	persons := []Person{
		{Name: "Alice", age: 30},
		{Name: "Bob", age: 25},
	}
	for i := range persons {
		persons[i].Name = "Changed"
	}
	fmt.Println(persons[0].Name)
}

// ЗАДАЧА 15: Что выведет?
func task15() {
	persons := []*Person{
		{Name: "Alice", age: 30},
		{Name: "Bob", age: 25},
	}
	for _, p := range persons {
		p.Name = "Changed"
	}
	fmt.Println(persons[0].Name)
}

// ЗАДАЧА 16: Что выведет?
func task16() {
	m := map[string]Person{
		"alice": {Name: "Alice", age: 30},
	}
	// m["alice"].age = 31 // ошибка компиляции!
	p := m["alice"]
	p.age = 31
	m["alice"] = p
	fmt.Println(m["alice"].age)
}

// ЗАДАЧА 17: Что выведет?
func task17() {
	m := map[string]*Person{
		"alice": {Name: "Alice", age: 30},
	}
	m["alice"].age = 31
	fmt.Println(m["alice"].age)
}

// ЗАДАЧА 18: Что выведет?
func task18() {
	var p Person
	fmt.Printf("%#v\n", p)
	fmt.Println(p.Name == "")
	fmt.Println(p.age == 0)
}

// ЗАДАЧА 19: Что выведет?
func task19() {
	p := new(Person)
	fmt.Println(p.Name == "")
	fmt.Println(p.age == 0)
	fmt.Printf("%T\n", p)
}

// ЗАДАЧА 20: Что выведет?
func task20() {
	p1 := Person{Name: "Alice", age: 30}
	p2 := Person{Name: "Alice", age: 30}
	fmt.Println(p1 == p2)
	p3 := Person{Name: "Alice", age: 31}
	fmt.Println(p1 == p3)
}

// ЗАДАЧА 21: Что выведет?
func task21() {
	p1 := &Person{Name: "Alice", age: 30}
	p2 := &Person{Name: "Alice", age: 30}
	fmt.Println(p1 == p2)
	fmt.Println(*p1 == *p2)
}

// ЗАДАЧА 22: Что выведет?
func task22() {
	c := Counter{value: 5}
	c.increment()
	fmt.Println(c.value)
}
func (c Counter) increment() {
	c.value++
}

// ЗАДАЧА 23: Что выведет?
func task23() {
	c := Counter{value: 5}
	c.incrementPtr()
	fmt.Println(c.value)
}
func (c *Counter) incrementPtr() {
	c.value++
}

// ЗАДАЧА 24: Что выведет?
func task24() {
	c := Counter{value: 5}
	c.incrementPtr()
	fmt.Println(c.value)
	(&c).incrementPtr()
	fmt.Println(c.value)
}

// ЗАДАЧА 25: Что выведет?
func task25() {
	jsonStr := `{"name":"John","salary":5000}`
	var e Employee
	json.Unmarshal([]byte(jsonStr), &e)
	fmt.Printf("%+v\n", e)
}

// ЗАДАЧА 26: Что выведет?
func task26() {
	jsonStr := `{"code":200,"msg":"OK","data":"secret","extra":"ignored"}`
	var r Response
	json.Unmarshal([]byte(jsonStr), &r)
	fmt.Printf("%+v\n", r)
}

// ЗАДАЧА 27: Что выведет?
func task27() {
	p := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("%+v\n", p)
}

// ЗАДАЧА 28: Что выведет?
func task28() {
	p1 := struct {
		Name string
		Age  int
	}{"Alice", 30}
	p2 := struct {
		Name string
		Age  int
	}{"Alice", 30}
	fmt.Println(p1 == p2)
}

// ЗАДАЧА 29: Что выведет?
func task29() {
	type Data struct {
		Name  string `json:"name,omitempty"`
		Value int    `json:"value,omitempty"`
	}
	d1 := Data{Name: "test", Value: 0}
	d2 := Data{Name: "", Value: 42}
	b1, _ := json.Marshal(d1)
	b2, _ := json.Marshal(d2)
	fmt.Println(string(b1))
	fmt.Println(string(b2))
}

// ЗАДАЧА 30: Что выведет?
func task30() {
	type Data struct {
		Public  string `json:"public"`
		Private string `json:"-"`
	}
	d := Data{Public: "visible", Private: "hidden"}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

// ЗАДАЧА 31: Что выведет?
func task31() {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n1.Next = n2
	n2.Next = n3
	fmt.Println(n1.Value)
	fmt.Println(n1.Next.Value)
	fmt.Println(n1.Next.Next.Value)
}

// ЗАДАЧА 32: Что выведет?
func task32() {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n1.Next = n2
	n2.Next = n1
	fmt.Println(n1.Value)
	fmt.Println(n1.Next.Value)
	fmt.Println(n1.Next.Next.Value)
}

// ЗАДАЧА 33: Что выведет?
func task33() {
	type Box struct {
		Value *int
	}
	x := 10
	b1 := Box{Value: &x}
	b2 := b1
	*b2.Value = 20
	fmt.Println(*b1.Value)
	fmt.Println(x)
}

// ЗАДАЧА 34: Что выведет?
func task34() {
	type Data struct {
		Values []int
	}
	d1 := Data{Values: []int{1, 2, 3}}
	d2 := d1
	d2.Values[0] = 100
	fmt.Println(d1.Values[0])
}

// ЗАДАЧА 35: Что выведет?
func task35() {
	type Base struct {
		ID int `json:"id"`
	}
	type Extended struct {
		Base
		Name string `json:"name"`
	}
	e := Extended{Base: Base{ID: 1}, Name: "test"}
	b, _ := json.Marshal(e)
	fmt.Println(string(b))
}

// ЗАДАЧА 36: Что выведет?
func task36() {
	type Base struct {
		ID int `json:"id"`
	}
	type Extended struct {
		*Base
		Name string `json:"name"`
	}
	e := Extended{Base: &Base{ID: 1}, Name: "test"}
	b, _ := json.Marshal(e)
	fmt.Println(string(b))
}

// ЗАДАЧА 37: Что выведет?
func task37() {
	type A struct {
		a bool
		b int64
		c bool
	}
	type B struct {
		a bool
		c bool
		b int64
	}
	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
}

// ЗАДАЧА 38: Что выведет?
func task38() {
	type Empty struct{}
	fmt.Println(unsafe.Sizeof(Empty{}))
	arr := [1000000]Empty{}
	fmt.Println(unsafe.Sizeof(arr))
}

// ЗАДАЧА 39: Что выведет?
func task39() {
	type Data struct {
		ID    int    `json:"id,string"`
		Value string `json:"value"`
	}
	d := Data{ID: 123, Value: "test"}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

// ЗАДАЧА 40: Что выведет?
func task40() {
	type Data struct {
		Value int `json:"value"`
	}
	jsonStr := `{"value":"not a number"}`
	var d Data
	err := json.Unmarshal([]byte(jsonStr), &d)
	fmt.Println(err != nil)
	fmt.Println(d.Value)
}

// ЗАДАЧА 41: Что выведет?
func task41() {
	type A struct {
		Value int
	}
	type B struct {
		Value int
	}
	type C struct {
		A
		B
	}
	c := C{A: A{Value: 1}, B: B{Value: 2}}
	// fmt.Println(c.Value) // ошибка: ambiguous selector
	fmt.Println(c.A.Value, c.B.Value)
}

// ЗАДАЧА 42: Что выведет?
func task42() {
	s := Stats{
		Metrics: Metrics{Count: 10},
		Total:   100,
	}
	s.increase()
	fmt.Println(s.Count)
	fmt.Println(s.Metrics.Count)
}
func (m *Metrics) increase() {
	m.Count++
}

// ЗАДАЧА 43: Что выведет?
func task43() {
	p := &Person{Name: "Bob", age: 30}
	changePtr(&p)
	fmt.Println(p.Name)
}
func changePtr(pp **Person) {
	*pp = &Person{Name: "Alice", age: 25}
}

// ЗАДАЧА 44: Что выведет?
func task44() {
	p := Point{10, 20}
	fmt.Println(p.X, p.Y)
	//q := Point{20} что тут будет?
	//fmt.Println(q)
}

// ЗАДАЧА 45: Что выведет?
func task45() {
	p := Point{X: 10}
	fmt.Println(p.X, p.Y)
}

// ЗАДАЧА 46: Что выведет?
func task46() {
	type Data struct {
		Name string          `json:"name"`
		Raw  json.RawMessage `json:"raw"`
	}
	jsonStr := `{"name":"test","raw":{"nested":"value"}}`
	var d Data
	json.Unmarshal([]byte(jsonStr), &d)
	fmt.Println(d.Name)
	fmt.Println(string(d.Raw))
}

// ЗАДАЧА 47: Что выведет?
func task47() {
	type Response struct {
		Status int `json:"status"`
		Data   struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"data"`
	}
	r := Response{Status: 200}
	r.Data.Name = "test"
	r.Data.Value = 42
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

// ЗАДАЧА 48: Что выведет?
func task48() {
	type Data struct {
		Value int    `json: "value" xml:"val"`
		Name  string `json:"name"`
	}
	d := Data{Value: 42, Name: "test"}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}

// ЗАДАЧА 49: Что выведет?
func task49() {
	type Inner struct {
		Value int
	}
	type Outer struct {
		Data Inner
		Name string
	}
	var o Outer
	fmt.Printf("%+v\n", o)
	fmt.Println(o.Data.Value)
}

// ЗАДАЧА 50: Что выведет?
func task50() {
	type PersonA struct {
		Name string
		Age  int
	}
	type PersonB struct {
		Name string
		Age  int
	}
	a := PersonA{Name: "Alice", Age: 30}
	// b := PersonB(a) // ошибка компиляции!
	b := PersonB{Name: a.Name, Age: a.Age}
	fmt.Printf("%+v\n", b)
}
