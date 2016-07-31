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

// User test table
type User struct {
	ID        int
	Name      string
	Status    string
	CreatedAt time.Time
}

// Event test table
type Event struct {
	ID        int
	UserID    int
	Category  string
	CreatedAt time.Time
}

// EventSummary test table
type EventSummary struct {
	ID        int
	UserID    int
	Category  string
	Count     int
	UpdatedAt time.Time
}

// Insert creates app_user in db
func (t *User) Insert(q Queryer) error {
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
func GetUserByID(q Queryer, id int) (*User, error) {
	var t User
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
func GetUserWeekAgoFromNow(q Queryer, n time.Time) ([]User, error) {
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
	var l []User
	for rows.Next() {
		var t User
		rows.Scan(&t.ID, &t.Name, &t.Status, &t.CreatedAt)
		l = append(l, t)
	}
	return l, nil
}

// GetUserByName gets app_user by val
func GetUserByName(q Queryer, name string) ([]User, error) {
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
	var l []User
	for rows.Next() {
		var t User
		rows.Scan(&t.ID, &t.Name, &t.Status, &t.CreatedAt)
		l = append(l, t)
	}
	return l, nil
}
