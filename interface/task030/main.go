package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для системы распределенного consensus (Raft/Paxos-подобный алгоритм).
// Система должна обеспечивать согласованное состояние кластера при наличии сбоев.
//
// Сложность:
// - Leader election (выбор лидера при сбое)
// - Log replication (репликация команд на все узлы)
// - Safety гарантии (linearizability)
// - Membership changes (добавление/удаление узлов)
// - Log compaction (снимки состояния)
// - Network partitions handling (split-brain prevention)

type NodeState string

const (
	StateFollower  NodeState = "follower"
	StateCandidate NodeState = "candidate"
	StateLeader    NodeState = "leader"
)

type LogEntry struct {
	Index     int64
	Term      int64
	Command   interface{}
	Timestamp time.Time
}

type Snapshot struct {
	LastIncludedIndex int64
	LastIncludedTerm  int64
	Data              []byte
	Timestamp         time.Time
}

type ConsensusNode interface {
	// Start запускает узел
	Start(ctx context.Context) error
	// Stop останавливает узел
	Stop(ctx context.Context) error

	// Propose предлагает новую команду (доступно только для лидера)
	Propose(ctx context.Context, command interface{}) error
	// GetState возвращает текущее состояние узла
	GetState() NodeInfo
	// IsLeader проверяет, является ли узел лидером
	IsLeader() bool

	// AddNode добавляет новый узел в кластер
	AddNode(ctx context.Context, nodeID string, address string) error
	// RemoveNode удаляет узел из кластера
	RemoveNode(ctx context.Context, nodeID string) error
}

type NodeInfo struct {
	ID           string
	State        NodeState
	CurrentTerm  int64
	VotedFor     string
	Leader       string
	LastLogIndex int64
	LastLogTerm  int64
	CommitIndex  int64
	LastApplied  int64
}

type RaftRPC interface {
	// RequestVote обрабатывает запрос на голосование
	RequestVote(ctx context.Context, req VoteRequest) (*VoteResponse, error)
	// AppendEntries обрабатывает добавление записей в лог
	AppendEntries(ctx context.Context, req AppendEntriesRequest) (*AppendEntriesResponse, error)
	// InstallSnapshot обрабатывает установку снимка
	InstallSnapshot(ctx context.Context, req SnapshotRequest) (*SnapshotResponse, error)
}

type VoteRequest struct {
	Term         int64
	CandidateID  string
	LastLogIndex int64
	LastLogTerm  int64
}

type VoteResponse struct {
	Term        int64
	VoteGranted bool
}

type AppendEntriesRequest struct {
	Term         int64
	LeaderID     string
	PrevLogIndex int64
	PrevLogTerm  int64
	Entries      []LogEntry
	LeaderCommit int64
}

type AppendEntriesResponse struct {
	Term    int64
	Success bool
	// Optimization: для быстрого нахождения места расхождения логов
	ConflictTerm  int64
	ConflictIndex int64
}

type SnapshotRequest struct {
	Term              int64
	LeaderID          string
	LastIncludedIndex int64
	LastIncludedTerm  int64
	Offset            int64
	Data              []byte
	Done              bool
}

type SnapshotResponse struct {
	Term int64
}

type StateMachine interface {
	// Apply применяет команду к state machine
	Apply(entry LogEntry) error
	// GetState возвращает текущее состояние
	GetState() interface{}
	// CreateSnapshot создает снимок состояния
	CreateSnapshot() ([]byte, error)
	// RestoreSnapshot восстанавливает состояние из снимка
	RestoreSnapshot(snapshot []byte) error
}

type LogStore interface {
	// Append добавляет записи в лог
	Append(entries []LogEntry) error
	// Get получает запись по индексу
	Get(index int64) (*LogEntry, error)
	// GetRange получает записи в диапазоне [from, to]
	GetRange(from, to int64) ([]LogEntry, error)
	// GetLastIndex возвращает индекс последней записи
	GetLastIndex() int64
	// GetLastTerm возвращает term последней записи
	GetLastTerm() int64
	// DeleteRange удаляет записи начиная с from
	DeleteRange(from int64) error
	// Compact компактифицирует лог до индекса (удаляет старые записи)
	Compact(index int64) error
}

type SnapshotStore interface {
	// Save сохраняет снимок
	Save(snapshot Snapshot) error
	// Load загружает последний снимок
	Load() (*Snapshot, error)
	// List возвращает список доступных снимков
	List() ([]Snapshot, error)
	// Delete удаляет снимок
	Delete(lastIncludedIndex int64) error
}

