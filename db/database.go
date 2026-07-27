package db

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// var err error

	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBName,
	)

	DB, err = sql.Open(

		"postgres",dsn)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL!")

	return DB

}