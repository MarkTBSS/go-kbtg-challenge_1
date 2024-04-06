package postgres

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Database *sql.DB
}

func New() (*Postgres, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	database, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &Postgres{Database: database}, nil
}
