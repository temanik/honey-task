# Honey-task

Коллекция практических задач по различным темам golang которые помогут закрепить знания.

## Доступные темы

### default block
- **[slices](slices/)** - Слайсы в Go (50 задач)
- **[maps](maps/)** - Maps в Go (50 задач)
- **[pointers](pointers/)** - Указатели в Go (50 задач)
- **[structs](structs/)** - Структуры в Go (50 задач)
- _interfaces_ - интерфейсы (soon)
- _defer|panic|recover_ - механизмы восстановления (soon)
- _errors_ - ошибки (soon)
- _strings_ - строки (its needed?)
- _func_ - функции (its needed?)

### concurrency block
- _channels_ - каналы и concurrency (soon)
- _goroutines_ - горутины (soon)
- _mutex_ - мьютексы (soon)
- _context_ - контекст (soon)
- _select_ - селект
- _sync_ - работа с пакетом sync
- _race cond_ - типичные проблемы и их решения

### Practice block
сборник ревью кода с собеседований (soon)
типовые задачи на написание n-программы с собеседований (soon)

```bash
# Запуск задачи
go run main.go <тема> <номера_задач>

# Смотри список доступных тем
go run main.go list

# Запуск одной задачи
go run main.go slices 3

# Запуск нескольких задач
go run main.go slices 1 5 10
```

## Почитать полезное:

- [Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Blog](https://go.dev/blog/)
