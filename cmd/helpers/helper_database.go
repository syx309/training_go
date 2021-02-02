package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "training_passwordku"
)

var DB *sql.DB

func notation() string {
	// connection string
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func InitDatabase() {
	if DB == nil {
		DB = openDatabase()
	}
}

func openDatabase() *sql.DB {
	db, err := sql.Open("postgres", notation())
	if err != nil {
		fmt.Println("Failed to connect")
		panic(err)
	}
	ping(db)
	fmt.Println("Successfully connected to Database!")
	return db
}

func CloseDatabase() {
	err := DB.Close()
	CheckError(err)
}

func ping(db *sql.DB) {
	err := db.Ping()
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

