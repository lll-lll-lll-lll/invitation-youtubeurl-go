package repository

import (
	"github.com/jmoiron/sqlx"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"
)

func InsertUser(req *fb.RegisterUser, db *sqlx.DB) error {
	if err := Transaction(db, req, InsertUserFunc); err != nil {
		return err
	}
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
