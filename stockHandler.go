package main

import (
	"net/http"
	"strconv"
)

func stockHandler(w http.ResponseWriter, r *http.Request) {
	var c client
	if err := c.readCookie(w, r); err != nil {
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		symbol := r.FormValue("symbol")
		qty, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s := stock{
			Symbol: symbol,
			Quantity: qty,
		}

		if err := s.insert(c.id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := c.updateStocks(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(w, "stock.html", c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
