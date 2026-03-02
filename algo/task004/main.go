package main

import "fmt"

// Условие задачи
// Есть бинарное дерево категорий, где на разных уровнях дублируются названия категорий.
// Нужно вернуть набор наименований с количеством повторений

// Пример
// in

type Category struct {
	Name  string
	Left  *Category
	Right *Category
}

var root = &Category{
	Name: "Машины",
	Left: &Category{
		Name: "Легковые",
		Left: &Category{
			Name: "Российские",
			Left: &Category{
				Name: "Легковые",
			},
			Right: &Category{
				Name: "Красивые",
			},
		},
	},
	Right: &Category{
		Name: "Грузовые",
		Left: &Category{
			Name: "Большие",
			Left: &Category{
				Name: "Зеленые",
			},
			Right: &Category{
				Name: "Незеленые",
			},
		},
		Right: &Category{
			Name: "Красивые",
			Left: &Category{
				Name: "Красные",
			},
			Right: &Category{
				Name: "Некрасные",
				Left: &Category{
					Name: "Машины",
					Right: &Category{
						Name: "Машины",
					},
				},
			},
		},
	},
}

/*
// out

{
  'Машины': 3,
  'Легковые': 2,
  'Российские': 1,
  'Красивые': 2,
  'Грузовые': 1,
  'Большие': 1,
  'Зеленые': 1,
  'Незеленые': 1,
  'Красные': 1,
  'Некрасные': 1
}
*/

// stack
func fn2(cat *Category) map[string]int {
	res := make(map[string]int)

	if cat == nil {
		return res
	}

	stack := make([]*Category, 0)
	stack = append(stack, cat)

	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res[c.Name]++

		if c.Left != nil {
			stack = append(stack, c.Left)
		}
		if c.Right != nil {
			stack = append(stack, c.Right)
		}
	}

	return res
}

// recur
func fn(cat *Category) map[string]int {
	res := make(map[string]int)

	var dfs func(cat *Category)
	dfs = func(cat *Category) {
		if cat == nil {
			return
		}
		res[cat.Name]++
		dfs(cat.Left)
		dfs(cat.Right)
	}

	dfs(cat)
	return res
}

func main() {
	fmt.Println(fn(root))
	fmt.Println(fn2(root))
}
