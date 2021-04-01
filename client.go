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
func (c *client) exist() error {
	err := db.QueryRow("SELECT id FROM client WHERE name = $1", c.Name).Scan(&c.id)
	if err != nil {
		return err
	}

	return nil
}
