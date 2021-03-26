package models

import (
	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

const (
	Name     = "テスト"
	Score    = 1234
	Photourl = "https://test"
	Password = "abcd12341231"
)

// ユーザ一覧取得テスト
func TestGetUsers(t *testing.T) {
	//DB接続
	db, err := connectDb()
	if err != nil {
		t.Errorf("failed to open mysql connection: %v", err)
	}
	//テーブル初期化
	err = initUserTable(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	//テストユーザー作成
	_, err = createUserForTest(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	//セッション登録
	session := &pb.Session{
		Uuid:   "test123",
		Name:   Name,
		Userid: 1,
	}
	err = CreateSession(session, db)
	if err != nil {
		t.Errorf("failed to CreateSession: %v", err)
	}

	//メタデータ設定
	ctx := context.Background()
	// m := map[string]string{"login_token": "test123"}
	// md := metadata.New(m)
	// ctx := metadata.NewIncomingContext(context.Background(), md)
	request := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, request)
	if err != nil {
		t.Errorf("failed to GetUsers: %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte("abcd12341231"))
	if err != nil {
		t.Errorf("failed to Password: %v", err)
	}
	assert.Equal(t, users[0].Name, "テスト")
	assert.Equal(t, users[0].Score, int32(1234))
	assert.Equal(t, users[0].Photourl, "https://test")
}

//ユーザー登録テスト
func TestCreateUser(t *testing.T) {
	db, err := connectDb()
	if err != nil {
		t.Errorf("failed to open mysql connection: %v", err)
	}
	ctx := context.Background()
	err = initUserTable(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	request := pb.CreateUserRequest{Name: Name, Score: Score, Photourl: Photourl, Password: Password}
	_, err = CreateUser(ctx, db, request)
	if err != nil {
		t.Errorf("failed to CreateUser: %v", err)
	}
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	if err != nil {
		t.Errorf("failed to GetUsers: %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte("abcd12341231"))
	if err != nil {
		t.Errorf("failed to Password: %v", err)
	}
	assert.Equal(t, users[0].Name, "テスト")
	assert.Equal(t, users[0].Score, int32(1234))
	assert.Equal(t, users[0].Photourl, "https://test")
}

//ユーザー取得テスト
func TestGetUserById(t *testing.T) {
	db, err := connectDb()
	if err != nil {
		t.Errorf("failed to open mysql connection: %v", err)
	}
	ctx := context.Background()
	err = initUserTable(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	_, err = createUserForTest(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	if err != nil {
		t.Errorf("failed to GetUsers: %v", err)
	}
	fetchedUser, err := GetUserById(ctx, db, users[0].Id)
	err = bcrypt.CompareHashAndPassword([]byte(fetchedUser[0].Password), []byte("abcd12341231"))
	if err != nil {
		t.Errorf("failed to Password: %v", err)
	}
	assert.Equal(t, fetchedUser[0].Name, "テスト")
	assert.Equal(t, fetchedUser[0].Score, int32(1234))
	assert.Equal(t, fetchedUser[0].Photourl, "https://test")
}

//ユーザー削除テスト
func TestDeleteUser(t *testing.T) {
	db, err := connectDb()
	if err != nil {
		t.Errorf("failed to open mysql connection: %v", err)
	}
	ctx := context.Background()
	err = initUserTable(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	_, err = createUserForTest(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	if err != nil {
		t.Errorf("failed to GetUsers: %v", err)
	}
	DeleteUser(ctx, db, users[0].Id)
	fetchedUser, err := GetUserById(ctx, db, users[0].Id)
	assert.Equal(t, len(fetchedUser), int(0))
}

//ログインテスト
func TestLoginUser(t *testing.T) {
	db, err := connectDb()
	if err != nil {
		t.Errorf("failed to open mysql connection: %v", err)
	}
	ctx := context.Background()
	err = initUserTable(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	//テストユーザー作成
	_, err = createUserForTest(db)
	if err != nil {
		t.Errorf("failed to initUserTable: %v", err)
	}
	//作成したユーザー取得
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	if err != nil {
		t.Errorf("failed to GetUsers: %v", err)
	}
	fetchedUser, err := GetUserById(ctx, db, users[0].Id)

	//セッションDB登録
	session := &pb.Session{
		Uuid:   createUUID(),
		Name:   fetchedUser[0].Name,
		Userid: fetchedUser[0].Id,
	}
	err = CreateSession(session, db)
	if err != nil {
		t.Errorf("failed to CreateSession: %v", err)
	}

	//ログイン処理テスト
	loginRequest := pb.LoginRequest{Name: Name, Password: Password}
	_, _, err = LoginUser(ctx, db, loginRequest)
	if err != nil {
		t.Errorf("failed to Login: %v", err)
	}
}

//テスト用データ作成用プライベートメソッド
func connectDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:test@tcp(127.0.0.1:13306)/test")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initUserTable(db *sqlx.DB) error {
	q := `DELETE FROM users;`
	_, err := db.Exec(q)
	if err != nil {
		log.Println(err)
		return err
	}
	q = `ALTER TABLE users auto_increment = 1;`
	_, err = db.Exec(q)
	if err != nil {
		log.Println(err)
		return err
	}
	q = `DELETE FROM session;`
	_, err = db.Exec(q)
	if err != nil {
		log.Println(err)
		return err
	}
	q = `ALTER TABLE session auto_increment = 1;`
	_, err = db.Exec(q)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//今後 user_idをもとにしたユーザー取得のテストにも用いるため、idをreturnする
func createUserForTest(db *sqlx.DB) (int32, error) {
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
			log.Printf("rollbackerror : %s", rollbackErr)
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
