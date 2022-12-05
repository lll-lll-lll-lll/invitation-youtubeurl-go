package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSql struct {
	Db *sqlx.DB
}

func NewPostgreSql() *PostgreSql {
	return &PostgreSql{}
}

type User struct {
	UserID   int
	Password string
}

func (ps *PostgreSql) Open() {
	Db, err := sqlx.Open("postgres", "host=postgres user=app_user password=password dbname=app_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	ps.Db = Db
}
