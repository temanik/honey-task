package main

import "time"

// Реализуйте интерфейс Scheduler для планирования и выполнения задач.
// Scheduler должен поддерживать одноразовые задачи, повторяющиеся задачи и cron-like расписания.

type ScheduledTask interface {
	ID() string
	Run() error
	Cancel()
}

type Schedule interface {
	// Next возвращает время следующего выполнения после указанного времени
	Next(after time.Time) time.Time
}

type Scheduler interface {
	// ScheduleOnce планирует одноразовое выполнение задачи
	ScheduleOnce(delay time.Duration, task func() error) (ScheduledTask, error)
	// ScheduleRepeat планирует повторяющееся выполнение с фиксированным интервалом
	ScheduleRepeat(interval time.Duration, task func() error) (ScheduledTask, error)
	// ScheduleWithSchedule планирует выполнение согласно расписанию
	ScheduleWithSchedule(schedule Schedule, task func() error) (ScheduledTask, error)
	// Cancel отменяет задачу по ID
	Cancel(taskID string) error
	// Start запускает планировщик
	Start() error
	// Stop останавливает планировщик
	Stop() error
}

// Реализуйте:
// 1. SimpleScheduler
// 2. PeriodicSchedule - расписание с фиксированным периодом
// 3. CronSchedule - cron-like расписание (упрощенный вариант)
