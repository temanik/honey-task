package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для распределенного кэша с поддержкой консистентности (Distributed Cache).
// Кэш должен поддерживать различные стратегии инвалидации, репликацию и консистентное хеширование.
//
// Сложность:
// - Консистентное хеширование для распределения ключей по нодам
// - Репликация данных (primary + replicas)
// - Write-through и Write-behind стратегии
// - Cache stampede protection (только один запрос к backend при cache miss)
// - Поддержка TTL и LRU eviction policy

type CacheEntry struct {
	Key        string
	Value      interface{}
	TTL        time.Duration
	CreatedAt  time.Time
	AccessedAt time.Time
	Version    int64
}

type EvictionPolicy string

const (
	EvictionLRU  EvictionPolicy = "lru"
	EvictionLFU  EvictionPolicy = "lfu"
	EvictionFIFO EvictionPolicy = "fifo"
)

type CacheStrategy string

const (
	WriteThrough CacheStrategy = "write_through"
	WriteBehind  CacheStrategy = "write_behind"
	WriteAround  CacheStrategy = "write_around"
)

type DistributedCache interface {
	// Get получает значение из кэша
	Get(ctx context.Context, key string) (interface{}, bool, error)
	// GetOrLoad получает из кэша или загружает через loader
	GetOrLoad(ctx context.Context, key string, loader func() (interface{}, error)) (interface{}, error)
	// Set сохраняет значение в кэш
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	// Delete удаляет значение из кэша
	Delete(ctx context.Context, key string) error
	// InvalidatePattern инвалидирует все ключи по паттерну
	InvalidatePattern(ctx context.Context, pattern string) error
	// GetMulti получает несколько значений атомарно
	GetMulti(ctx context.Context, keys []string) (map[string]interface{}, error)
	// SetMulti сохраняет несколько значений атомарно
	SetMulti(ctx context.Context, entries map[string]interface{}, ttl time.Duration) error
}

type Node interface {
	// ID возвращает идентификатор ноды
	ID() string
	// IsAlive проверяет, жива ли нода
	IsAlive() bool
	// Get получает значение с этой ноды
	Get(ctx context.Context, key string) (interface{}, bool, error)
	// Set сохраняет значение на эту ноду
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}

type ConsistentHash interface {
	// Add добавляет ноду в кольцо
	Add(node Node) error
	// Remove удаляет ноду из кольца
	Remove(nodeID string) error
	// GetNode возвращает ноду для ключа
	GetNode(key string) (Node, error)
	// GetNodes возвращает N нод для ключа (primary + replicas)
	GetNodes(key string, count int) ([]Node, error)
}

type CacheStats struct {
	Hits        int64
	Misses      int64
	Evictions   int64
	Size        int64
	AvgLoadTime time.Duration
	HitRate     float64
}

// Реализуйте:
// 1. ConsistentHashRing - консистентное хеширование с виртуальными нодами
// 2. ReplicatedCache - кэш с репликацией на N нод
// 3. SingleFlightCache - защита от cache stampede (дедупликация одновременных запросов)
// 4. TieredCache - многоуровневый кэш (L1: local memory, L2: distributed)
