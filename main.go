package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/honeynil/honey-task/slices"
)

// Topic представляет тему с задачами
type Topic struct {
	Name        string
	Tasks       map[int]func()
	Count       int
	Description string
}

var topics = map[string]Topic{
	"slices": {
		Name:        "slices",
		Tasks:       slices.GetTasks(),
		Count:       slices.Count(),
		Description: "Слайсы в Go - 50 задач",
	},
	// Здесь можно добавлять новые темы:
	// "maps": {
	//     Name:        "maps",
	//     Tasks:       maps.GetTasks(),
	//     Count:       maps.Count(),
	//     Description: "Maps в Go - задачи",
	// },
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	topicName := os.Args[1]

	// Специальная команда для списка тем
	if topicName == "list" {
		listTopics()
		return
	}

	topic, ok := topics[topicName]
	if !ok {
		fmt.Printf("Неизвестная тема: %s\n\n", topicName)
		printHelp()
		return
	}

	// Если нет аргументов после темы, показываем справку по теме
	if len(os.Args) < 3 {
		printTopicHelp(topic)
		return
	}

	// Обработка команды "all"
	if os.Args[2] == "all" {
		runAllTasks(topic)
		return
	}

	// Запуск конкретных задач
	runTasks(topic, os.Args[2:])
}

func printHelp() {
	fmt.Println("🎯 Go Practice Tasks - Практические задачи по Go")
	fmt.Println("\nИспользование:")
	fmt.Println("  go run main.go <тема> <задачи>")
	fmt.Println("\nКоманды:")
	fmt.Println("  go run main.go list                    - список доступных тем")
	fmt.Println("  go run main.go <тема>                  - справка по теме")
	fmt.Println("  go run main.go <тема> <номер>          - запустить задачу")
	fmt.Println("  go run main.go <тема> <номера>         - запустить несколько задач")
	fmt.Println("  go run main.go <тема> all              - запустить все задачи")
	fmt.Println("\nПримеры:")
	fmt.Println("  go run main.go slices 1                - запустить задачу 1 по слайсам")
	fmt.Println("  go run main.go slices 1 5 10           - запустить задачи 1, 5 и 10")
	fmt.Println("  go run main.go slices all              - все задачи по слайсам")
	fmt.Println("\nДоступные темы:")
	for name, topic := range topics {
		fmt.Printf("  %-10s - %s (%d задач)\n", name, topic.Description, topic.Count)
	}
}

func listTopics() {
	fmt.Println("📚 Доступные темы:\n")
	for name, topic := range topics {
		fmt.Printf("  %-10s - %s (%d задач)\n", name, topic.Description, topic.Count)
	}
	fmt.Println("\nИспользуйте: go run main.go <тема> для подробностей")
}

func printTopicHelp(topic Topic) {
	fmt.Printf("📖 %s\n\n", topic.Description)
	fmt.Printf("Доступно задач: %d\n\n", topic.Count)
	fmt.Println("Использование:")
	fmt.Printf("  go run main.go %s <номер>          - запустить одну задачу\n", topic.Name)
	fmt.Printf("  go run main.go %s 1 5 10           - запустить несколько задач\n", topic.Name)
	fmt.Printf("  go run main.go %s all              - запустить все задачи\n", topic.Name)
	fmt.Println("\nПримеры:")
	fmt.Printf("  go run main.go %s 1\n", topic.Name)
	fmt.Printf("  go run main.go %s 1 5 10\n", topic.Name)
}

func runAllTasks(topic Topic) {
	fmt.Printf("\n🚀 Запуск всех задач: %s\n", topic.Description)
	for i := 1; i <= topic.Count; i++ {
		fmt.Printf("\n" + strings.Repeat("=", 50) + "\n")
		fmt.Printf("ЗАДАЧА %d\n", i)
		fmt.Printf(strings.Repeat("=", 50) + "\n\n")
		if task, ok := topic.Tasks[i]; ok {
			task()
		}
	}
}

func runTasks(topic Topic, args []string) {
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil || num < 1 || num > topic.Count {
			fmt.Printf("❌ Неверный номер задачи: %s (доступны 1-%d)\n", arg, topic.Count)
			continue
		}

		fmt.Printf("\n" + strings.Repeat("=", 50) + "\n")
		fmt.Printf("ЗАДАЧА %d\n", num)
		fmt.Printf(strings.Repeat("=", 50) + "\n\n")

		if task, ok := topic.Tasks[num]; ok {
			task()
		}
	}
}
