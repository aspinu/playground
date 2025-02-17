package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("home.html"))

type Spending struct {
	Name   string
	Amount int
	Tag    string
	Total  int
}

var X int = 0

func addSpedingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("spending-name")
	amount := r.PostFormValue("spending-amount")

	tag := r.PostFormValue("spending-tag")

	amountInt, err := strconv.Atoi(strings.Trim(amount, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	ptrt := &X
	*ptrt = *ptrt + amountInt
	tmpl.ExecuteTemplate(w, "spending-list-element", Spending{Name: name, Amount: amountInt, Tag: tag, Total: *ptrt})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	spending := make(map[string]Spending)

	tmpl.Execute(w, spending)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/add-spending/", addSpedingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
