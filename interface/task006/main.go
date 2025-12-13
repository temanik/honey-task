package main

import "context"

// Реализуйте интерфейс Worker для пула воркеров, обрабатывающих задачи.
// WorkerPool должен ограничивать количество одновременно выполняющихся задач
// и поддерживать graceful shutdown через context.

type Task func() error

type WorkerPool interface {
	// Submit добавляет задачу в очередь на выполнение
	Submit(task Task) error
	// Start запускает пул воркеров
	Start(ctx context.Context) error
	// Shutdown корректно завершает работу пула, дожидаясь выполнения всех задач
	Shutdown() error
	// ShutdownNow немедленно останавливает пул, отменяя невыполненные задачи
	ShutdownNow() error
}

// Реализуйте структуру SimpleWorkerPool с конструктором:
// func NewWorkerPool(workerCount int, queueSize int) *SimpleWorkerPool
