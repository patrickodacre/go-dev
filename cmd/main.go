package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func main() {

	db, err := sql.Open("mysql", "admin:secret@tcp(go_dev_db)/go_dev")

	if err != nil { panic(err) }

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

		stmt, err_stmt := db.Prepare(`SELECT id, name FROM users`)

		if err_stmt != nil {
			http.Error(w, "Cannot retrieve users.", http.StatusInternalServerError)
			return
		}

		defer stmt.Close()

		rows, err_query := stmt.Query()

		if err_query != nil {
			http.Error(w, "Cannot retrieve users.", http.StatusInternalServerError)
			return
		}

		users := []User{}

		for rows.Next() {
			u := User{}
			err = rows.Scan(&u.ID, &u.Name)

			if err != nil {
				http.Error(w, "Cannot retrieve users.", http.StatusInternalServerError)
				return
			}

			users = append(users, u)
		}

		w.WriteHeader(http.StatusOK)

		for _, u := range users {
			fmt.Fprintf(w, "%v\n", u.Name)
		}
	})

	log.Println("listening on 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}