package main

import (
	"testing"
	"time"
)

func TestUserInsert(t *testing.T) {
	tx, cleanup := setupModelTest(t)
	defer cleanup()

	t1 := User{
		Name:      "testname",
		Status:    UserStatusActive,
		CreatedAt: time.Now(),
	}
	if err := t1.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if t1.ID == 0 {
		t.Errorf("want id other than 0")
	}
	t.Logf("%+v", t1)
}

func TestGetUserByID(t *testing.T) {
	tx, cleanup := setupModelTest(t)
	defer cleanup()

	t1 := &User{
		Name:      "t01",
		Status:    UserStatusActive,
		CreatedAt: time.Now(),
	}
	TestCreateUserData(t, tx, t1)
	t2 := &User{
		Name:      "t02",
		Status:    UserStatusActive,
		CreatedAt: time.Now(),
	}
	TestCreateUserData(t, tx, t2)

	target, err := GetUserByID(tx, t1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if target.Name != t1.Name {
		t.Errorf("got %s want %s", target.Name, t1.Name)
	}
}

func TestGetUserWeekAgoFromNow(t *testing.T) {
	tx, cleanup := setupModelTest(t)
	defer cleanup()

	n := time.Now()
	t0 := &User{
		Name:      "t00",
		Status:    UserStatusActive,
		CreatedAt: n.AddDate(0, 0, -7),
	}
	TestCreateUserData(t, tx, t0)
	t1 := &User{
		Name:      "t01",
		Status:    UserStatusActive,
		CreatedAt: n.AddDate(0, 0, -10),
	}
	TestCreateUserData(t, tx, t1)
	t2 := &User{
		Name:      "t02",
		Status:    UserStatusActive,
		CreatedAt: n.AddDate(0, 0, -1),
	}
	TestCreateUserData(t, tx, t2)

	targets, err := GetUserWeekAgoFromNow(tx, n)
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range targets {
		t.Logf("%+v", a)
	}
}
