package main

import (
	"database/sql"
	"testing"
)

func setupModelTest(t *testing.T) (*sql.Tx, func()) {
	db, err := sql.Open("postgres", "user=pgtest dbname=pgtest sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		tx.Rollback()
		db.Close()
	}
	return tx, cleanup
}

// TestCreateT1Data creates t1 test data
func TestCreateT1Data(t *testing.T, q Queryer, t1 *T1) {
	if err := t1.Insert(q); err != nil {
		t.Fatal(err)
	}
}
