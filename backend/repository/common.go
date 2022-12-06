package repository

import "github.com/jmoiron/sqlx"

func Transaction(db *sqlx.DB, req interface{}, f func(req interface{}, db *sqlx.DB) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() error {
		if err := recover(); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
		}
		return nil
	}()

	if err := f(req, db); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
