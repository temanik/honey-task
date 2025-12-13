package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для распределенного rate limiter с поддержкой квот и burst.
// Rate limiter должен работать в распределенной среде и обеспечивать точность ограничений.
//
// Сложность:
// - Распределенный учет запросов (синхронизация между нодами)
// - Различные алгоритмы: Token Bucket, Leaky Bucket, Fixed/Sliding Window
// - Поддержка burst (кратковременное превышение лимита)
// - Квоты на разных уровнях (per user, per IP, per API key, global)
// - Динамическое изменение лимитов без перезапуска
// - Fair queueing при достижении лимита

type RateLimitAlgorithm string

const (
	TokenBucket   RateLimitAlgorithm = "token_bucket"
	LeakyBucket   RateLimitAlgorithm = "leaky_bucket"
	FixedWindow   RateLimitAlgorithm = "fixed_window"
	SlidingWindow RateLimitAlgorithm = "sliding_window"
)

type RateLimit struct {
	Rate      int           // запросов в период
	Period    time.Duration // период (например, 1 секунда)
	Burst     int           // максимальный burst
	Algorithm RateLimitAlgorithm
}

type QuotaLevel string

const (
	QuotaGlobal  QuotaLevel = "global"
	QuotaPerUser QuotaLevel = "per_user"
	QuotaPerIP   QuotaLevel = "per_ip"
	QuotaPerKey  QuotaLevel = "per_key"
)

type Request struct {
	UserID    string
	IP        string
	APIKey    string
	Resource  string
	Timestamp time.Time
}

type RateLimitResult struct {
	Allowed      bool
	Remaining    int
	ResetAt      time.Time
	RetryAfter   time.Duration
	CurrentUsage int64
	QuotaLimit   int64
}

type DistributedRateLimiter interface {
	// Allow проверяет, разрешен ли запрос
	Allow(ctx context.Context, req Request) (RateLimitResult, error)
	// Reserve резервирует N запросов (для batch операций)
	Reserve(ctx context.Context, req Request, count int) (RateLimitResult, error)
	// Wait блокирует до тех пор, пока запрос не будет разрешен
	Wait(ctx context.Context, req Request) error
	// GetUsage возвращает текущее использование квоты
	GetUsage(ctx context.Context, level QuotaLevel, identifier string) (int64, error)
	// ResetQuota сбрасывает квоту для указанного идентификатора
	ResetQuota(ctx context.Context, level QuotaLevel, identifier string) error
}

type QuotaManager interface {
	// SetLimit устанавливает лимит для уровня и идентификатора
	SetLimit(ctx context.Context, level QuotaLevel, identifier string, limit RateLimit) error
	// GetLimit получает лимит
	GetLimit(ctx context.Context, level QuotaLevel, identifier string) (*RateLimit, error)
	// DeleteLimit удаляет лимит
	DeleteLimit(ctx context.Context, level QuotaLevel, identifier string) error
	// ListLimits возвращает все активные лимиты
	ListLimits(ctx context.Context) (map[string]RateLimit, error)
}

type RateLimitStore interface {
	// Increment увеличивает счетчик
	Increment(ctx context.Context, key string, window time.Duration) (int64, error)
	// IncrementBy увеличивает счетчик на указанное значение
	IncrementBy(ctx context.Context, key string, value int64, window time.Duration) (int64, error)
	// Get получает текущее значение счетчика
	Get(ctx context.Context, key string) (int64, error)
	// Reset сбрасывает счетчик
	Reset(ctx context.Context, key string) error
	// GetWindow получает все записи в окне (для sliding window)
	GetWindow(ctx context.Context, key string, start, end time.Time) ([]int64, error)
}

type FairQueue interface {
	// Enqueue добавляет запрос в очередь ожидания
	Enqueue(ctx context.Context, req Request, priority int) error
	// Dequeue получает следующий запрос из очереди
	Dequeue(ctx context.Context) (*Request, error)
	// Size возвращает размер очереди
	Size(ctx context.Context) (int, error)
}

type RateLimitMetrics struct {
	TotalRequests    int64
	AllowedRequests  int64
	BlockedRequests  int64
	CurrentWaiting   int64
	AvgWaitTime      time.Duration
	QuotaUtilization map[string]float64 // level:identifier -> utilization %
}

type AdaptiveRateLimiter interface {
	DistributedRateLimiter
	// AdjustLimit динамически изменяет лимит на основе нагрузки
	AdjustLimit(ctx context.Context, level QuotaLevel, identifier string, factor float64) error
	// GetMetrics возвращает метрики
	GetMetrics(ctx context.Context) (RateLimitMetrics, error)
}

// Реализуйте:
// 1. TokenBucketLimiter - реализация Token Bucket алгоритма
// 2. SlidingWindowLimiter - реализация Sliding Window алгоритма
// 3. HierarchicalQuotaManager - менеджер квот с иерархией (global -> user -> resource)
// 4. DistributedRateLimitStore - распределенное хранилище счетчиков
//    (можно использовать sync.Map для упрощения, но структура должна быть расширяема для Redis)
// 5. AdaptiveLimiter - rate limiter с автоматической адаптацией лимитов
//
// Подсказки:
// - Для точности в распределенной среде можно использовать векторные часы
// - Sliding window можно реализовать через sorted set с timestamp
// - Fair queue можно реализовать через priority queue с round-robin
