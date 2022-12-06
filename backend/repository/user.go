package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"
)

func InsertUser(req *fb.RegisterUser, db *sqlx.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("INSERT INTO users(id, name) VALUES($1,$2)", req.ID, req.Name)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if rows != 1 {
		tx.Rollback()
		return fmt.Errorf("expected 1 row, got %d", rows)
	}
	tx.Commit()
	return nil
}

func InsertUserFunc(req interface{}, db *sqlx.DB) error {
	castedReq := req.(fb.RegisterUser)
	stmt, err := db.Prepare("INSERT INTO users(id, name) VALUES($1,$2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(castedReq.ID, castedReq.Name)
	if err != nil {
		return err
	}
	return nil
}
