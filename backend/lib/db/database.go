package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlDB interface {
	InitDB() (*sqlx.DB, error)
	Open() (*sqlx.DB, error)
}

type MySql struct {
	datasource string
}

func NewMySql(datasource string) *MySql {
	return &MySql{
		datasource: datasource,
	}
}

func (md *MySql) InitDB(count int) (*sqlx.DB, error) {
	db, err := md.Open()
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(25)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return md.InitDB(count)
	}
	return db, nil
}

func (md *MySql) Open() (*sqlx.DB, error) {
	dbcon, err := sqlx.Open("mysql", md.datasource)
	if err != nil {
		return nil, fmt.Errorf("failed db init. %s", err)
	}
	return dbcon, nil
}

func SetUpDB(count int) (*sqlx.DB, error) {
	d := DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
	}.String()

	m := NewMySql(d)
	DB, err := m.InitDB(count)
	if err != nil {
		DB.Close()
		return nil, err
	}
	return DB, nil
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (dc DBConfig) String() string {
	if dc.User == "" {
		dc.User = "user"
	}
	if dc.Password == "" {
		dc.Password = "password"
	}
	if dc.Host == "" {
		dc.Host = "db"
	}
	if dc.Port == "" {
		dc.Port = "3306"
	}
	if dc.DBName == "" {
		dc.DBName = "cyberDB"
	}
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dc.User, dc.Password, dc.Host, dc.Port, dc.DBName)
	return s
}
