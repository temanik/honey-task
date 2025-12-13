package main

// Реализуйте интерфейс для потоковой обработки данных (Stream Processing).
// Stream должен поддерживать операции фильтрации, трансформации и агрегации.

type Stream interface {
	// Filter оставляет только элементы, удовлетворяющие предикату
	Filter(predicate func(interface{}) bool) Stream
	// Map применяет функцию трансформации к каждому элементу
	Map(transform func(interface{}) interface{}) Stream
	// Reduce агрегирует элементы в одно значение
	Reduce(initial interface{}, accumulator func(interface{}, interface{}) interface{}) interface{}
	// Collect собирает все элементы в слайс
	Collect() []interface{}
	// ForEach применяет функцию к каждому элементу
	ForEach(consumer func(interface{}))
	// Count возвращает количество элементов
	Count() int
}

// Реализуйте:
// 1. SliceStream - стрим на основе слайса
//    func NewSliceStream(items []interface{}) *SliceStream
// 2. ChannelStream - стрим на основе канала
//    func NewChannelStream(ch <-chan interface{}) *ChannelStream
//
// Дополнительно (опционально):
// 3. ParallelStream - параллельная обработка элементов
