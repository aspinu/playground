package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type spending struct {
	Id               int
	SpendingName     string
	SpendingAmount   int
	SpendingCategory string
}

type spendingLong struct {
	Id               int
	SpendingName     string
	SpendingAmount   int
	SpendingCategory string
	Day              int
	Month            string
	Year             int
}

var Year, Month, Day = time.Now().Date()

func addSpending(db *sql.DB, newSpending spendingLong) {
	stmt, err := db.Prepare("INSERT INTO spendings (id, spendings_name, spendings_amount, spendings_category, year, month, day) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(nil, newSpending.SpendingName, newSpending.SpendingAmount, newSpending.SpendingCategory, Year, Month, Day)
	defer stmt.Close()
}

func selectFilteredSpendings(db *sql.DB, slectedYear, selectedMonth string) []spending {
	rows, err := db.Query("SELECT id, spendings_name, spendings_amount, spendings_category FROM  spendings WHERE year = '" + slectedYear + "' and month = '" + selectedMonth + "'")
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
