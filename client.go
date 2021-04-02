package main

import (
	"net/http"
	"errors"
)

type client struct {
	id   int
	Name string
	pass string
}

// Insert values into database. This inserts the password as plain-text. Do NOT
// do this in a production setting :). This is for debugging purposes.
func (c *client) insert() error {
	if c.Name == "" {
		return errors.New("Name is not set")
	} else if c.pass == "" {
		return errors.New("pass is not set")
	}

	_, err := db.Exec("INSERT INTO client (name, pass) VALUES ($1, $2)", c.Name, c.pass)
	if err != nil {
		return err
	}

	return nil
}

// Check if in database.
func (c *client) exist() bool {
	err := db.QueryRow("SELECT id FROM client WHERE name = $1", c.Name).Scan(&c.id)
	if err != nil {
		return false
	}

	return true
}

func (c *client) passCorrect() error {
	var pass string

	err := db.QueryRow("SELECT pass FROM client WHERE name = $1", c.Name).Scan(&pass)
	if err != nil {
		return err
	}

	if c.pass != pass {
		return errors.New("invalid pass")
	}

	return nil
}

// If used in production, it might be preferably to create some sort of cookie
// session manager that manages cookies more securely. perhaps implement bcrypt
// later for some security at least? https://godocs.io/golang.org/x/crypto/bcrypt
func (c *client) login(w http.ResponseWriter) error {
	if err := c.passCorrect(); err != nil {
		return err
	}

	name := http.Cookie {
		Name: "name",
		Value: c.Name,
		MaxAge: 86400, // 24 hours from now
		Secure: true,
	}

	pass := http.Cookie {
		Name: "pass",
		Value: c.pass,
		MaxAge: 86400, // 24 hours from now
		Secure: true,
	}

	http.SetCookie(w, &name)
	http.SetCookie(w, &pass)

	return nil
}

func (c *client) readCookie(w http.ResponseWriter, r *http.Request) bool {
	name, _ := r.Cookie("name")
	pass, _ := r.Cookie("pass")

	c.Name = name.Value
	c.pass = pass.Value
	return true
}
