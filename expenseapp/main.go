package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	tmpl    = template.Must(template.ParseFiles("home.html"))
	newTime = time.Now()
)

type handlerData struct {
	DbData []Spending
	Total  float64
}

func newSpendingLong(name, category string, amount float64) SpendingLong {
	newSpending := SpendingLong{}
	newSpending.Spending.SpendingAmount = amount
	newSpending.Spending.SpendingName = name
	newSpending.Spending.SpendingCategory = category
	newSpending.Day = newTime.Day()
	newSpending.Month = newTime.Month().String()
	newSpending.Year = newTime.Year()
	return newSpending
}

func addSpedingHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := r.PostFormValue("spending-name")
	amountStr := r.PostFormValue("spending-amount")
	if amountStr == "" {
		fmt.Fprint(w, "This field is required")
		return
	}
	category := r.PostFormValue("spending-category")
	amount, err := strconv.ParseFloat(strings.Trim(amountStr, "\n"), 64)
	if err != nil {
		fmt.Fprint(w, "Please enter a valid decimal number")
		return
	}

	addSpending(db, newSpendingLong(name, category, amount))
	tmpl.ExecuteTemplate(w, "spending-location-element", nil)
}

func showSpendingHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dataFromSql, total := selectAllSpendings(db)
	var sumAmount float64
	for _, vals := range total {
		sumAmount += vals
	}
	spendingData := handlerData{
		DbData: dataFromSql,
		Total:  sumAmount,
	}
	tmpl := template.Must(template.ParseFiles("spending_list.html"))
	tmpl.Execute(w, spendingData)
}

func showFilteredHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var year string
	if r.PostFormValue("year") == "" {
		year = strconv.Itoa(newTime.Year())
	} else {
		year = r.PostFormValue("year")
	}
	month := r.PostFormValue("months")

	dataFromSql, total := selectFilteredSpendings(db, year, month)
	var sumAmount float64
	for _, vals := range total {
		sumAmount += vals
	}
	spendingFilteredData := handlerData{
		DbData: dataFromSql,
		Total:  sumAmount,
	}

	tmpl := template.Must(template.ParseFiles("spending_list.html"))
	tmpl.Execute(w, spendingFilteredData)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("months_filter.html"))
	tmp.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/show/", showSpendingHandler)
	mux.HandleFunc("/add-spending/", addSpedingHandler)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/show-filtered/", showFilteredHandler)
	mux.HandleFunc("/filter-page/", filterHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
