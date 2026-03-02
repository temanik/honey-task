package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://user:password@localhost/mydb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// создание таблицы
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	if err != nil {
		panic(err)
	}
	// создание индекса на два столбца
	_, err = db.Exec("CREATE INDEX users_name_age_idx ON users (name, age)")
	if err != nil {
		panic(err)
	}
	// вставка данных
	_, err = db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)", 1, "John", 30)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)", 2, "Mary", 25)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)", 3, "Peter", 40)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)", 4, "Jane", 35)
	if err != nil {
		panic(err)
	}
	// запрос с использованием покрывающего индекса
	rows, err := db.Query("SELECT id, name, age FROM users WHERE name = $1 AND age > $2", "Mary", 20)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}
}

/*
В этом примере мы создаем таблицу пользователей и индекс для столбцов name и age.
Затем мы вставляем четыре строки в таблицу. Затем мы выполняем запрос, который использует покрывающий индекс.
В этом запросе мы запрашиваем все строки, в которых name равен "Mary" и age больше 20.
Так как мы создали индекс на оба столбца, PostgreSQL может использовать его для выполнения запроса без обращения к самой таблице.
В результате мы получим две строки, где name равен "Mary" и age больше 20. Код будет выводить: id: 2, name: Mary, age: 25
*/
