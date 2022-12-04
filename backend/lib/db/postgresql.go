package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type PostgreSql struct {
	datasource string
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
	sql := "SELECT user_id, user_password FROM TEST_USER WHERE user_id=$1;"

	// preparedstatement の生成
	pstatement, err := Db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	// 検索パラメータ（ユーザID）
	queryID := 1
	// 検索結果格納用の TestUser
	var user User

	// queryID を埋め込み SQL の実行、検索結果1件の取得
	err = pstatement.QueryRow(queryID).Scan(&user.UserID, &user.Password)
	if err != nil {
		log.Fatal(err)
	}

	// 検索結果の表示
	fmt.Println(user.UserID, user.Password)

}
