package postgres

import (
	"database/sql"
	"log"

	"github.com/MarkTBSS/go-kbtg-challenge_1/configs"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Database *sql.DB
}

func New() (*Postgres, error) {
	database, err := sql.Open("postgres", configs.DatabaseURL())
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
