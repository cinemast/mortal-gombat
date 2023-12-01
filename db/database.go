package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

type Db struct {
	Queries
	c  *sql.DB
	Tx *sql.Tx
}

func (db *Db) Connection() *sql.DB {
	return db.c
}

func NewDb(c *sql.DB) *Db {
	return &Db{
		Queries: *New(c),
		c:       c,
		Tx:      nil,
	}
}

func Tx(ctx context.Context, db *Db, statements func(ctx context.Context, tx *Db) error) error {
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	txDb := &Db{db.Queries, db.c, tx}
	err = statements(ctx, txDb)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func TxR[T any](ctx context.Context, db *Db, statements func(ctx context.Context, tx *Db) (*T, error)) (*T, error) {
	tx, err := db.c.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	txDb := &Db{db.Queries, db.c, tx}
	r, err := statements(ctx, txDb)
	if err != nil {
		return nil, err
	}
	return r, tx.Commit()
}

type SqliteDateTime struct {
	time.Time
}

func (ts *SqliteDateTime) Scan(src any) error {
	t, err := time.Parse(time.DateTime, src.(string))
	if err != nil {
		return err
	}
	*ts = SqliteDateTime{t}
	return nil
}

func (ts SqliteDateTime) Value() (driver.Value, error) {
	return driver.Value(ts.Format(time.DateTime)), nil
}
