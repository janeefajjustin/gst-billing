package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		product_code INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT NOT NULL,
		product_price FLOAT NOT NULL,
		product_gst FLOAT NOT NULL
	)
	`

	_, err = DB.Exec(createProductsTable)

	if err != nil {
		panic("Could not create products table.")
	}

	createBillingTable := `
	CREATE TABLE IF NOT EXISTS billings (
		billing_id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT NOT NULL,
		product_code INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		amount FLOAT NOT NULL,
		FOREIGN KEY(product_code) REFERENCES products(product_code)
	)
	`
	_, err = DB.Exec(createBillingTable)

	if err != nil {
		panic("Could not create Billing table.")
	}
}
