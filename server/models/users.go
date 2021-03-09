package models

import (
	// "fmt"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func GetUsers(ctx context.Context, db *sqlx.DB, request pb.GetUsersRequest) ([]*pb.User, error) {
	var userlist []*pb.User
	q := "SELECT * FROM users"
	err := db.SelectContext(ctx, &userlist, q)
	if err != nil {
		log.Println(err)
	}
	// log.Printf("userList : %s", userlist)
	return userlist, nil
}

func CreateUser(ctx context.Context, db *sqlx.DB, request pb.CreateUserRequest) (string, error) {
	//パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(request.GetPassword()), 10)
	if err != nil {
		return "", err
	}
	hash_password := string(hash)
	user := pb.User{
		Name:     request.GetName(),
		Score:    request.GetScore(),
		Photourl: request.GetPhotourl(),
		Password: hash_password,
	}
	query := `INSERT INTO users (id, name, score, photourl, password) VALUES (:id, :name, :score, :photourl, :password);`
	tx := db.MustBegin()
	_, err = tx.NamedExecContext(ctx, query, &user)
	if err != nil {
		log.Printf("error : %s", err)
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

func LoginUser(ctx context.Context, db *sqlx.DB, request pb.LoginRequest) (int32, error) {
	// var token string
	var user []*pb.User
	q := `SELECT * FROM users WHERE NAME = ?;`
	err := db.SelectContext(ctx, &user, q, request.GetName())
	// err := user_service.CheckLoginUserRequest(request)
	log.Printf("user : %s", user[0].Id)
	// db.Where("email = ?", request.Email).First(&user)
	err = bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(request.GetPassword()))

	if err != nil {

		return -1, status.New(codes.InvalidArgument, "ユーザー名 または　パスワードが間違っています").Err()
	}

	// token, err = user_service.CreateToken(user)
	// if err != nil {
	// 	return -1, "", status.New(codes.Unknown, "作成失敗").Err()
	// }
	return user[0].Id, nil
}
