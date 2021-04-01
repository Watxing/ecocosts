package main

import (
	"fmt"
	"net/http"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	var c client

	if c.readCookie(w, r) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		c.Name = r.FormValue("name")
		c.pass = r.FormValue("pass")
		action := r.FormValue("action")
		fmt.Println(c.Name)

		if action == "login" {
			if err := c.login(w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if action == "register" {
			if c.exist() {
				http.Error(w, "username taken", http.StatusInternalServerError)
				return
			}
			if err := c.insert(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, "/auth", http.StatusFound)
	}

	if err := templates.ExecuteTemplate(w, "auth.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
