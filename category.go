package main

type category struct {
	ID          int
	Description string
}

func getCategories() ([]category, error) {
	var categories []category

	rows, err := db.Query("SELECT id, description FROM category")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c category
		if err := rows.Scan(&c.ID, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
