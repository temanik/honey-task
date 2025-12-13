package main

// Реализуйте интерфейс Repository для работы с пользователями.
// Нужно создать две реализации: InMemoryRepository и CachedRepository.
// CachedRepository должен кэшировать результаты операций GetByID.

type User struct {
	ID    int
	Name  string
	Email string
}

type Repository interface {
	Save(user User) error
	GetByID(id int) (User, error)
	GetAll() ([]User, error)
	Delete(id int) error
}

// Реализуйте:
// 1. InMemoryRepository - простое хранилище в памяти
// 2. CachedRepository - обертка над любым Repository с кэшированием GetByID
//
// func NewInMemoryRepository() *InMemoryRepository
// func NewCachedRepository(repo Repository, ttl time.Duration) *CachedRepository
