package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type (
	queryEngine struct {
		db *sql.DB
	}
)

var (
	currentConnection *queryEngine
)

func GetQueryEngine() *queryEngine {
	if currentConnection != nil {
		return currentConnection
	}

	var q *queryEngine
	q = new(queryEngine)
	q.db = q.getDb()
	currentConnection = q
	return q
}

func (q queryEngine) getDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./movies.db")
	log.Printf("Opening connection to movies database")

	if err != nil {
		log.Fatal("Could not open connection to DB: %q", err)
	}

	return db
}

func (q queryEngine) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return q.db.Query(query, args...)
}

func (q queryEngine) Exec(query string, args ...interface{}) (*sql.Result, error) {
	result, err := q.db.Exec(query, args...)
	return &result, err
}
