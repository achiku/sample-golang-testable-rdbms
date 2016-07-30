package main

import "testing"

func TestNewDB(t *testing.T) {
	c := &DBConfig{
		Host:    "localhost",
		User:    "pgtest",
		DBName:  "pgtest",
		SSLMode: "disable",
	}
	db, err := NewDB(c)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
