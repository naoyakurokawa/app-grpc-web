package models

import (
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
	db, err := sqlx.Open("mysql", "root:test@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Fatalln(err)
	}
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

//今後 user_idをもとにしたユーザー取得のテストにも用いるため、idをreturnする
func CreateUserForTest(db *sqlx.DB) (int32, error) {
	//パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 10)
	if err != nil {
		log.Printf("error : %s", err)
		return -1, err
	}
	hash_password := string(hash)
	user := pb.User{
		Name:     Name,
		Score:    Score,
		Photourl: Photourl,
		Password: hash_password,
	}
	query := `INSERT INTO users (name, score, photourl, password) VALUES (:name, :score, :photourl, :password);`
	tx, err := db.Beginx()
	if err != nil {
		log.Printf("error : %s", err)
		return -1, err
	}
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
