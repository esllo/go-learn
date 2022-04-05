package example

import (
	"database/sql"
	"log"
	"os"
)

func preparedStatement(db *sql.DB) {
	statement, errPrepare := db.Prepare("UPDATE golang SET etc=? WHERE etc=? LIMIT 1")
	if errPrepare != nil {
		log.Fatal(errPrepare)
	}

	_, errExec := statement.Exec("description", "modified")
	if errExec != nil {
		log.Fatal(errExec)
	}
}

func transactionSQL(db *sql.DB) {
	transaction, errBegin := db.Begin()
	if errBegin != nil {
		log.Panic(errBegin)
	}
	defer transaction.Rollback()

	_, err := transaction.Exec("INSERT INTO golang VALUES (?, ?, ?)", nil, "Kim Cheolsoo", "inserted")
	if err != nil {
		log.Panic(err)
	}
	_, err = transaction.Exec("UPDATE golang SET etc=? WHERE etc=?", "transaction", "inserted")
	if err != nil {
		log.Panic(err)
	}

	err = transaction.Commit()
	if err != nil {
		log.Panic(err)
	}
}

func ExmapleDBAdvanced() {
	DBURL := os.Getenv("DBURL")
	db, errConnect := sql.Open("mysql", DBURL)
	if errConnect != nil {
		panic(errConnect)
	}
	defer db.Close()

	preparedStatement(db)
	transactionSQL(db)
}
