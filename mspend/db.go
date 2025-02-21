package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type spending struct {
	Id               int
	SpendingName     string
	SpendingAmount   int
	SpendingCategory string
}

func addSpending(db *sql.DB, newSpending spending) {
	stmt, err := db.Prepare("INSERT INTO spendings (id, spendings_name, spendings_amount, spendings_category) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(nil, newSpending.SpendingName, newSpending.SpendingAmount, newSpending.SpendingCategory)
	defer stmt.Close()
}

func selectAllSpendings(db *sql.DB) []spending {
	rows, err := db.Query("SELECT id, spendings_name, spendings_amount, spendings_category FROM  spendings")
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
		err = rows.Scan(&curentSpending.Id, &curentSpending.SpendingName, &curentSpending.SpendingAmount, &curentSpending.SpendingCategory)
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
