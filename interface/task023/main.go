package main

import (
	"context"
	"time"
)

// Реализуйте интерфейс для транзакционной системы с поддержкой SAGA паттерна.
// Система должна обеспечивать консистентность в распределенных транзакциях.
//
// Сложность:
// - Реализация SAGA паттерна (choreography или orchestration)
// - Компенсирующие транзакции при откате
// - Идемпотентность операций
// - Персистентность состояния транзакции
// - Поддержка timeout и deadlines

type TransactionStatus string

const (
	StatusPending      TransactionStatus = "pending"
	StatusInProgress   TransactionStatus = "in_progress"
	StatusCompleted    TransactionStatus = "completed"
	StatusFailed       TransactionStatus = "failed"
	StatusCompensating TransactionStatus = "compensating"
	StatusCompensated  TransactionStatus = "compensated"
)

type Step struct {
	Name       string
	Action     func(ctx context.Context, data interface{}) (interface{}, error)
	Compensate func(ctx context.Context, data interface{}) error
	Timeout    time.Duration
}

type Transaction struct {
	ID          string
	Steps       []Step
	Status      TransactionStatus
	Data        map[string]interface{}
	StartedAt   time.Time
	CompletedAt *time.Time
	Error       error
}

type SagaOrchestrator interface {
	// Execute выполняет транзакцию (все шаги последовательно)
	Execute(ctx context.Context, tx *Transaction) error
	// GetStatus возвращает текущий статус транзакции
	GetStatus(ctx context.Context, txID string) (TransactionStatus, error)
	// Compensate откатывает транзакцию (выполняет компенсации в обратном порядке)
	Compensate(ctx context.Context, txID string) error
	// Resume возобновляет прерванную транзакцию
	Resume(ctx context.Context, txID string) error
}

type TransactionLog interface {
	// LogStepStarted логирует начало шага
	LogStepStarted(ctx context.Context, txID string, stepName string, data interface{}) error
	// LogStepCompleted логирует завершение шага
	LogStepCompleted(ctx context.Context, txID string, stepName string, result interface{}) error
	// LogStepFailed логирует ошибку шага
	LogStepFailed(ctx context.Context, txID string, stepName string, err error) error
	// LogCompensationStarted логирует начало компенсации
	LogCompensationStarted(ctx context.Context, txID string, stepName string) error
	// GetTransactionState восстанавливает состояние транзакции
	GetTransactionState(ctx context.Context, txID string) (*Transaction, error)
}

type DistributedTransaction interface {
	// Begin начинает распределенную транзакцию
	Begin(ctx context.Context) (string, error)
	// AddStep добавляет шаг в транзакцию
	AddStep(ctx context.Context, txID string, step Step) error
	// Commit фиксирует транзакцию
	Commit(ctx context.Context, txID string) error
	// Rollback откатывает транзакцию
	Rollback(ctx context.Context, txID string) error
	// GetHistory возвращает историю выполнения транзакции
	GetHistory(ctx context.Context, txID string) ([]TransactionEvent, error)
}

type TransactionEvent struct {
	TxID      string
	StepName  string
	EventType string
	Data      interface{}
	Timestamp time.Time
}

// Реализуйте:
// 1. SagaOrchestrator - оркестратор SAGA транзакций
// 2. InMemoryTransactionLog - лог транзакций в памяти
// 3. CompensatingTransactionManager - менеджер компенсирующих транзакций
//
// Пример использования:
// tx := &Transaction{
//     ID: "order-123",
//     Steps: []Step{
//         {Name: "reserve-inventory", Action: reserveInventory, Compensate: releaseInventory},
//         {Name: "charge-payment", Action: chargePayment, Compensate: refundPayment},
//         {Name: "create-shipment", Action: createShipment, Compensate: cancelShipment},
//     },
// }
// orchestrator.Execute(ctx, tx)
