package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	container "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"

	db "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/db"
)

func main() {
	router := gin.Default()
	postgresql := db.NewPostgreSql()
	postgresql.Open()
	defer postgresql.Db.Close()
	router.POST("/invitation_code", func(ctx *gin.Context) {
		var input container.Input
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}
		con, err := container.New(input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := insert(con, postgresql.Db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": con.Code})

	})
	router.Run(":8080")
}
func transaction(db *sqlx.DB, req interface{}, f func(req interface{}, db *sqlx.DB) error) error {
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
func insert(con *container.Container, db *sqlx.DB) error {

	postCode := PostCode{Code: con.Code}
	if err := transaction(db, postCode, insertCodeFunc); err != nil {
		return err
	}
	return nil
}

type PostCode struct {
	Code string `json:"code"`
}

func insertCodeFunc(req interface{}, db *sqlx.DB) error {
	castedReq := req.(PostCode)
	stmt, err := db.Prepare("INSERT INTO invitation_codes(code) VALUES($1) returning code")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(castedReq.Code)
	if err != nil {
		return err
	}
	return nil
}
