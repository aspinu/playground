package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type spending struct {
	id               int
	spendingName     string
	spendingCategory string
	spendingAmount   int
}

func addSpending(db *sql.DB, newSpending spending) {
	stmt, _ := db.Prepare("INSERT INTO spendings (id, spending_name, spending_amount, spending_cathegory) VALUES (?, ?, ?, ?)")
	stmt.Exec(nil, newSpending.spendingName, newSpending.spendingCategory, newSpending.spendingAmount)
	defer stmt.Close()
}

func selectAllSpendings(db *sql.DB) []spending {
	rows, err := db.Query("SELECT * FROM spendings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	mySpending := make([]spending, 0)
	for rows.Next() {
		curentSpending := spending{}
		err = rows.Scan(&curentSpending.id, &curentSpending.spendingName, &curentSpending.spendingAmount, &curentSpending.spendingCategory)
		if err != nil {
			log.Fatal(err)
		}
		mySpending = append(mySpending, curentSpending)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return mySpending
}
