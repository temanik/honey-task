package main

// Реализуйте интерфейс Validator для валидации данных.
// Создайте композитный валидатор, который может объединять несколько валидаторов.
// Также реализуйте несколько базовых валидаторов.

type ValidationError struct {
	Field   string
	Message string
}

type Validator interface {
	Validate(data interface{}) []ValidationError
}

type CompositeValidator interface {
	Validator
	// Add добавляет валидатор в композицию
	Add(validator Validator)
	// ValidateField валидирует конкретное поле
	ValidateField(field string, value interface{}) []ValidationError
}

// Реализуйте следующие валидаторы:
// 1. RequiredValidator - проверяет, что значение не пустое
// 2. LengthValidator - проверяет длину строки (min, max)
// 3. EmailValidator - проверяет формат email
// 4. ChainValidator - последовательно применяет несколько валидаторов
// 5. ParallelValidator - применяет валидаторы параллельно и собирает все ошибки
