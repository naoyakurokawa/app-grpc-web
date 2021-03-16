package models

import (
	"context"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func GetUsers(ctx context.Context, db *sqlx.DB, request pb.GetUsersRequest) ([]*pb.User, error) {
	//メタデータ取得
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return nil, nil
	}
	//メターデータの中のlogin_tokenを参照
	login_token := md["login_token"][0]
	//sessionテーブルにlogin_tokenに紐づくデータが存在するか確認
	s, err := GetSessionByUuid(ctx, db, login_token)
	//sessionテーブルに存在しなければreturn
	if len(s) == 0 {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var userlist []*pb.User
	q := "SELECT * FROM users"
	err = db.SelectContext(ctx, &userlist, q)
	if err != nil {
		log.Println(err)
		return nil, err
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
	tx, err := db.Beginx()
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

func LoginUser(ctx context.Context, db *sqlx.DB, request pb.LoginRequest) (int32, string, error) {
	var user []*pb.User

	//ユーザー名必須チェック
	err := CheckRequired("名前", request.GetName())
	if err != nil {
		return -1, "", err
	}

	//パスワード必須チェック
	err = CheckRequired("パスワード", request.GetPassword())
	if err != nil {
		return -1, "", err
	}

	//ユーザー存在チェック
	q := `SELECT * FROM users WHERE NAME = ?;`
	err = db.SelectContext(ctx, &user, q, request.GetName())
	if err != nil {
		return -1, "", status.New(codes.InvalidArgument, "ユーザー名が間違っています").Err()
	}

	//パスワード一致チェック
	err = CheckMatchPassword(user[0].Password, request.GetPassword())
	if err != nil {
		return -1, "", err
	}

	//セッションDB登録
	session := &pb.Session{
		Uuid:   createUUID(),
		Name:   user[0].Name,
		Userid: user[0].Id,
	}
	err = CreateSession(session, db)
	if err != nil {
		log.Println(err)
		return -1, "", err
	}

	return user[0].Id, session.Uuid, nil
}

func createUUID() (uuidobj string) {
	u, _ := uuid.NewUUID()
	uuidobj = u.String()
	return uuidobj
}

func CreateSession(sess *pb.Session, db *sqlx.DB) error {
	query := `INSERT INTO session (id, uuid, name, userid) VALUES (:id, :uuid, :name, :userid);`
	tx, err := db.Beginx()
	_, err = tx.NamedExec(query, &sess)
	if err != nil {
		log.Printf("error : %s", err)
		// エラーが発生した場合はロールバックします。
		tx.Rollback()
		// エラー内容を返却します。
		return err
	}
	tx.Commit()
	return err
}

func GetSessionByUuid(ctx context.Context, db *sqlx.DB, uuid string) ([]*pb.Session, error) {
	var session []*pb.Session
	q := `SELECT * FROM session WHERE Uuid = ?;`
	err := db.SelectContext(ctx, &session, q, uuid)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return session, nil
}

//Validation関連のメソッド

//必須チェック
func CheckRequired(feild string, input string) error {
	if input == "" {
		return status.New(codes.InvalidArgument, feild+"は必須です").Err()
	} else {
		return nil
	}
}

//パスワード一致確認
func CheckMatchPassword(dbData string, input string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbData), []byte(input))
	if err != nil {
		return status.New(codes.InvalidArgument, "パスワードが間違っています").Err()
	} else {
		return nil
	}
}
