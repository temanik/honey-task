package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для распределенной очереди задач (Distributed Task Queue).
// Очередь должна поддерживать приоритеты, retry логику, dead letter queue и exactly-once семантику.
//
// Сложность:
// - Нужно обеспечить exactly-once delivery (задача не должна выполниться дважды)
// - Поддержка приоритетов задач
// - Автоматический retry с экспоненциальной задержкой
// - Dead Letter Queue для задач, которые не удалось выполнить после N попыток
// - Graceful shutdown с завершением текущих задач

type Priority int

const (
	PriorityLow Priority = iota
	PriorityNormal
	PriorityHigh
	PriorityCritical
)

type Task struct {
	ID          string
	Payload     interface{}
	Priority    Priority
	MaxRetries  int
	Timeout     time.Duration
	ScheduledAt time.Time
}

type TaskResult struct {
	TaskID    string
	Success   bool
	Result    interface{}
	Error     error
	Attempts  int
	StartedAt time.Time
	EndedAt   time.Time
}

type TaskHandler func(ctx context.Context, task Task) (interface{}, error)

type TaskQueue interface {
	// Enqueue добавляет задачу в очередь
	Enqueue(ctx context.Context, task Task) error
	// EnqueueBatch добавляет несколько задач атомарно
	EnqueueBatch(ctx context.Context, tasks []Task) error
	// Dequeue получает задачу из очереди (с учетом приоритета)
	Dequeue(ctx context.Context, timeout time.Duration) (*Task, error)
	// Complete помечает задачу как выполненную
	Complete(ctx context.Context, taskID string, result interface{}) error
	// Fail помечает задачу как проваленную (может вызвать retry)
	Fail(ctx context.Context, taskID string, err error) error
	// GetDeadLetterQueue возвращает задачи из DLQ
	GetDeadLetterQueue(ctx context.Context, limit int) ([]Task, error)
	// Stats возвращает статистику очереди
	Stats(ctx context.Context) (QueueStats, error)
}

type QueueStats struct {
	PendingTasks    int
	ProcessingTasks int
	CompletedTasks  int
	FailedTasks     int
	DLQTasks        int
}

type Worker interface {
	// RegisterHandler регистрирует обработчик задач
	RegisterHandler(taskType string, handler TaskHandler) error
	// Start запускает воркер
	Start(ctx context.Context) error
	// Stop останавливает воркер (graceful)
	Stop(ctx context.Context) error
}

// Реализуйте:
// 1. DistributedTaskQueue - основная реализация очереди
// 2. TaskWorker - воркер для обработки задач из очереди
// 3. TaskMonitor - мониторинг и статистика выполнения задач
//
// Подсказки:
// - Используйте каналы и контексты для координации
// - Для exactly-once можно использовать идемпотентные ключи
// - Приоритетную очередь можно реализовать через heap или несколько каналов
