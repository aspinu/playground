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

//
// type Spending struct {
// 	Name     string
// 	Amount   int
// 	Category string
// 	Total    int
// }
//
// var X int = 0
//
// func displayDB(w http.ResponseWriter, db *sql.DB)                {}
// func addInDB(w http.ResponseWriter, r *http.Request, db *sql.DB) {}
//
// func addSpedingHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.PostFormValue("spending-name")
// 	amount := r.PostFormValue("spending-amount")
//
// 	category := r.PostFormValue("spending-category")
//
// 	amountInt, err := strconv.Atoi(strings.Trim(amount, "\n"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	ptrt := &X
// 	*ptrt = *ptrt + amountInt
// 	tmpl.ExecuteTemplate(w, "spending-list-element", Spending{Name: name, Amount: amountInt, Category: category, Total: *ptrt})
// }
//

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

	newSpending := spending{
		spendingAmount:   amount,
		spendingName:     name,
		spendingCategory: category,
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

	dataFormSql := selectAllSpendings(db)
	tmpl.ExecuteTemplate(w, "spendings-list", dataFormSql)
}

func main() {
	// db, err := sql.Open("sqlite3", "./names.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	//
	http.HandleFunc("/add-spending/", addSpedingHandler)
	http.HandleFunc("/", showSpendingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
