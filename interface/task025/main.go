package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для графовой базы данных с поддержкой транзакций и индексов.
// База должна эффективно хранить и обрабатывать графы с возможностью выполнения сложных запросов.
//
// Сложность:
// - Эффективное хранение графа (adjacency list, adjacency matrix, edge list)
// - Поддержка ACID транзакций
// - Индексы для быстрого поиска вершин и ребер
// - Алгоритмы обхода графа (BFS, DFS, кратчайший путь)
// - Поддержка свойств на вершинах и ребрах
// - Оптимистичные и пессимистичные блокировки

type Vertex struct {
	ID         string
	Label      string
	Properties map[string]interface{}
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Edge struct {
	ID         string
	Label      string
	FromVertex string
	ToVertex   string
	Properties map[string]interface{}
	Weight     float64
	CreatedAt  time.Time
}

type Path struct {
	Vertices []Vertex
	Edges    []Edge
	Length   int
	Cost     float64
}

type GraphDB interface {
	// Transaction management
	BeginTx(ctx context.Context, opts TxOptions) (Transaction, error)

	// Vertex operations
	AddVertex(ctx context.Context, vertex Vertex) error
	GetVertex(ctx context.Context, id string) (*Vertex, error)
	UpdateVertex(ctx context.Context, id string, properties map[string]interface{}) error
	DeleteVertex(ctx context.Context, id string) error

	// Edge operations
	AddEdge(ctx context.Context, edge Edge) error
	GetEdge(ctx context.Context, id string) (*Edge, error)
	DeleteEdge(ctx context.Context, id string) error

	// Traversal
	GetNeighbors(ctx context.Context, vertexID string, direction Direction) ([]Vertex, error)
	GetEdges(ctx context.Context, vertexID string, direction Direction) ([]Edge, error)
}

type Direction string

const (
	DirectionIn   Direction = "in"
	DirectionOut  Direction = "out"
	DirectionBoth Direction = "both"
)

type TxOptions struct {
	ReadOnly  bool
	Isolation IsolationLevel
}

type IsolationLevel string

const (
	ReadUncommitted IsolationLevel = "read_uncommitted"
	ReadCommitted   IsolationLevel = "read_committed"
	RepeatableRead  IsolationLevel = "repeatable_read"
	Serializable    IsolationLevel = "serializable"
)

type Transaction interface {
	// Vertex operations within transaction
	AddVertex(vertex Vertex) error
	GetVertex(id string) (*Vertex, error)
	UpdateVertex(id string, properties map[string]interface{}) error
	DeleteVertex(id string) error

	// Edge operations within transaction
	AddEdge(edge Edge) error
	GetEdge(id string) (*Edge, error)
	DeleteEdge(id string) error

	// Transaction control
	Commit() error
	Rollback() error
}

type GraphQuery interface {
	// Pattern matching (упрощенный Cypher-like синтаксис)
	Match(pattern string) GraphQuery
	Where(condition string, params map[string]interface{}) GraphQuery
	Return(fields ...string) GraphQuery
	OrderBy(field string, ascending bool) GraphQuery
	Limit(limit int) GraphQuery
	Execute(ctx context.Context) ([]map[string]interface{}, error)
}

type GraphAlgorithms interface {
	// ShortestPath находит кратчайший путь между двумя вершинами
	ShortestPath(ctx context.Context, fromID, toID string) (*Path, error)
	// AllPaths находит все пути между двумя вершинами
	AllPaths(ctx context.Context, fromID, toID string, maxDepth int) ([]Path, error)
	// BFS обход в ширину
	BFS(ctx context.Context, startID string, visitor func(Vertex) bool) error
	// DFS обход в глубину
	DFS(ctx context.Context, startID string, visitor func(Vertex) bool) error
	// PageRank вычисляет PageRank для вершин
	PageRank(ctx context.Context, iterations int, dampingFactor float64) (map[string]float64, error)
	// DetectCycle определяет наличие циклов
	DetectCycle(ctx context.Context) (bool, []string, error)
	// ConnectedComponents находит связные компоненты
	ConnectedComponents(ctx context.Context) ([][]string, error)
}

type Index interface {
	// Add добавляет запись в индекс
	Add(key string, vertexID string) error
	// Remove удаляет запись из индекса
	Remove(key string, vertexID string) error
	// Search ищет вершины по ключу
	Search(key string) ([]string, error)
	// RangeSearch ищет вершины в диапазоне ключей
	RangeSearch(from, to string) ([]string, error)
}

// Реализуйте:
// 1. InMemoryGraphDB - графовая БД в памяти с поддержкой транзакций
// 2. GraphIndex - индекс для быстрого поиска (B-Tree или Hash)
// 3. GraphQueryEngine - движок для выполнения запросов
// 4. GraphAlgorithmsImpl - реализация алгоритмов на графах
// 5. MVCC (Multi-Version Concurrency Control) для изоляции транзакций
