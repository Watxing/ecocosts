package main

import (
	"net/http"
)

func viewTransaction() ([]transaction, error) {
	var t []transaction
	rows, err := db.Query(`SELECT id, client_id, cat_id, amount, balance,
		description, time FROM transaction`)
	if err != nil {
		return t, err
	}

	for rows.Next() {
		var r transaction
		err := rows.Scan(&r.ID, &r.Client_id, &r.Cat_id, &r.Amount, &r.Balance,
			&r.Description, &r.Time)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	if err = rows.Err(); err != nil {
		return t, err
	}

	return t, nil
}

func ledgerHandler(w http.ResponseWriter, r *http.Request) {
	t, err := viewTransaction()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := templates.ExecuteTemplate(w, "ledger.html", t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
