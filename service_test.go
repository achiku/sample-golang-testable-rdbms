package main

import (
	"testing"
	"time"
)

func TestCreateAppUser(t *testing.T) {
	db, cleanup := setupServiceTest(t)
	defer cleanup()

	user := &AppUser{
		Name:      "testuser",
		Status:    UserStatusActive,
		CreatedAt: time.Now(),
	}
	u, e, err := CreateAppUser(db, user)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("AppUser: %+v", u)
	t.Logf("AppEvent: %+v", e)
}

func TestCreateAppUserFailure(t *testing.T) {
	db, cleanup := setupServiceTest(t)
	defer cleanup()

	user := &AppUser{
		Name:      "testuser",
		Status:    UserStatusActive,
		CreatedAt: time.Now(),
	}
	u, e, err := CreateAppUser(db, user)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("AppUser: %+v", u)
	t.Logf("AppEvent: %+v", e)
}
