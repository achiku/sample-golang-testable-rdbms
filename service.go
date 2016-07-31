package main

import (
	"time"

	"github.com/pkg/errors"
)

// CreateAppUser create app user
func CreateAppUser(db DBer, u *AppUser) (*AppUser, *AppEvent, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create tx")
	}
	defer tx.Rollback()

	if err := u.Insert(tx); err != nil {
		return nil, nil, errors.Wrap(err, "failed to create AppUser")
	}
	e := &AppEvent{
		UserID:    u.ID,
		Category:  "create_user",
		CreatedAt: time.Now(),
	}
	if err := e.Insert(tx); err != nil {
		return nil, nil, errors.Wrap(err, "failed to create AppEvent")
	}
	tx.Commit()
	return u, e, nil
}
