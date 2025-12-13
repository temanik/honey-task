package main

import "time"

// Реализуйте интерфейс для работы с различными хранилищами (Storage).
// Storage должен поддерживать операции чтения/записи и быть расширяемым.
// Создайте реализации для файловой системы, памяти и композитного хранилища.

type Metadata struct {
	Size         int64
	LastModified time.Time
	ContentType  string
	Custom       map[string]string
}

type Storage interface {
	// Put сохраняет данные по ключу
	Put(key string, data []byte, metadata Metadata) error
	// Get получает данные по ключу
	Get(key string) ([]byte, Metadata, error)
	// Delete удаляет данные по ключу
	Delete(key string) error
	// Exists проверяет существование ключа
	Exists(key string) (bool, error)
	// List возвращает список ключей с заданным префиксом
	List(prefix string) ([]string, error)
}

// Реализуйте:
// 1. MemoryStorage - хранилище в памяти
// 2. FileStorage - хранилище на файловой системе
// 3. TieredStorage - многоуровневое хранилище (например, память + диск)
//    При чтении сначала проверяет быструю память, потом медленный диск
//    При записи пишет в оба хранилища
