package main

// Реализуйте интерфейс EventBus для системы событий.
// EventBus должен позволять подписываться на события по типу и рассылать события всем подписчикам.
// Важно: подписчики должны получать события асинхронно, но в порядке публикации.

type Event struct {
	Type string
	Data interface{}
}

type EventHandler interface {
	Handle(event Event) error
}

type EventBus interface {
	// Subscribe регистрирует обработчик для событий определенного типа
	Subscribe(eventType string, handler EventHandler) error
	// Unsubscribe удаляет обработчик для событий определенного типа
	Unsubscribe(eventType string, handler EventHandler) error
	// Publish публикует событие всем подписчикам
	Publish(event Event) error
	// Close завершает работу EventBus и ожидает обработки всех событий
	Close() error
}

// Реализуйте структуру SimpleEventBus, которая удовлетворяет интерфейсу EventBus
