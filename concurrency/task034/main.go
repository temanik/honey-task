package main

import (
	"fmt"
	"runtime"
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
	return &RequestDeduplicator{
		calls: make(map[string]*call),
	}
}

// Do гарантирует, что для данного ключа функция fn будет выполнена только один раз
// среди всех конкурентных вызовов с этим ключом. Остальные вызовы будут ждать
// завершения первого и получат тот же результат.
// После завершения fn запись о ключе удаляется, чтобы будущие вызовы снова вызывали fn.
func (rd *RequestDeduplicator) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	rd.mu.Lock()
	c, ok := rd.calls[key]
	if ok {
		rd.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c = &call{}
	c.wg.Add(1)
	rd.calls[key] = c
	rd.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	go func() {
		time.Sleep(3 * time.Second)
		rd.mu.Lock()
		delete(rd.calls, key)
		rd.mu.Unlock()
	}()

	return c.val, c.err
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
				runtime.Gosched()
				return "Weather: sunny", fmt.Errorf("api error")
			})
			fmt.Printf("  [%d] Результат %d: %v, err: %v\n", id, i%2, result, err)
		}(i)
	}
	wg.Wait()

	// Проверка, что следующий вызов запустит функцию снова (кеш хранится)
	time.Sleep(1 * time.Second)
	fmt.Println("\nНовый вызов через 1 секунду:")
	result, err := rd.Do(keys[0], func() (interface{}, error) {
		fmt.Println("  Реальный вызов API (новый)")
		return "Weather: cloudy", nil
	})
	fmt.Printf("Cached Результат: %v, err: %v\n", result, err)

	// Проверка, что следующий вызов запустит функцию снова (кеш удалился)
	time.Sleep(3 * time.Second)
	fmt.Println("\nНовый вызов через 4 секунду:")
	result, err = rd.Do(keys[0], func() (interface{}, error) {
		fmt.Println("  Реальный вызов API (новый)")
		return "Weather: cloudy", nil
	})
	fmt.Printf("New Результат: %v, err: %v\n", result, err)
}
