package main

import "time"

// Реализуйте интерфейс Metrics для сбора метрик приложения.
// Metrics должен поддерживать счетчики, гистограммы и gauge метрики.

type MetricType string

const (
	Counter   MetricType = "counter"
	Gauge     MetricType = "gauge"
	Histogram MetricType = "histogram"
)

type Metric struct {
	Name      string
	Type      MetricType
	Value     float64
	Labels    map[string]string
	Timestamp time.Time
}

type Metrics interface {
	// Inc увеличивает счетчик на 1
	Inc(name string, labels map[string]string)
	// Add увеличивает счетчик на значение
	Add(name string, value float64, labels map[string]string)
	// Set устанавливает значение gauge
	Set(name string, value float64, labels map[string]string)
	// Observe добавляет наблюдение в гистограмму
	Observe(name string, value float64, labels map[string]string)
	// GetAll возвращает все собранные метрики
	GetAll() []Metric
	// Reset сбрасывает все метрики
	Reset()
}

// Реализуйте структуру SimpleMetrics и декоратор BufferedMetrics:
// 1. SimpleMetrics - базовая реализация
// 2. BufferedMetrics - буферизует метрики и отправляет их батчами
//    func NewBufferedMetrics(underlying Metrics, bufferSize int, flushInterval time.Duration) *BufferedMetrics
