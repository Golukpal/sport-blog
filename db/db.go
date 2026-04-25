package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr:= "user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}

	fmt.Println("db connected")

	DB = db 

	createTable()
}

func createTable(){
	userTable := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		password TEXT
	);`

	postTable := `
	CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		title TEXT,
		content TEXT,
		user_id INT
	);`

	DB.Exec(userTable)
	DB.Exec(postTable)
}

