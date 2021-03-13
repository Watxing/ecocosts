CREATE TABLE clients (
	id SERIAL,
	name TEXT UNIQUE NOT NULL,
	pass TEXT NOT NULL,
	PRIMARY KEY (id),
	CHECK (LENGTH(name) < 20)
);

CREATE TABLE budgets (
	id SERIAL,
	client_id INT NOT NULL,
	cat_id INT NOT NULL,
	amount FLOAT,
	PRIMARY KEY (id),
	FOREIGN KEY(client_id) REFERENCES clients(id) ON DELETE CASCADE,
	FOREIGN KEY(cat_id) REFERENCES categories(id)
);

CREATE TABLE transaction (
	id SERIAL,
	client_id INT NOT NULL,
	cat_id INT,
	amount FLOAT NOT NULL,
	balance FLOAT NOT NULL,
	description TEXT,
	time TIMESTAMP NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
	FOREIGN KEY (cat_id) REFERENCES categories(id)
);

CREATE TABLE categories (
	id SERIAL,
	description TEXT NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE stocks (
	client_id INT NOT NULL,
	stock TEXT NOT NULL,
	qty INT NOT NULL,
	FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
	CHECK (qty > 0)
);	
