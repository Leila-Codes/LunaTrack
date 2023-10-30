package postgres

import (
	"Luna_Track/log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// TODO: schema deployment/implementation

type PostgesDB struct {
	conn *sql.DB
}

func NewPostgresDB(host, user, pass, dbname string, port int) *PostgesDB {
	var logger = log.GetLogger()
	var connString = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, dbname)
	logger.Debugf("Attempting DB connection... %s\n", connString)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Database Connection Successful")

	return &PostgesDB{conn: db}
}
