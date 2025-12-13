package main

// Реализуйте интерфейс Middleware для построения цепочки обработчиков.
// Middleware должен поддерживать композицию и возможность прерывания цепочки.

type Context interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Next() error
	Abort(err error)
	IsAborted() bool
	Error() error
}

type Handler func(ctx Context) error

type Middleware interface {
	Handle(ctx Context) error
}

type MiddlewareChain interface {
	// Use добавляет middleware в цепочку
	Use(middleware Middleware) MiddlewareChain
	// Execute выполняет цепочку middleware
	Execute(ctx Context) error
}

// Реализуйте:
// 1. SimpleContext - базовая реализация Context
// 2. Chain - цепочка middleware
// 3. Несколько базовых middleware:
//    - LoggingMiddleware - логирует запросы
//    - RecoveryMiddleware - обрабатывает панику
//    - TimeoutMiddleware - ограничивает время выполнения
//    - AuthMiddleware - проверяет аутентификацию (stub)
