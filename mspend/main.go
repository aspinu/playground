package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("home.html"))

type handlerData struct {
	DbData []spending
	Total  int
}

func addSpedingHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name := r.PostFormValue("spending-name")
	amountStr := r.PostFormValue("spending-amount")
	category := r.PostFormValue("spending-category")
	amount, err := strconv.Atoi(strings.Trim(amountStr, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	newSpending := spendingLong{
		SpendingAmount:   amount,
		SpendingName:     name,
		SpendingCategory: category,
	}

	addSpending(db, newSpending)
	tmpl.ExecuteTemplate(w, "spending-location-element", nil)
}

func showSpendingHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dataFromSql, total := selectAllSpendings(db)
	var sumTotal int
	for _, vals := range total {
		sumTotal += vals
	}
	spendingData := handlerData{
		DbData: dataFromSql,
		Total:  sumTotal,
	}
	tmpl := template.Must(template.ParseFiles("listsp.html"))
	tmpl.Execute(w, spendingData)
}

func showFilteredHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	year := "2025"
	month := r.PostFormValue("months")

	data, total := selectFilteredSpendings(db, year, month)
	var sumTotal int
	for _, vals := range total {
		sumTotal += vals
	}
	spendingFilteredData := handlerData{
		DbData: data,
		Total:  sumTotal,
	}

	tmpl := template.Must(template.ParseFiles("listsp.html"))
	tmpl.Execute(w, spendingFilteredData)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("filterm.html"))
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
