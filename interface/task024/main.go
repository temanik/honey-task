package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для streaming data pipeline с поддержкой backpressure и windowing.
// Pipeline должен обрабатывать потоковые данные с поддержкой различных окон агрегации.
//
// Сложность:
// - Backpressure mechanism (замедление producer при перегрузке consumer)
// - Windowing: tumbling, sliding, session windows
// - Watermarks для обработки late arriving data
// - Exactly-once processing semantics
// - Поддержка checkpointing для fault tolerance

type Message struct {
	Key       string
	Value     interface{}
	Timestamp time.Time
	Metadata  map[string]string
}

type Window struct {
	Start time.Time
	End   time.Time
	Type  WindowType
}

type WindowType string

const (
	TumblingWindow WindowType = "tumbling" // фиксированные непересекающиеся окна
	SlidingWindow  WindowType = "sliding"  // перекрывающиеся окна
	SessionWindow  WindowType = "session"  // окна на основе активности
)

type Processor interface {
	// Process обрабатывает сообщение
	Process(ctx context.Context, msg Message) ([]Message, error)
}

type Aggregator interface {
	// Aggregate агрегирует сообщения в окне
	Aggregate(ctx context.Context, window Window, messages []Message) (interface{}, error)
}

type StreamSource interface {
	// Read читает сообщения из источника
	Read(ctx context.Context, maxBatch int) ([]Message, error)
	// Commit подтверждает обработку до указанного offset
	Commit(ctx context.Context, offset int64) error
	// Seek перемещает позицию чтения
	Seek(ctx context.Context, offset int64) error
}

type StreamSink interface {
	// Write записывает сообщения в приемник
	Write(ctx context.Context, messages []Message) error
	// Flush форсирует запись буферизованных данных
	Flush(ctx context.Context) error
}

type Pipeline interface {
	// AddSource добавляет источник данных
	AddSource(name string, source StreamSource) error
	// AddProcessor добавляет процессор в pipeline
	AddProcessor(name string, processor Processor) error
	// AddSink добавляет приемник данных
	AddSink(name string, sink StreamSink) error
	// AddWindow добавляет окно агрегации
	AddWindow(name string, windowType WindowType, size time.Duration, aggregator Aggregator) error
	// Start запускает pipeline
	Start(ctx context.Context) error
	// Stop останавливает pipeline
	Stop(ctx context.Context) error
	// Checkpoint сохраняет состояние pipeline
	Checkpoint(ctx context.Context) error
	// Restore восстанавливает состояние из checkpoint
	Restore(ctx context.Context, checkpointID string) error
}

type BackpressureStrategy interface {
	// ShouldBlock проверяет, нужно ли заблокировать producer
	ShouldBlock(ctx context.Context, queueSize int, maxSize int) bool
	// OnBackpressure вызывается при backpressure
	OnBackpressure(ctx context.Context) error
}

type Watermark interface {
	// UpdateWatermark обновляет watermark
	UpdateWatermark(timestamp time.Time)
	// GetWatermark возвращает текущий watermark
	GetWatermark() time.Time
	// IsLate проверяет, является ли сообщение опоздавшим
	IsLate(msgTimestamp time.Time) bool
}

type PipelineMetrics struct {
	MessagesProcessed  int64
	MessagesDropped    int64
	BytesProcessed     int64
	Lag                time.Duration
	Throughput         float64
	BackpressureEvents int64
}

// Реализуйте:
// 1. StreamingPipeline - основная реализация pipeline
// 2. WindowManager - управление окнами агрегации
// 3. BackpressureController - контроль backpressure
// 4. CheckpointManager - управление checkpoints для fault tolerance
//
// Пример использования:
// pipeline := NewStreamingPipeline()
// pipeline.AddSource("kafka", kafkaSource)
// pipeline.AddProcessor("filter", filterProcessor)
// pipeline.AddWindow("5min", TumblingWindow, 5*time.Minute, sumAggregator)
// pipeline.AddSink("clickhouse", chSink)
// pipeline.Start(ctx)
