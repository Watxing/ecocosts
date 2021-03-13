CREATE TABLE client (
	id SERIAL,
	name TEXT UNIQUE NOT NULL,
	pass TEXT NOT NULL,
	PRIMARY KEY (id),
	CHECK (LENGTH(name) < 20)
);

CREATE TABLE budget (
	id SERIAL,
	client_id INT NOT NULL,
	cat_id INT NOT NULL,
	amount FLOAT,
	PRIMARY KEY (id),
	FOREIGN KEY(client_id) REFERENCES client(id) ON DELETE CASCADE,
	FOREIGN KEY(cat_id) REFERENCES category(id)
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
	FOREIGN KEY (client_id) REFERENCES client(id) ON DELETE CASCADE,
	FOREIGN KEY (cat_id) REFERENCES category(id)
);

CREATE TABLE category (
	id SERIAL,
	description TEXT NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE stock (
	client_id INT NOT NULL,
	symbol TEXT NOT NULL,
	quantity INT NOT NULL,
	FOREIGN KEY (client_id) REFERENCES client(id) ON DELETE CASCADE,
	CHECK (quantity > 0)
);	
