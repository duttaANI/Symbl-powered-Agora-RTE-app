package models

import (
	"github.com/jinzhu/gorm"

	// Importing postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database contains a pointer to the database object
type Database struct {
	*gorm.DB
}

// CreateDB is used to initialize a new database connection
func CreateDB(dbURL string) (*Database, error) {
	// db, err := gorm.Open("postgres","postgres://<user>:<password>@<host>/<db_name>?sslmode=disable")
	// db, err := gorm.Open("postgres", "postgres://myprojectuser:password@aniruddha-G3-3579/postgres?sslmode=disable")
	db, err := gorm.Open("postgres", "postgres://postgres:password@localhost/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// TODO: Setup Production Migrations
	return &Database{db}, nil
}
