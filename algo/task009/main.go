package main

import (
	"container/heap"
	"fmt"
	"slices"
)

/*
В массиве А хранятся цены на N предметов. Есть K купонов, которые уменьшают цену предмета на X.
Если применить купонов на предмет с ценой a, то его итоговая стоимость будет max(a - t*x, 0) (то есть купоны не могут сделать цену предмета отрицательной).
 Необходимо вернуть минимальное количество денег, которое придется потратить, чтобы купить все предметы. A = 8, 3, 10, 5, 13 K = 4 X = 7 result = 12

8, 3, 10, 5, 13
1, 3, 3, 5, 0
*/

func apply(items []int, k, x int) int {
	if len(items) == 0 {
		return 0
	}

	for range k {
		slices.SortFunc(items, func(a, b int) int {
			return b - a
		})

		items[0] = max(items[0]-x, 0)
	}

	result := 0
	for _, item := range items {
		result += item
	}

	return result
}

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}
func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func applyHeap(items []int, k, x int) int {
	h := MaxHeap(items)
	heap.Init(&h)

	for range k {
		item := heap.Pop(&h).(int)
		item = max(item-x, 0)
		heap.Push(&h, item)
	}

	sum := 0
	for _, v := range h {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(apply([]int{8, 3, 10, 5, 13}, 4, 7))
	fmt.Println(applyHeap([]int{8, 3, 10, 5, 13}, 4, 7))
}
