package main

// Реализуйте интерфейс Observer (паттерн Observer) для системы уведомлений.
// Subject должен уведомлять подписчиков об изменениях своего состояния.

type Observer interface {
	// Update вызывается при изменении состояния Subject
	Update(event string, data interface{}) error
}

type Subject interface {
	// Attach добавляет наблюдателя
	Attach(observer Observer) error
	// Detach удаляет наблюдателя
	Detach(observer Observer) error
	// Notify уведомляет всех наблюдателей
	Notify(event string, data interface{}) error
}

// Реализуйте:
// 1. SimpleSubject - базовая реализация Subject
// 2. FilteredSubject - Subject, который уведомляет наблюдателей только о определенных событиях
//    Каждый Observer регистрируется с фильтром событий
// 3. AsyncSubject - Subject, который уведомляет наблюдателей асинхронно
//
// Пример использования:
// subject := NewSimpleSubject()
// observer1 := &LogObserver{}
// observer2 := &EmailObserver{}
// subject.Attach(observer1)
// subject.Attach(observer2)
// subject.Notify("user.created", User{ID: 1, Name: "John"})
