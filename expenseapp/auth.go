package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Login struct {
	Username         string
	HashedPassowrd   string
	SessionToken     string
	CSRFToken        string
	UserExpenseTable string
}

var users = map[string]Login{}

func createExpUsrTable(db *sql.DB, username string) error {
	createUsersTable := "CREATE TABLE " + username + "(id, spendings_name, spendings_amount, spendings_category, year, month, day, PRIMARY KEY(id AUTOINCREMENT));"

	_, err := db.Exec(createUsersTable)

	return err
}

func createUsersTable(db *sql.DB) error {
	const createUsersTable = `
  CREATE TABLE IF NOT EXISTS "users" (
  "UserID" INTEGER, 
  "Username" TEXT,
  "HashedPassowrd" TEXT,
  "SessionToken" TEXT,
  "CSRFToken" TEXT,
  "UserExpenseTable" TEXT,
  PRIMARY KEY("UserID" AUTOINCREMENT)
);`
	_, err := db.Exec(createUsersTable)

	return err
}

func registerUser(db *sql.DB, newUser Login) error {
	stmt, err := db.Prepare("INSERT INTO users (UserID, Username, HashedPassowrd, SessionToken, CSRFToken, UserExpenseTable) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(nil, newUser.Username, newUser.HashedPassowrd, newUser.SessionToken, newUser.CSRFToken, newUser.UserExpenseTable)
	defer stmt.Close()
	er := createExpUsrTable(db, newUser.Username)
	if er != nil {
		log.Fatal(er)
	}

	return err
}

func checkUsername(db *sql.DB, username string) error {
	checkUserExists := "SELECT * from users where Username like '%" + username + "'%"
	_, err := db.Exec(checkUserExists)

	return err
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) < 8 || len(password) < 8 {
		err := http.StatusNotAcceptable
		http.Error(w, "Invalid username or password!", err)
		return
	}

	hashedPassowrd, err := hashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	error := checkUsername(db, username)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	newUser := Login{
		username,
		hashedPassowrd,
		"x",
		"x",
		username,
	}

	er := registerUser(db, newUser)
	if er != nil {
		err := http.StatusConflict
		http.Error(w, "User already exists", err)
		return
	}
	tmpl := template.Must(template.ParseFiles("home.html"))
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./spendings.db")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// user, ok := users[username]
	errs := checkUsername(db, username)
	if errs != nil {

		err := http.StatusUnauthorized
		http.Error(w, "Invalid username or password!", err)
		return
	}
	if err != nil || !chekcPasswordHash(password, user.HashedPassowrd) {
		err := http.StatusUnauthorized
		http.Error(w, "Invalid username or password!", err)
		return
	}

	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_Token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	fmt.Println("Login succesful!")
}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return

	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_Token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})
	username := r.FormValue("username")
	user, _ := users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	users[username] = user
	fmt.Fprintf(w, "Logged out succesfully")
}

func protected(w http.ResponseWriter, r *http.Request) {
	// http POST

	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return

	}
	username := r.FormValue("username")
	fmt.Fprintf(w, "Welcome: %s", username)
}
