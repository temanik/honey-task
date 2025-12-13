package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс Retry для повторного выполнения операций при ошибках.
// Поддержите различные стратегии повтора: константная задержка, экспоненциальная, с jitter.

type RetryStrategy interface {
	// NextDelay возвращает задержку перед следующей попыткой
	// attempt - номер попытки (начиная с 0)
	NextDelay(attempt int) time.Duration
}

type Retry interface {
	// Do выполняет операцию с повторами согласно стратегии
	Do(ctx context.Context, fn func() error) error
	// DoWithData выполняет операцию с возвратом данных
	DoWithData(ctx context.Context, fn func() (interface{}, error)) (interface{}, error)
}

// Реализуйте:
// 1. ConstantBackoff - константная задержка между попытками
// 2. ExponentialBackoff - экспоненциально растущая задержка
// 3. RetryExecutor с конструктором:
//    func NewRetryExecutor(maxAttempts int, strategy RetryStrategy) *RetryExecutor
