package main

import (
	"fmt"
	"sync"
	"time"
)

// call представляет собой один выполняющийся или выполненный запрос.
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// RequestDeduplicator дедуплицирует параллельные запросы с одинаковым ключом.
type RequestDeduplicator struct {
	mu    sync.Mutex
	calls map[string]*call
}

// New создает новый экземпляр RequestDeduplicator.
func New() *RequestDeduplicator {
	// YOUR CODE HERE
	// Инициализируйте карту calls и верните указатель на RequestDeduplicator
	return nil
}

// Do гарантирует, что для данного ключа функция fn будет выполнена только один раз
// среди всех конкурентных вызовов с этим ключом. Остальные вызовы будут ждать
// завершения первого и получат тот же результат.
// После завершения fn запись о ключе удаляется, чтобы будущие вызовы снова вызывали fn.
func (rd *RequestDeduplicator) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	// YOUR CODE HERE
	// Реализуйте логику дедупликации, используя rd.mu, rd.calls и call.wg
	return nil, nil
}

// Пример использования (не требует изменений)
func main() {
	rd := New()
	var wg sync.WaitGroup
	keys := [2]string{"Moscow", "Piter"}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			result, err := rd.Do(keys[i%2], func() (interface{}, error) {
				fmt.Printf("  [%d] Выполняется реальный вызов API...\n", id)
				time.Sleep(2 * time.Second)
				return "Weather: sunny", nil
			})
			fmt.Printf("  [%d] Результат %d: %v, err: %v\n", id, i%2, result, err)
		}(i)
	}
	wg.Wait()

	// Проверка, что следующий вызов запустит функцию снова (кеш не хранится)
	time.Sleep(1 * time.Second)
	fmt.Println("\nНовый вызов через 1 секунду:")
	result, err := rd.Do(keys[0], func() (interface{}, error) {
		fmt.Println("  Реальный вызов API (новый)")
		return "Weather: cloudy", nil
	})
	fmt.Printf("Результат: %v, err: %v\n", result, err)
}
