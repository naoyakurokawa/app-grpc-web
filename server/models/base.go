package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB
var err error

func init() {
	db, err = sqlx.Open("mysql", "root:test@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Fatalln(err)
	}
}
