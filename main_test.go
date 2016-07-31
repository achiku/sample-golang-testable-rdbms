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
	CREATE TABLE IF NOT EXISTS app_user (
		id serial
		, name text not null
		, status text not null
		, created_at timestamp with time zone not null
	);
	CREATE TABLE IF NOT EXISTS app_event (
		id serial
		, user_id int not null
		, category text not null
		, created_at timestamp with time zone not null
	);
	CREATE TABLE IF NOT EXISTS app_event_summary (
		id serial
		, user_id int not null
		, category text not null
		, count int not null
		, updated_at timestamp with time zone not null
	);
	`)
	if err != nil {
		log.Fatalf("failed to create test table: %s", err.Error())
	}
	code := m.Run()
	_, err = db.Exec(`
	DROP TABLE IF EXISTS app_user;
	DROP TABLE IF EXISTS app_event;
	DROP TABLE IF EXISTS app_event_summary;
	`)
	if err != nil {
		log.Fatalf("failed to create test table: %s", err.Error())
	}
	os.Exit(code)
}
