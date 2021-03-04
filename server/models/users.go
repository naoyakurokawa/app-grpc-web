package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"log"
)

func GetUsers(request pb.GetUsersRequest) ([]*pb.User, error) {
	// var users []User
	var userlist []*pb.User
	db, err := sqlx.Open("mysql", "root:test@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	q := "SELECT * FROM users"
	err = db.Select(&userlist, q)
	if err != nil {
		log.Println(err)
	}
	return userlist, nil
}
