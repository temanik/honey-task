package main

import "context"

// Реализуйте интерфейс Connection Pool для управления пулом соединений.
// Pool должен переиспользовать соединения, контролировать их количество и lifetime.

type Connection interface {
	// Execute выполняет операцию через соединение
	Execute(query string) (interface{}, error)
	// Close закрывает соединение
	Close() error
	// IsValid проверяет, валидно ли соединение
	IsValid() bool
}

type ConnectionFactory interface {
	// Create создает новое соединение
	Create() (Connection, error)
}

type Pool interface {
	// Acquire получает соединение из пула
	Acquire(ctx context.Context) (Connection, error)
	// Release возвращает соединение в пул
	Release(conn Connection) error
	// Stats возвращает статистику пула
	Stats() PoolStats
	// Close закрывает пул и все соединения
	Close() error
}

type PoolStats struct {
	TotalConnections  int
	IdleConnections   int
	ActiveConnections int
}

// Реализуйте структуру ConnectionPool с конструктором:
// func NewConnectionPool(factory ConnectionFactory, minSize, maxSize int) *ConnectionPool
// minSize - минимальное количество соединений
// maxSize - максимальное количество соединений
//
// Дополнительно:
// - Поддержка timeout при Acquire
// - Проверка валидности соединения перед выдачей
// - Автоматическое переоткрытие невалидных соединений
