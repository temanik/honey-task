package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для системы версионирования данных (Temporal Database / Time Travel).
// Система должна хранить всю историю изменений и позволять запрашивать данные на любой момент времени.
//
// Сложность:
// - Эффективное хранение версий (без дублирования неизменных данных)
// - Запросы "as of" - получение данных на определенный момент времени
// - Запросы "between" - получение всех изменений в диапазоне времени
// - Поддержка индексов с учетом времени
// - Garbage collection старых версий
// - Оптимистичные блокировки с версионированием

type Version struct {
	ID            int64
	RecordID      string
	Data          map[string]interface{}
	ValidFrom     time.Time
	ValidTo       *time.Time // nil означает текущую версию
	CreatedBy     string
	Operation     Operation
	TransactionID string
}

type Operation string

const (
	OperationInsert Operation = "insert"
	OperationUpdate Operation = "update"
	OperationDelete Operation = "delete"
)

type TimeQuery struct {
	AsOf    *time.Time // точка во времени
	From    *time.Time // начало диапазона
	To      *time.Time // конец диапазона
	Version *int64     // конкретная версия
}

type TemporalDB interface {
	// Insert вставляет новую запись
	Insert(ctx context.Context, id string, data map[string]interface{}) error
	// Update обновляет запись (создает новую версию)
	Update(ctx context.Context, id string, data map[string]interface{}) error
	// Delete удаляет запись (помечает как удаленную, но сохраняет историю)
	Delete(ctx context.Context, id string) error

	// Get получает текущую версию записи
	Get(ctx context.Context, id string) (map[string]interface{}, error)
	// GetAsOf получает запись на указанный момент времени
	GetAsOf(ctx context.Context, id string, timestamp time.Time) (map[string]interface{}, error)
	// GetVersion получает конкретную версию
	GetVersion(ctx context.Context, id string, version int64) (*Version, error)

	// Query выполняет запрос с учетом времени
	Query(ctx context.Context, query TimeQuery, filter map[string]interface{}) ([]map[string]interface{}, error)
}

type VersionHistory interface {
	// GetHistory получает всю историю изменений записи
	GetHistory(ctx context.Context, id string) ([]Version, error)
	// GetChanges получает изменения в диапазоне времени
	GetChanges(ctx context.Context, from, to time.Time) ([]Version, error)
	// GetChangesByUser получает изменения, сделанные пользователем
	GetChangesByUser(ctx context.Context, userID string, from, to time.Time) ([]Version, error)
	// Diff вычисляет разницу между двумя версиями
	Diff(ctx context.Context, id string, version1, version2 int64) (map[string]interface{}, error)
}

type TemporalIndex interface {
	// Index индексирует поле с учетом времени
	Index(ctx context.Context, field string) error
	// Search ищет записи по индексу на момент времени
	Search(ctx context.Context, field string, value interface{}, timestamp time.Time) ([]string, error)
	// RangeSearch поиск в диапазоне значений
	RangeSearch(ctx context.Context, field string, from, to interface{}, timestamp time.Time) ([]string, error)
}

type Snapshot interface {
	// CreateSnapshot создает снимок всей БД на момент времени
	CreateSnapshot(ctx context.Context, timestamp time.Time) (string, error)
	// RestoreSnapshot восстанавливает БД из снимка
	RestoreSnapshot(ctx context.Context, snapshotID string) error
	// ListSnapshots возвращает список всех снимков
	ListSnapshots(ctx context.Context) ([]SnapshotInfo, error)
	// DeleteSnapshot удаляет снимок
	DeleteSnapshot(ctx context.Context, snapshotID string) error
}

type SnapshotInfo struct {
	ID          string
	Timestamp   time.Time
	Size        int64
	RecordCount int64
}

type VersionGC interface {
	// SetRetentionPolicy устанавливает политику хранения
	SetRetentionPolicy(policy RetentionPolicy) error
	// Compact удаляет старые версии согласно политике
	Compact(ctx context.Context) (CompactionStats, error)
	// Vacuum освобождает место от удаленных версий
	Vacuum(ctx context.Context) error
}

type RetentionPolicy struct {
	// KeepVersions - сколько версий хранить для каждой записи
	KeepVersions int
	// KeepDuration - как долго хранить версии
	KeepDuration time.Duration
	// CompactAfter - когда начинать компактификацию
	CompactAfter time.Duration
	// KeepSnapshots - сколько снимков хранить
	KeepSnapshots int
}

type CompactionStats struct {
	VersionsRemoved int64
	SpaceFreed      int64
	Duration        time.Duration
}

type TemporalTransaction interface {
	// Begin начинает транзакцию с версионированием
	Begin(ctx context.Context) (string, error)
	// Insert вставка в рамках транзакции
	Insert(id string, data map[string]interface{}) error
	// Update обновление в рамках транзакции
	Update(id string, data map[string]interface{}, expectedVersion int64) error
	// Delete удаление в рамках транзакции
	Delete(id string) error
	// Commit фиксирует транзакцию
	Commit() error
	// Rollback откатывает транзакцию
	Rollback() error
}

type ConflictResolver interface {
	// DetectConflict проверяет конфликт версий
	DetectConflict(ctx context.Context, id string, expectedVersion, currentVersion int64) (bool, error)
	// ResolveConflict разрешает конфликт версий
	ResolveConflict(ctx context.Context, id string, version1, version2 *Version, strategy MergeStrategy) (*Version, error)
}

type MergeStrategy string

const (
	MergeStrategyLastWrite  MergeStrategy = "last_write_wins"
	MergeStrategyFirstWrite MergeStrategy = "first_write_wins"
	MergeStrategyManual     MergeStrategy = "manual"
	MergeStrategyMerge      MergeStrategy = "three_way_merge"
)

// Реализуйте:
// 1. TemporalDatabase - основная реализация temporal БД
// 2. VersionStore - эффективное хранилище версий (Copy-on-Write)
// 3. TemporalIndexManager - индексы с учетом времени
// 4. SnapshotManager - управление снимками
// 5. GarbageCollector - сборщик мусора для старых версий
// 6. OptimisticLock - оптимистичные блокировки с версиями
//
// Подсказки:
// - Для эффективного хранения используйте структурное разделение (structural sharing)
// - Индексы можно реализовать через B-Tree с timestamp в ключе
// - Для diff можно использовать алгоритм Myers или simple field comparison
