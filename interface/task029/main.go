package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для системы распределенных вычислений (Map-Reduce framework).
// Система должна распределять вычисления по узлам кластера с поддержкой fault tolerance.
//
// Сложность:
// - Распределение задач по воркерам (partitioning)
// - Fault tolerance - перезапуск упавших задач
// - Shuffling данных между Map и Reduce фазами
// - Combiner для оптимизации (локальная агрегация перед shuffle)
// - Speculative execution (запуск дублирующих задач для медленных воркеров)
// - Locality awareness (запуск задач рядом с данными)

type MapFunc func(key, value interface{}) ([]KeyValue, error)
type ReduceFunc func(key interface{}, values []interface{}) (interface{}, error)
type CombineFunc func(key interface{}, values []interface{}) ([]interface{}, error)

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

type Job struct {
	ID          string
	Name        string
	MapFunc     MapFunc
	ReduceFunc  ReduceFunc
	CombineFunc CombineFunc // опциональный
	Input       []InputSplit
	NumReducers int
	Config      JobConfig
}

type JobConfig struct {
	MaxAttempts          int
	TaskTimeout          time.Duration
	SpeculativeExecution bool
	LocalityPreference   bool
	Compression          bool
}

type InputSplit struct {
	ID       string
	Location string // где находятся данные
	Offset   int64
	Length   int64
	Data     []KeyValue
}

type TaskType string

const (
	TaskTypeMap    TaskType = "map"
	TaskTypeReduce TaskType = "reduce"
)

type Task struct {
	ID          string
	JobID       string
	Type        TaskType
	SplitID     string // для map задач
	Partition   int    // для reduce задач
	Attempt     int
	Status      TaskStatus
	WorkerID    string
	StartedAt   *time.Time
	CompletedAt *time.Time
	Error       error
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
	TaskStatusKilled    TaskStatus = "killed"
)

type MapReduceFramework interface {
	// SubmitJob отправляет задание на выполнение
	SubmitJob(ctx context.Context, job Job) (string, error)
	// GetJobStatus получает статус задания
	GetJobStatus(ctx context.Context, jobID string) (*JobStatus, error)
	// CancelJob отменяет задание
	CancelJob(ctx context.Context, jobID string) error
	// WaitForCompletion ждет завершения задания
	WaitForCompletion(ctx context.Context, jobID string) (*JobResult, error)
}

type JobStatus struct {
	JobID          string
	Status         TaskStatus
	Progress       float64
	MapProgress    float64
	ReduceProgress float64
	StartedAt      time.Time
	CompletedAt    *time.Time
	TotalTasks     int
	CompletedTasks int
	FailedTasks    int
	RunningTasks   int
}

type JobResult struct {
	JobID     string
	Success   bool
	Output    []KeyValue
	Duration  time.Duration
	TaskStats TaskStatistics
	Error     error
}

type TaskStatistics struct {
	TotalMapTasks     int
	TotalReduceTasks  int
	FailedMapTasks    int
	FailedReduceTasks int
	AvgMapDuration    time.Duration
	AvgReduceDuration time.Duration
	DataShuffled      int64
	SpeculativeTasks  int
}

type Master interface {
	// Schedule распределяет задачи по воркерам
	Schedule(ctx context.Context, tasks []Task) error
	// HandleTaskCompletion обрабатывает завершение задачи
	HandleTaskCompletion(ctx context.Context, task Task, output []KeyValue) error
	// HandleTaskFailure обрабатывает ошибку задачи
	HandleTaskFailure(ctx context.Context, task Task, err error) error
	// MonitorWorkers мониторит состояние воркеров
	MonitorWorkers(ctx context.Context) error
	// Rebalance перебалансирует задачи при изменении кластера
	Rebalance(ctx context.Context) error
}

