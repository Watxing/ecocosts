package main

import (
	"database/sql"
	"time"
)

type transaction struct {
	ID          int
	Client_id   int
	Cat_id      sql.NullInt64
	Amount      float64
	Balance     float64
	Description sql.NullString
	Time        time.Time
}

// Inserts values into the database with incremented ID.
func (t *transaction) insert() error {
	// check values not null
	// insert into
	// error checks
	return nil
}
