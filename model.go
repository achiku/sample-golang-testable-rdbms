package main

import (
	"time"

	"github.com/pkg/errors"
)

// UserStatus
const (
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
	UserStatusClosed   = "closed"
)

// AppUser test table
type AppUser struct {
	ID        int
	Name      string
	Status    string
	CreatedAt time.Time
}

// Insert creates app_user in db
func (t *AppUser) Insert(q Queryer) error {
	err := q.QueryRow(`
	INSERT INTO app_user (
		name
		, status
		, created_at
	) VALUES ($1, $2, $3) RETURNING id
	`, t.Name, t.Status, t.CreatedAt).Scan(&t.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create app_user")
	}
	return nil
}

// GetUserByID get app_user by id
func GetUserByID(q Queryer, id int) (*AppUser, error) {
	var t AppUser
	err := q.QueryRow(`
	SELECT
		id
		, name
		, status
		, created_at
	FROM app_user
	WHERE id = $1
	`, id).Scan(
		&t.ID,
		&t.Name,
		&t.Status,
		&t.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get app_user")
	}
	return &t, nil
}

// GetUserWeekAgoFromNow gets app_user created a week ago
func GetUserWeekAgoFromNow(q Queryer, n time.Time) ([]AppUser, error) {
	wkAgo := n.AddDate(0, 0, -7)
	rows, err := q.Query(`
	SELECT
		id
		, name
		, status
		, created_at
	FROM app_user
	WHERE created_at <= $1
	`, wkAgo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get app_user")
	}
	var l []AppUser
	for rows.Next() {
		var t AppUser
		rows.Scan(&t.ID, &t.Name, &t.Status, &t.CreatedAt)
		l = append(l, t)
	}
	return l, nil
}

// GetUserByName gets app_user by val
func GetUserByName(q Queryer, name string) ([]AppUser, error) {
	rows, err := q.Query(`
	SELECT
		id
		, name
		, status
		, created_at
	FROM app_user
	WHERE name = $1
	`, name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get app_user")
	}
	var l []AppUser
	for rows.Next() {
		var t AppUser
		rows.Scan(&t.ID, &t.Name, &t.Status, &t.CreatedAt)
		l = append(l, t)
	}
	return l, nil
}

// AppEvent test table
type AppEvent struct {
	ID        int
	UserID    int
	Category  string
	CreatedAt time.Time
}

// Insert insert event
func (e *AppEvent) Insert(q Queryer) error {
	err := q.QueryRow(`
	INSERT INTO app_event (
		user_id
		, category
		, created_at
	) VALUES ($1, $2, $3) RETURNING id
	`, e.UserID, e.Category, e.CreatedAt).Scan(&e.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create app_event")
	}
	return nil
}

// AppEventSummary test table
type AppEventSummary struct {
	ID        int
	UserID    int
	Category  string
	Count     int
	UpdatedAt time.Time
}
