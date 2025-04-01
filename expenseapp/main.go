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
	Total  float64
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
	amount, err := strconv.ParseFloat(strings.Trim(amountStr, "\n"), 64)
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
	year := "2025"
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
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	er := createUsersTable(db)
	if er != nil {
		log.Fatal(er)
	}
	defer db.Close()
	tmpl := template.Must(template.ParseFiles("auth.html"))

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
	mux.HandleFunc("/auth-login/", registerHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
