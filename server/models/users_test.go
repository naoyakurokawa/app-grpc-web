package models

import (
	// "fmt"

	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"golang.org/x/crypto/bcrypt"
)

const (
	Name     = "テスト"
	Score    = 123
	Photourl = "https://test"
	Password = "abcd12341231"
)

// ユーザ取得 正常
func TestGetUsersSuccess(t *testing.T) {
	var db *sqlx.DB
	ctx := context.Background()
	InitUserTable(db)
	CreateUserForTest(db)
	request := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	if users[0].Name != "テスト" {
		t.Errorf("%v != %v", users[0].Name, "テスト")
	}
	if users[0].Score != 123 {
		t.Errorf("%v != %v", users[0].Score, 123)
	}
	if users[0].Photourl != "https://test" {
		t.Errorf("%v != %v", users[0].Photourl, "https://test")
	}
}

func InitUserTable(db *sqlx.DB) error {
	q := `DELETE FROM users;`
	_, err := db.Exec(q)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateUserForTest(db *sqlx.DB) (int32, error) {
	//パスワードのハッシュ化
	hash, _ := bcrypt.GenerateFromPassword([]byte(Password), 10)
	hash_password := string(hash)
	user := pb.User{
		Name:     Name,
		Score:    Score,
		Photourl: Photourl,
		Password: hash_password,
	}
	query := `INSERT INTO users (name, score, photourl, password) VALUES (:name, :score, :photourl, :password);`
	tx, err := db.Beginx()
	_, err = tx.NamedExec(query, &user)
	if err != nil {
		log.Printf("error : %s", err)
		// エラーが発生した場合はロールバックします。
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return -1, rollbackErr
		}
		// エラー内容を返却します。
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return user.Id, nil
}
