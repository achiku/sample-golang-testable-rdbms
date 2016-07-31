package main

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq" // sql database
)

func init() {
	txdb.Register("txdb", "postgres", "user=pgtest dbname=pgtest sslmode=disable")
}

func setupModelTest(t *testing.T) (*sql.Tx, func()) {
	db, err := sql.Open("txdb", "dummy")
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

// TestDB db for test
type TestDB struct {
	*sql.DB
}

func setupServiceTest(t *testing.T) (DBer, func()) {
	db, err := sql.Open("txdb", "dummy")
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		db.Close()
	}
	return db, cleanup
}

// TestCreateUserData creates t1 test data
func TestCreateUserData(t *testing.T, q Queryer, u *AppUser) *AppUser {
	if err := u.Insert(q); err != nil {
		t.Fatal(err)
	}
	return u
}
