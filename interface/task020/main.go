package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для распределенных блокировок (Distributed Lock).
// Lock должен обеспечивать взаимоисключающий доступ к ресурсу в распределенной системе.

type Lock interface {
	// Acquire пытается захватить блокировку
	Acquire(ctx context.Context, resource string, ttl time.Duration) (bool, error)
	// Release освобождает блокировку
	Release(ctx context.Context, resource string) error
	// Refresh продлевает время жизни блокировки
	Refresh(ctx context.Context, resource string, ttl time.Duration) error
	// IsLocked проверяет, заблокирован ли ресурс
	IsLocked(ctx context.Context, resource string) (bool, error)
}

type DistributedLock interface {
	Lock
	// AcquireWithRetry пытается захватить блокировку с повторными попытками
	AcquireWithRetry(ctx context.Context, resource string, ttl time.Duration, retries int) error
	// WithLock выполняет функцию с захваченной блокировкой
	WithLock(ctx context.Context, resource string, ttl time.Duration, fn func() error) error
}

// Реализуйте:
// 1. InMemoryLock - блокировка в памяти (для одного процесса)
// 2. RetryableLock - обертка над Lock с автоматическими повторными попытками
// 3. AutoRefreshLock - обертка, которая автоматически продлевает блокировку
//
// Для упрощения можно использовать sync.Mutex для InMemoryLock,
// но структура должна быть расширяема для работы с Redis/etcd в будущем.
