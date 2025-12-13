package main

// Реализуйте интерфейс Encoder для сериализации данных.
// Создайте реализации для JSON и XML форматов.
// Также реализуйте MultiEncoder, который может кодировать в несколько форматов одновременно.

type Encoder interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
	ContentType() string
}

type MultiEncoder interface {
	// AddEncoder добавляет энкодер в список
	AddEncoder(name string, encoder Encoder)
	// Encode кодирует данные всеми зарегистрированными энкодерами
	Encode(data interface{}) (map[string][]byte, error)
	// DecodeWith декодирует данные используя указанный энкодер
	DecodeWith(name string, data []byte, v interface{}) error
}

// Реализуйте:
// 1. JSONEncoder
// 2. XMLEncoder
// 3. CompositeEncoder, который реализует MultiEncoder
