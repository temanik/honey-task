package main

import "fmt"

// Есть список отзывов на товар, который содержит текст отзыва и оценку товара в звездах(от 1 до 5)
// Необходимо сгруппировать отзывы с одинаковыми оценками

/* ### in
reviews = [
    {text: "Отлично!", rating: 5}, {text: "Хороший товар", rating: 4}, {text: "Ожидал большего", rating: 3},
    {text: "Не оправдал ожиданий", rating: 1}, {text: "Все как в описании", rating: 5}, {text: "Не понравилось", rating: 1}
]
### out
{ 5: ["Отлично!", "Все как в описании"], 4: ["Хороший товар"], 3: ["Ожидал большего"], 1: ["Не оправдал ожиданий", "Не понравилось"] }
*/

type Feed struct {
	text   string
	rating int
}

func feedback(feeds []Feed) map[int][]string {
	res := make(map[int][]string)

	for _, f := range feeds {
		rating := f.rating
		res[f.rating] = append(res[rating], f.text)
	}

	return res
}

func main() {
	feeds := []Feed{
		{"Отлично!", 5},
		{"Хороший товар", 4},
		{"Ожидал большего", 3},
		{"Не оправдал ожиданий", 1},
		{"Все как в описании", 5},
		{"Не понравилось", 1},
	}

	fmt.Println(feedback(feeds))
}