type Worker interface {
	// ID возвращает идентификатор воркера
	ID() string
	// ExecuteMap выполняет map задачу
	ExecuteMap(ctx context.Context, task Task, mapFunc MapFunc) ([]KeyValue, error)
	// ExecuteReduce выполняет reduce задачу
	ExecuteReduce(ctx context.Context, task Task, reduceFunc ReduceFunc, input []KeyValue) (interface{}, error)
	// GetStatus возвращает статус воркера
	GetStatus() WorkerStatus
	// Heartbeat отправляет heartbeat мастеру
	Heartbeat(ctx context.Context) error
}

type WorkerStatus struct {
	ID             string
	Location       string
	IsAlive        bool
	CurrentTasks   []string
	CPUUsage       float64
	MemoryUsage    int64
	LastHeartbeat  time.Time
	TasksCompleted int64
	TasksFailed    int64
}

type Partitioner interface {
	// Partition определяет, в какую partition отправить ключ
	Partition(key interface{}, numPartitions int) int
}

type Shuffler interface {
	// Shuffle перемещает данные от map к reduce воркерам
	Shuffle(ctx context.Context, mapOutput []KeyValue, numReducers int) (map[int][]KeyValue, error)
	// Fetch получает данные для reduce задачи
	Fetch(ctx context.Context, jobID string, partition int) ([]KeyValue, error)
	// Store сохраняет промежуточные результаты
	Store(ctx context.Context, jobID string, mapTaskID string, output []KeyValue) error
}

type Scheduler interface {
	// AssignTask назначает задачу воркеру
	AssignTask(ctx context.Context, task Task) (string, error)
	// SelectWorker выбирает оптимального воркера для задачи
	SelectWorker(ctx context.Context, task Task, workers []WorkerStatus) (*WorkerStatus, error)
	// HandleStragglers обрабатывает медленные задачи
	HandleStragglers(ctx context.Context, job Job) error
}

type FaultTolerance interface {
	// DetectFailure определяет упавшие задачи/воркеры
	DetectFailure(ctx context.Context) ([]Task, []string, error)
	// RecoverTask восстанавливает упавшую задачу
	RecoverTask(ctx context.Context, task Task) error
	// Checkpoint создает контрольную точку
	Checkpoint(ctx context.Context, jobID string, state JobState) error
	// Restore восстанавливает состояние из контрольной точки
	Restore(ctx context.Context, jobID string) (*JobState, error)
}

type JobState struct {
	JobID            string
	CompletedTasks   map[string][]KeyValue
	PendingTasks     []Task
	IntermediateData map[string][]KeyValue
	Timestamp        time.Time
}

type LocalityScheduler interface {
	Scheduler
	// GetDataLocality возвращает информацию о расположении данных
	GetDataLocality(ctx context.Context, splitID string) ([]string, error)
	// CalculateLocality вычисляет степень локальности
	CalculateLocality(ctx context.Context, task Task, worker WorkerStatus) float64
}

// Реализуйте:
// 1. MapReduceMaster - координатор заданий
// 2. MapReduceWorker - воркер для выполнения задач
// 3. HashPartitioner - распределение по хешу
// 4. NetworkShuffler - перемещение данных по сети
// 5. LocalityAwareScheduler - планировщик с учетом расположения данных
// 6. FaultTolerantExecutor - выполнение с поддержкой fault tolerance
// 7. SpeculativeScheduler - спекулятивное выполнение медленных задач
//
// Пример использования:
// job := Job{
//     MapFunc: func(key, value interface{}) ([]KeyValue, error) {
//         // word count: emit each word
//         words := strings.Split(value.(string), " ")
//         result := make([]KeyValue, len(words))
//         for i, word := range words {
//             result[i] = KeyValue{Key: word, Value: 1}
//         }
//         return result, nil
//     },
//     ReduceFunc: func(key interface{}, values []interface{}) (interface{}, error) {
//         // sum all counts
//         sum := 0
//         for _, v := range values {
//             sum += v.(int)
//         }
//         return sum, nil
//     },
//     NumReducers: 10,
// }
// framework.SubmitJob(ctx, job)
