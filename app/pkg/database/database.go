package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func New(user string, password string, host string, port int, dbname string) (*pgx.Conn, error) {

	psqlUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		user, password, host, port, dbname,
	)

	dbCon, err := pgx.Connect(context.Background(), psqlUrl)
	if err != nil {
		log.Fatalf("Unable to connect database %s", err)
	}

	return dbCon, nil
}