type ClusterMembership interface {
	// GetMembers возвращает текущий состав кластера
	GetMembers() []Member
	// AddMember добавляет члена кластера
	AddMember(member Member) error
	// RemoveMember удаляет члена кластера
	RemoveMember(nodeID string) error
	// GetMember получает информацию о члене
	GetMember(nodeID string) (*Member, error)
	// GetQuorum возвращает размер кворума
	GetQuorum() int
}

type Member struct {
	ID       string
	Address  string
	Status   MemberStatus
	Role     MemberRole
	JoinedAt time.Time
}

type MemberStatus string

const (
	MemberStatusAlive   MemberStatus = "alive"
	MemberStatusSuspect MemberStatus = "suspect"
	MemberStatusDead    MemberStatus = "dead"
)

type MemberRole string

const (
	MemberRoleVoter    MemberRole = "voter"
	MemberRoleNonVoter MemberRole = "non_voter"
	MemberRoleLearner  MemberRole = "learner"
)

type LeaderElection interface {
	// StartElection начинает процесс выборов
	StartElection(ctx context.Context) error
	// BecomeLeader переводит узел в состояние лидера
	BecomeLeader(ctx context.Context) error
	// BecomeFollower переводит узел в состояние последователя
	BecomeFollower(ctx context.Context, term int64, leader string) error
	// ResetElectionTimeout сбрасывает таймаут выборов
	ResetElectionTimeout()
}

type Replicator interface {
	// ReplicateLog реплицирует лог на последователей
	ReplicateLog(ctx context.Context) error
	// SendHeartbeat отправляет heartbeat последователям
	SendHeartbeat(ctx context.Context) error
	// UpdateCommitIndex обновляет commit index на основе репликации
	UpdateCommitIndex(ctx context.Context) error
}

type ConsensusConfig struct {
	NodeID              string
	ElectionTimeoutMin  time.Duration
	ElectionTimeoutMax  time.Duration
	HeartbeatInterval   time.Duration
	SnapshotInterval    int64 // количество записей в логе до создания снимка
	MaxEntriesPerAppend int
	MaxSnapshotSize     int64
}

type ClusterObserver interface {
	// OnLeaderChange вызывается при смене лидера
	OnLeaderChange(oldLeader, newLeader string)
	// OnMembershipChange вызывается при изменении состава кластера
	OnMembershipChange(change MembershipChange)
	// OnCommit вызывается при commit записи
	OnCommit(entry LogEntry)
}

type MembershipChange struct {
	Type       ChangeType
	Member     Member
	OldMembers []Member
	NewMembers []Member
}

type ChangeType string

const (
	ChangeTypeAdd    ChangeType = "add"
	ChangeTypeRemove ChangeType = "remove"
	ChangeTypeUpdate ChangeType = "update"
)

type ConsensusMetrics struct {
	CurrentTerm        int64
	State              NodeState
	Leader             string
	CommitIndex        int64
	LastApplied        int64
	LogSize            int64
	SnapshotSize       int64
	Elections          int64
	HeartbeatsSent     int64
	HeartbeatsReceived int64
	RPCErrors          int64
}

// Реализуйте:
// 1. RaftNode - узел Raft кластера
// 2. InMemoryLogStore - хранилище лога в памяти
// 3. SnapshotManager - управление снимками состояния
// 4. MembershipManager - управление составом кластера
// 5. ReplicationManager - репликация лога
// 6. ElectionController - контроллер выборов
// 7. StateMachineExecutor - применение команд к state machine
//
// Подсказки:
// - Election timeout должен быть рандомизированным для избежания split votes
// - Heartbeat interval должен быть значительно меньше election timeout
// - Используйте векторные часы или monotonic timestamps для упорядочивания
// - Для network partitions реализуйте pre-vote фазу (raft extension)
// - Log compaction критичен для долгоживущих систем
//
// Пример использования:
// config := ConsensusConfig{
//     NodeID: "node-1",
//     ElectionTimeoutMin: 150 * time.Millisecond,
//     ElectionTimeoutMax: 300 * time.Millisecond,
//     HeartbeatInterval: 50 * time.Millisecond,
// }
// node := NewRaftNode(config, stateMachine)
// node.Start(ctx)
//
// // На лидере:
// if node.IsLeader() {
//     node.Propose(ctx, command)
// }
