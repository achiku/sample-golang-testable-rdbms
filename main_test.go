package main

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // db
)

// TestMain service package setup/teardonw
func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", "user=pgtest dbname=pgtest sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect test db: %s", err.Error())
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS t1 (
		id serial
		, val text not null
		, created_at timestamp with time zone not null
	)`)
	if err != nil {
		log.Fatalf("failed to create test table: %s", err.Error())
	}
	code := m.Run()
	_, err = db.Exec(`DROP TABLE IF EXISTS t1`)
	if err != nil {
		log.Fatalf("failed to create test table: %s", err.Error())
	}
	os.Exit(code)
}
