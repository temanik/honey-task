// ЗАДАЧА 15: RingBuffer
// Реализуйте методы структуры ringBuffer так, что main отрабатывал без паники

package main

import (
	"fmt"
	"reflect"
)

type ringBuffer struct {
	buff   chan int
	closed bool
}

func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{
		buff:   make(chan int, size),
		closed: false,
	}
}

func (b *ringBuffer) write(v int) {
	if b.closed {
		return
	}

	select {
	case b.buff <- v:
	default:
		<-b.buff
		b.buff <- v
	}
}

func (b *ringBuffer) close() {
	if !b.closed {
		b.closed = true
	}
	close(b.buff)
}

func (b *ringBuffer) read() (v int, ok bool) {
	v, ok = <-b.buff
	return
}

func main() {
	buff := newRingBuffer(3)

	for i := 1; i <= 6; i++ {
		buff.write(i)
	}

	buff.close()

	res := make([]int, 0)
	for {
		if v, ok := buff.read(); ok {
			res = append(res, v)
		} else {
			break
		}
	}

	if !reflect.DeepEqual(res, []int{4, 5, 6}) {
		panic(fmt.Sprintf("wrong code, res is %v", res))
	}
}
