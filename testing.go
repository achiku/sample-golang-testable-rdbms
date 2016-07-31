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

// TestCreateUserData creates t1 test data
func TestCreateUserData(t *testing.T, q Queryer, u *AppUser) *AppUser {
	if err := u.Insert(q); err != nil {
		t.Fatal(err)
	}
	return u
}
