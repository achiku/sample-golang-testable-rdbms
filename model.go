package main

import (
	"time"

	"github.com/pkg/errors"
)

// T1 test table
type T1 struct {
	ID        int
	Val       string
	Status    string
	CreatedAt time.Time
}

// Insert creates t1 in db
func (t *T1) Insert(q Queryer) error {
	err := q.QueryRow(`
	INSERT INTO t1 (
		val
		, created_at
	) VALUES ($1, $2) RETURNING id
	`, t.Val, t.CreatedAt).Scan(&t.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create t1")
	}
	return nil
}

// GetT1ByID get t1 by id
func GetT1ByID(q Queryer, id int) (*T1, error) {
	var t T1
	err := q.QueryRow(`
	SELECT
		id
		, val
		, created_at
	FROM t1
	WHERE id = $1
	`, id).Scan(
		&t.ID,
		&t.Val,
		&t.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get t1")
	}
	return &t, nil
}

// GetT1WeekAgoFromNow gets t1 created a week ago
func GetT1WeekAgoFromNow(q Queryer, n time.Time) ([]T1, error) {
	wkAgo := n.AddDate(0, 0, -7)
	rows, err := q.Query(`
	SELECT
		id
		, val
		, created_at
	FROM t1
	WHERE created_at <= $1
	`, wkAgo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get t1")
	}
	var l []T1
	for rows.Next() {
		var t T1
		rows.Scan(&t.ID, &t.Val, &t.CreatedAt)
		l = append(l, t)
	}
	return l, nil
}

// GetT1ByVal gets t1 by val
func GetT1ByVal(q Queryer, val string) ([]T1, error) {
	rows, err := q.Query(`
	SELECT
		id
		, val
		, created_at
	FROM t1
	WHERE val = $1
	`, val)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get t1")
	}
	var l []T1
	for rows.Next() {
		var t T1
		rows.Scan(&t.ID, &t.Val, &t.CreatedAt)
		l = append(l, t)
	}
	return l, nil
}
