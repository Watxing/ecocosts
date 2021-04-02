package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type payload struct {
	Name string
	Data interface{}
}

func fetchClientStocks(c client) ([]stock, error) {
	if c.id == 0 {
		return nil, errors.New("id is not set")
	}

	rows, err := db.Query("SELECT symbol, quantity FROM stock WHERE client_id = $1", c.id)
	if err != nil {
		return nil, err
	}

	stocks := make([]stock, 0)

	for rows.Next() {
		var s stock
		err := rows.Scan(&s.symbol, &s.quantity)
		if err != nil {
			return nil, err
		}
		s.getPrice()
		stocks = append(stocks, s)
	}

	return stocks, nil
}

func dashHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")

	if path != "" && path != "index" {
		http.NotFound(w, r)
		return
	}

	var c client
	if err := c.readCookie(w, r); err != nil {
		http.Redirect(w, r, "/auth", http.StatusFound)
	}

	stocks, err := fetchClientStocks(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(stocks)

	p := payload{c.Name, nil}

	if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
