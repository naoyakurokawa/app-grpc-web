package models

import (
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	"log"
)

func GetUsers(db, request pb.GetUsersRequest) ([]*pb.User, error) {
	var userlist []*pb.User
	var err error
	q := "SELECT * FROM users"
	err = db.Select(&userlist, q)
	if err != nil {
		log.Println(err)
	}
	return userlist, nil
}

func CreateUser(db, request pb.CreateUserRequest) (string, error) {
	var err error
	log.Printf("request : %s", request)
	user := pb.User{
		Name:     request.GetName(),
		Score:    request.GetScore(),
		Photourl: request.GetPhotourl(),
	}
	query := `INSERT INTO users (id, name, score, photourl) VALUES (:id, :name, :score, :photourl);`
	tx := db.MustBegin()
	_, err := tx.NamedExec(query, &user)
	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()
		// エラー内容を返却します。
		return "登録失敗", err
	}
	tx.Commit()
	return "登録成功", err
}
