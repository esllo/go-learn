package example

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func insertDB(db *sql.DB) {
	result, errInsert := db.Exec("INSERT INTO golang VALUES (?, ?, ?)", nil, "Hong Gildong", "description")
	if errInsert != nil {
		log.Fatal(errInsert)
	}

	count, errResult := result.RowsAffected()
	if errResult != nil {
		fmt.Printf("%d row inserted\n", count)
	}
}

func selectDB(db *sql.DB) {
	rows, errSelect := db.Query("SELECT id, name, etc FROM golang")
	if errSelect != nil {
		log.Fatal(errSelect)
	}

	var id int
	var name, etc string

	for rows.Next() {
		errNext := rows.Scan(&id, &name, &etc)
		if errNext != nil {
			log.Fatal(errNext)
		}

		fmt.Printf("[%d] %s : %s\n", id, name, etc)
	}
}

func ExampleDB() {
	DBURL := os.Getenv("DBURL")
	db, errConnect := sql.Open("mysql", DBURL)
	if errConnect != nil {
		panic(errConnect)
	}
	defer db.Close()

	insertDB(db)
	selectDB(db)
}
