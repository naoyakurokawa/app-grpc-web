package models

import (
	// "fmt"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"log"
)

func GetUsers(ctx context.Context, db *sqlx.DB, request pb.GetUsersRequest) ([]*pb.User, error) {
	var userlist []*pb.User
	q := "SELECT * FROM users"
	err := db.SelectContext(ctx, &userlist, q)
	if err != nil {
		log.Println(err)
	}
	return userlist, nil
}

func CreateUser(ctx context.Context, db *sqlx.DB, request pb.CreateUserRequest) (string, error) {
	log.Printf("request : %s", request)
	user := pb.User{
		Name:     request.GetName(),
		Score:    request.GetScore(),
		Photourl: request.GetPhotourl(),
	}
	query := `INSERT INTO users (id, name, score, photourl) VALUES (:id, :name, :score, :photourl);`
	tx := db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, &user)
	if err != nil {
		// エラーが発生した場合はロールバックします。
		tx.Rollback()
		// エラー内容を返却します。
		return "登録失敗", err
	}
	tx.Commit()
	return "登録成功", err
}

func GetUserById(ctx context.Context, db *sqlx.DB, id int32) ([]*pb.User, error) {
	log.Println(id)
	var user []*pb.User
	q := `SELECT * FROM users WHERE ID = ?;`
	err := db.SelectContext(ctx, &user, q, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func DeleteUser(ctx context.Context, db *sqlx.DB, id int32) error {
	q := `DELETE FROM users WHERE ID = ?;`
	_, err := db.ExecContext(ctx, q, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
