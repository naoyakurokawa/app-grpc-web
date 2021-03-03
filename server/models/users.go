package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"log"
)

var db *sqlx.DB

type User struct {
	ID    int
	Name  string
	Score int
}

type Userlist []User

func GetUsers(request pb.GetUsersRequest) ([]*pb.User, error) {
	// var users []User
	var userlist []*pb.User
	db, err := sqlx.Open("mysql", "root:test@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Println(err)
	}
	q := "SELECT * FROM users"
	err = db.Select(&userlist, q)
	if err != nil {
		panic(err)
	}
	return userlist, nil
}
