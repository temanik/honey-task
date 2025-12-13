package main

// Реализуйте интерфейс для построения query builder.
// QueryBuilder должен позволять строить SQL запросы через fluent interface.

type QueryBuilder interface {
	// Select указывает поля для выборки
	Select(fields ...string) QueryBuilder
	// From указывает таблицу
	From(table string) QueryBuilder
	// Where добавляет условие WHERE
	Where(condition string, args ...interface{}) QueryBuilder
	// Join добавляет JOIN
	Join(table string, condition string) QueryBuilder
	// OrderBy добавляет сортировку
	OrderBy(field string, desc bool) QueryBuilder
	// Limit устанавливает лимит
	Limit(limit int) QueryBuilder
	// Offset устанавливает offset
	Offset(offset int) QueryBuilder
	// Build строит финальный SQL запрос
	Build() (query string, args []interface{}, err error)
}

type InsertBuilder interface {
	// Into указывает таблицу для вставки
	Into(table string) InsertBuilder
	// Values добавляет значения для вставки
	Values(values map[string]interface{}) InsertBuilder
	// Build строит финальный SQL запрос
	Build() (query string, args []interface{}, err error)
}

type UpdateBuilder interface {
	// Table указывает таблицу для обновления
	Table(table string) UpdateBuilder
	// Set устанавливает значение поля
	Set(field string, value interface{}) UpdateBuilder
	// Where добавляет условие
	Where(condition string, args ...interface{}) UpdateBuilder
	// Build строит финальный SQL запрос
	Build() (query string, args []interface{}, err error)
}

// Реализуйте SQLBuilder, который поддерживает все три типа запросов
