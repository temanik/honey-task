package main

/*
Есть матрица NxN, состоящая из 0 и 1, и отражающая расположения кораблей на поле для морского боя. - Кораблей может быть любое количество.
Размер кораблей - от 1х1 до 1хN.
Корабли никак не соприкасаются друг с другом.
Корабли располагаются горизонтально и вертикально.
Необходимо подсчитать количество кораблей.
Пример: [1, 1, 0, 0, 1, 0], [0, 0, 0, 0, 1, 0], [1, 0, 1, 0, 1, 0], [0, 0, 0, 0, 0, 0], [1, 0, 1, 1, 1, 1], [0, 0, 0, 0, 0, 0],
package main

import "fmt"

func solve(field [][]int) int {
  return 0
}

func main() {
  var field = [][]int{
    {1, 1, 0, 0, 1, 0},
    {1, 0, 0, 0, 1, 0},
    {1, 0, 1, 0, 1, 0},
    {0, 0, 0, 0, 0, 0},
    {1, 0, 1, 1, 1, 1},
    {0, 0, 0, 0, 0, 0},
  }

  fmt.Println(solve(field))
}
*/

import "fmt"

func dfs(x, y int, field [][]int) {
	if y < 0 || y >= len(field) || x < 0 || x >= len(field[y]) {
		return
	}

	if field[y][x] == 0 {
		return
	}

	field[y][x] = 0

	dfs(x+1, y, field)
	dfs(x, y+1, field)
	dfs(x-1, y, field)
	dfs(x, y-1, field)
}

func solve(field [][]int) int {
	shipCount := 0

	for y := range field {
		for x, v := range field[y] {
			if v == 1 {
				shipCount++
				dfs(x, y, field)
			}
		}
	}

	return shipCount
}

func findShips(fields [][]int) int {
	shipCount := 0

	for y := range fields {
		for x, v := range fields[y] {
			if (x-1 < 0 || fields[y][x-1] == 0) &&
				(y-1 < 0 || fields[y-1][x] == 0) &&
				v == 1 {
				shipCount++
			}
		}
	}

	return shipCount
}

func main() {
	var field = [][]int{
		{1, 1, 0, 0, 1, 0},
		{1, 0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0},
	}

	var f = [][]int{
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1, 0},
		{0, 0, 1, 1, 0, 0},
		{1, 0, 0, 1, 1, 1},
		{1, 0, 0, 0, 0, 1},
	}

	fmt.Println(solve(field))
	fmt.Println(findShips(f))
}
