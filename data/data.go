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

// If the connection has already been created we return it. Otherwise
// we allocate a new queryEngin and connect to the database
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

// These are very basic functions, but it establishes a pattern that I
// would want to keep. The connect itself is hidden. This would allow
// us to manage multiple connects or do connection pooling more easily
func (q queryEngine) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return q.db.Query(query, args...)
}

func (q queryEngine) Exec(query string, args ...interface{}) (*sql.Result, error) {
	result, err := q.db.Exec(query, args...)
	return &result, err
}
