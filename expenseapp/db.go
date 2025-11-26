package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Spending struct {
	Id               int
	SpendingName     string
	SpendingAmount   float64
	SpendingCategory string
}

type SpendingLong struct {
	Spending Spending
	Day      int
	Month    string
	Year     int
}

func addSpending(db *sql.DB, newSpending SpendingLong) {
	stmt, err := db.Prepare("INSERT INTO spendings (id, spendings_name, spendings_amount, spendings_category, year, month, day) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(nil, newSpending.Spending.SpendingName, newSpending.Spending.SpendingAmount, newSpending.Spending.SpendingCategory, newSpending.Year, newSpending.Month, newSpending.Day)
	defer stmt.Close()
}

func selectFilteredSpendings(db *sql.DB, slectedYear, selectedMonth string) ([]Spending, []float64) {
	rows, err := db.Query("SELECT id, spendings_name, spendings_amount, spendings_category FROM spendings WHERE year = '" + slectedYear + "' and month = '" + selectedMonth + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	total := []float64{}
	expense := make([]Spending, 0)
	for rows.Next() {
		curentSpending := Spending{}
		err = rows.Scan(&curentSpending.Id, &curentSpending.SpendingName, &curentSpending.SpendingAmount, &curentSpending.SpendingCategory)
		if err != nil {
			log.Fatal(err)
		}
		total = append(total, curentSpending.SpendingAmount)
		expense = append(expense, curentSpending)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return expense, total
}

func selectAllSpendings(db *sql.DB) ([]Spending, []float64) {
	rows, err := db.Query("SELECT id, spendings_name, spendings_amount, spendings_category FROM  spendings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	total := []float64{}
	expense := make([]Spending, 0)
	for rows.Next() {
		curentSpending := Spending{}
		err = rows.Scan(&curentSpending.Id, &curentSpending.SpendingName, &curentSpending.SpendingAmount, &curentSpending.SpendingCategory)
		if err != nil {
			log.Fatal(err)
		}
		total = append(total, curentSpending.SpendingAmount)

		expense = append(expense, curentSpending)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return expense, total
}
