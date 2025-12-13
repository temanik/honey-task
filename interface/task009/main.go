package main

// Реализуйте интерфейс CircuitBreaker для защиты от каскадных сбоев.
// CircuitBreaker имеет три состояния: Closed, Open, HalfOpen.
// - Closed: все запросы проходят, при превышении порога ошибок -> Open
// - Open: все запросы отклоняются, через timeout -> HalfOpen
// - HalfOpen: пропускается ограниченное число запросов для проверки, при успехе -> Closed

type State string

const (
	StateClosed   State = "closed"
	StateOpen     State = "open"
	StateHalfOpen State = "half-open"
)

type CircuitBreaker interface {
	// Call выполняет функцию через circuit breaker
	Call(fn func() (interface{}, error)) (interface{}, error)
	// State возвращает текущее состояние
	State() State
	// Reset сбрасывает circuit breaker в состояние Closed
	Reset()
}

// Реализуйте структуру SimpleCircuitBreaker с конструктором:
// func NewCircuitBreaker(maxFailures int, timeout time.Duration) *SimpleCircuitBreaker
// maxFailures - количество ошибок для перехода в Open
// timeout - время в Open состоянии до перехода в HalfOpen
