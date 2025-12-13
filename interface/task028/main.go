package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для Content Delivery Network (CDN) edge cache.
// Система должна эффективно кэшировать и доставлять контент с учетом географии и актуальности.
//
// Сложность:
// - Multi-tier caching (edge -> regional -> origin)
// - Cache coherence protocol (инвалидация кэша на всех узлах)
// - Content routing на основе геолокации и латентности
// - Поддержка Range requests для больших файлов
// - Cache warming и prefetching
// - Conditional requests (ETag, Last-Modified)

type Content struct {
	Key          string
	Data         []byte
	ContentType  string
	ETag         string
	LastModified time.Time
	CacheControl CacheControl
	Metadata     map[string]string
}

type CacheControl struct {
	MaxAge               time.Duration
	SMaxAge              time.Duration // для shared cache
	MustRevalidate       bool
	NoCache              bool
	NoStore              bool
	Public               bool
	Private              bool
	Immutable            bool
	StaleWhileRevalidate time.Duration
}

type CacheTier string

const (
	TierEdge     CacheTier = "edge"     // ближайший к клиенту
	TierRegional CacheTier = "regional" // региональный
	TierOrigin   CacheTier = "origin"   // источник
)

type Location struct {
	Latitude   float64
	Longitude  float64
	Region     string
	DataCenter string
}

type EdgeCache interface {
	// Get получает контент из кэша
	Get(ctx context.Context, key string, opts RequestOptions) (*Content, CacheStatus, error)
	// Put сохраняет контент в кэш
	Put(ctx context.Context, content Content) error
	// Invalidate инвалидирует контент
	Invalidate(ctx context.Context, key string) error
	// InvalidatePattern инвалидирует по паттерну (например, /images/*)
	InvalidatePattern(ctx context.Context, pattern string) error
	// Purge немедленно удаляет из всех тиров
	Purge(ctx context.Context, key string) error
}

type RequestOptions struct {
	IfNoneMatch     string // ETag для условного запроса
	IfModifiedSince *time.Time
	Range           *RangeSpec
	ClientLocation  *Location
	AcceptEncoding  []string
}

type RangeSpec struct {
	Start int64
	End   int64 // -1 означает до конца
}

type CacheStatus string

const (
	StatusHit         CacheStatus = "hit"         // найдено в кэше
	StatusMiss        CacheStatus = "miss"        // не найдено
	StatusStale       CacheStatus = "stale"       // устаревшее
	StatusRevalidated CacheStatus = "revalidated" // ревалидировано
	StatusBypassed    CacheStatus = "bypassed"    // кэш обойден
)

type MultiTierCache interface {
	EdgeCache
	// GetFromTier получает из конкретного тира
	GetFromTier(ctx context.Context, tier CacheTier, key string) (*Content, error)
	// Promote продвигает контент в более высокий тир
	Promote(ctx context.Context, key string, fromTier, toTier CacheTier) error
	// GetTierStats возвращает статистику по тиру
	GetTierStats(ctx context.Context, tier CacheTier) (TierStats, error)
}

type TierStats struct {
	Tier       CacheTier
	HitRate    float64
	Size       int64
	ItemCount  int64
	Evictions  int64
	AvgLatency time.Duration
}

type CacheCoherence interface {
	// Broadcast отправляет команду инвалидации всем узлам
	Broadcast(ctx context.Context, msg InvalidationMessage) error
	// Subscribe подписывается на сообщения инвалидации
	Subscribe(ctx context.Context, handler func(InvalidationMessage) error) error
	// GetClusterState возвращает состояние кластера
	GetClusterState(ctx context.Context) (ClusterState, error)
}

type InvalidationMessage struct {
	Type      InvalidationType
	Keys      []string
	Pattern   string
	Timestamp time.Time
	NodeID    string
}

type InvalidationType string

const (
	InvalidateKey     InvalidationType = "key"
	InvalidatePattern InvalidationType = "pattern"
	InvalidatePurge   InvalidationType = "purge"
)

type ClusterState struct {
	Nodes        []NodeInfo
	TotalSize    int64
	TotalItems   int64
	Synchronized bool
}

type NodeInfo struct {
	ID          string
	Location    Location
	Tier        CacheTier
	Status      NodeStatus
	Load        float64
	Latency     time.Duration
	LastContact time.Time
}

type NodeStatus string

const (
	NodeStatusHealthy   NodeStatus = "healthy"
	NodeStatusDegraded  NodeStatus = "degraded"
	NodeStatusUnhealthy NodeStatus = "unhealthy"
)

type ContentRouter interface {
	// Route выбирает оптимальный узел для запроса
	Route(ctx context.Context, key string, clientLocation Location) (*NodeInfo, error)
	// GetClosestNodes возвращает N ближайших узлов
	GetClosestNodes(ctx context.Context, location Location, count int) ([]NodeInfo, error)
	// UpdateRouting обновляет таблицу маршрутизации
	UpdateRouting(ctx context.Context) error
}

type Prefetcher interface {
	// Prefetch предзагружает контент в кэш
	Prefetch(ctx context.Context, keys []string, tier CacheTier) error
	// WarmUp прогревает кэш популярным контентом
	WarmUp(ctx context.Context, tier CacheTier) error
	// PredictAndPrefetch предсказывает и предзагружает
	PredictAndPrefetch(ctx context.Context, accessPattern AccessPattern) error
}

type AccessPattern struct {
	Key          string
	AccessCount  int64
	LastAccessed time.Time
	Pattern      []time.Time // история доступа
	Predicted    []string    // предсказанные следующие запросы
}

type CompressionHandler interface {
	// Compress сжимает контент
	Compress(ctx context.Context, content []byte, encoding string) ([]byte, error)
	// Decompress распаковывает контент
	Decompress(ctx context.Context, content []byte, encoding string) ([]byte, error)
	// NegotiateEncoding выбирает оптимальное сжатие
	NegotiateEncoding(acceptedEncodings []string) string
}

type OriginShield interface {
	// Shield защищает origin от множественных запросов
	Shield(ctx context.Context, key string, fetcher func() (*Content, error)) (*Content, error)
	// CollapseRequests объединяет одновременные запросы
	CollapseRequests(ctx context.Context, key string) (*Content, error)
}

type CDNMetrics struct {
	TotalRequests  int64
	CacheHitRate   float64
	BandwidthSaved int64
	AvgLatency     time.Duration
	OriginRequests int64
	TierHitRates   map[CacheTier]float64
	TopContent     []ContentStats
}

type ContentStats struct {
	Key         string
	AccessCount int64
	HitRate     float64
	Size        int64
}

// Реализуйте:
// 1. MultiTierEdgeCache - многоуровневый кэш
// 2. GeoRouter - роутер на основе геолокации
// 3. CacheCoherenceProtocol - протокол согласованности кэша
// 4. SmartPrefetcher - умная предзагрузка на основе паттернов
// 5. OriginShieldLayer - защита origin сервера
// 6. CompressionMiddleware - middleware для сжатия
//
// Подсказки:
// - Для геороутинга можно использовать haversine формулу для расчета расстояния
// - Cache coherence можно реализовать через pub/sub паттерн
// - Для prefetching используйте статистику доступа и ML паттерны
// - Origin shield можно реализовать через singleflight паттерн
