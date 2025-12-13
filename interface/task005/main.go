package main

// Реализуйте интерфейс RateLimiter для ограничения скорости выполнения операций.
// Используйте алгоритм Token Bucket.
// RateLimiter должен быть потокобезопасным.

type RateLimiter interface {
	// Allow проверяет, можно ли выполнить операцию сейчас
	Allow() bool
	// Wait блокирует выполнение до тех пор, пока операция не станет разрешенной
	Wait() error
	// WaitN блокирует выполнение до тех пор, пока не будет доступно n токенов
	WaitN(n int) error
}

// Реализуйте структуру TokenBucket со следующим конструктором:
// func NewTokenBucket(rate int, capacity int) *TokenBucket
// где rate - количество токенов, добавляемых в секунду
// capacity - максимальное количество токенов в bucket
