package main

import (
	"fmt"
	"log"
	"net/http"
)

type Login struct {
	HashedPassowrd string
	SessionToken   string
	CSRFToken      string
}

var users = map[string]Login{}

func register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) < 8 || len(password) < 8 {
		err := http.StatusNotAcceptable
		http.Error(w, "Invalid username or password!", err)
		return
	}
	if _, ok := users[username]; ok {
		err := http.StatusConflict
		http.Error(w, "User already exists", err)
		return
	}
	hashedPassowrd, err := hashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	users[username] = Login{
		HashedPassowrd: hashedPassowrd,
	}
	fmt.Fprintln(w, "User registered succesfully!")
}

func login(w http.ResponseWriter, r *http.Request) {
}

func logout(w http.ResponseWriter, r *http.Request) {
}

func protected(w http.ResponseWriter, r *http.Request) {
}
