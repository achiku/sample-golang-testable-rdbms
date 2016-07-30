package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // sql database
	"github.com/pkg/errors"
)

// DB database
type DB struct {
	*sql.DB
}

// DBConfig config
type DBConfig struct {
	Host    string
	User    string
	DBName  string
	SSLMode string
}

// Queryer query/exec level interface
// Queryerが渡される場所は限定される。modelの中だけで利用されるような感じ。
// serviceレベル(modelの一つ上のレイヤ)ではBegin/Rollback/Commit等ができる。
// 通常のsql.DBでもsql.Txでもどちらでも使えるようなinterface。
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// DBer database level interface
// 後でテスト用にモックする為アプリケーション内ではこのinterfaceを利用する。
type DBer interface {
	Queryer
	Begin() (*sql.Tx, error)
	Ping() error
}

// NewDB creates DB
func NewDB(c *DBConfig) (DBer, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s", c.User, c.DBName, c.SSLMode))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db")
	}
	return &DB{db}, nil
}
