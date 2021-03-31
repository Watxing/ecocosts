package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

var db *sql.DB

type client struct {
	id   int
	name string
	pass string
}

type category struct {
	id          int
	description string
}

type budget struct {
	client_id int
	cat_id    int
	amount    float64
}

type transaction struct {
	id          int
	client_id   int
	cat_id      int
	amount      float64
	balance     float64
	description string
	time        Time
}

type stock struct {
	client_id int
	symbol    string
	quantity  int
}

// initializes the database. panics if a failure.
func init() {
	// do NOT use in production
	connStr := "user=postgres password=? dbname=ecocosts sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
