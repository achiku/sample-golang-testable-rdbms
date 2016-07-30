package main

import (
	"testing"
	"time"
)

func TestT1Insert(t *testing.T) {
	tx, cleanup := setupModelTest(t)
	defer cleanup()

	t1 := T1{
		Val:       "testval",
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

func TestGetT1ByID(t *testing.T) {
	tx, cleanup := setupModelTest(t)
	defer cleanup()

	t1 := &T1{
		Val:       "t01",
		CreatedAt: time.Now(),
	}
	TestCreateT1Data(t, tx, t1)
	t2 := &T1{
		Val:       "t02",
		CreatedAt: time.Now(),
	}
	TestCreateT1Data(t, tx, t2)

	target, err := GetT1ByID(tx, t1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if target.Val != t1.Val {
		t.Errorf("got %s want %s", target.Val, t1.Val)
	}
}

func TestGetT1WeekAgoFromNow(t *testing.T) {
	tx, cleanup := setupModelTest(t)
	defer cleanup()

	n := time.Now()
	t0 := &T1{
		Val:       "t00",
		CreatedAt: n.AddDate(0, 0, -7),
	}
	TestCreateT1Data(t, tx, t0)
	t1 := &T1{
		Val:       "t01",
		CreatedAt: n.AddDate(0, 0, -10),
	}
	TestCreateT1Data(t, tx, t1)
	t2 := &T1{
		Val:       "t02",
		CreatedAt: n.AddDate(0, 0, -1),
	}
	TestCreateT1Data(t, tx, t2)

	targets, err := GetT1WeekAgoFromNow(tx, n)
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range targets {
		t.Logf("%+v", a)
	}
}
