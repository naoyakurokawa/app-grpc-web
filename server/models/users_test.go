package models

import (
	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
)

const (
	Name     = "テスト"
	Score    = 1234
	Photourl = "https://test"
	Password = "abcd12341231"
)

// ユーザ一覧取得テスト
func TestGetUsers(t *testing.T) {
	db, err := connectDb()
	require.NoError(t, err)
	ctx, err := prepare(db)
	require.NoError(t, err)

	request := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, request)
	require.NoError(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte("abcd12341231"))
	require.NoError(t, err)
	assert.Equal(t, users[0].Name, "テスト")
	assert.Equal(t, users[0].Score, int32(1234))
	assert.Equal(t, users[0].Photourl, "https://test")
}

//ユーザー登録テスト
func TestCreateUser(t *testing.T) {
	db, err := connectDb()
	require.NoError(t, err)
	m := map[string]string{"login_token": "test123"}
	md := metadata.New(m)
	ctx := metadata.NewIncomingContext(context.Background(), md)
	err = initUserTable(db)
	require.NoError(t, err)
	//ユーザー作成
	request := pb.CreateUserRequest{Name: Name, Score: Score, Photourl: Photourl, Password: Password}
	userId, err := CreateUser(ctx, db, request)
	require.NoError(t, err)
	//セッション登録
	session := &pb.Session{
		Uuid:   "test123",
		Name:   Name,
		Userid: userId,
	}
	err = CreateSession(session, db)
	require.NoError(t, err)
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	require.NoError(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte("abcd12341231"))
	require.NoError(t, err)
	assert.Equal(t, users[0].Name, "テスト")
	assert.Equal(t, users[0].Score, int32(1234))
	assert.Equal(t, users[0].Photourl, "https://test")
}

//ユーザー取得テスト
func TestGetUserById(t *testing.T) {
	db, err := connectDb()
	require.NoError(t, err)
	ctx, err := prepare(db)
	require.NoError(t, err)
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	require.NoError(t, err)
	fetchedUser, err := GetUserById(ctx, db, users[0].Id)
	require.NoError(t, err)
	err = bcrypt.CompareHashAndPassword([]byte(fetchedUser[0].Password), []byte("abcd12341231"))
	require.NoError(t, err)
	assert.Equal(t, fetchedUser[0].Name, "テスト")
	assert.Equal(t, fetchedUser[0].Score, int32(1234))
	assert.Equal(t, fetchedUser[0].Photourl, "https://test")
}

//ユーザー削除テスト
func TestDeleteUser(t *testing.T) {
	db, err := connectDb()
	require.NoError(t, err)
	ctx, err := prepare(db)
	require.NoError(t, err)
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	require.NoError(t, err)
	DeleteUser(ctx, db, users[0].Id)
	fetchedUser, err := GetUserById(ctx, db, users[0].Id)
	require.NoError(t, err)
	assert.Equal(t, len(fetchedUser), int(0))
}

// //ログインテスト
func TestLoginUser(t *testing.T) {
	db, err := connectDb()
	require.NoError(t, err)
	ctx, err := prepare(db)
	require.NoError(t, err)
	//作成したユーザー取得
	getUsersRequest := pb.GetUsersRequest{}
	users, err := GetUsers(ctx, db, getUsersRequest)
	require.NoError(t, err)
	fetchedUser, err := GetUserById(ctx, db, users[0].Id)
	require.NoError(t, err)
	//セッションDB登録
	session := &pb.Session{
		Uuid:   createUUID(),
		Name:   fetchedUser[0].Name,
		Userid: fetchedUser[0].Id,
	}
	err = CreateSession(session, db)
	require.NoError(t, err)

	//ログイン処理テスト
	loginRequest := pb.LoginRequest{Name: Name, Password: Password}
	_, _, err = LoginUser(ctx, db, loginRequest)
	require.NoError(t, err)
}

//テスト用データ作成用プライベートメソッド
func prepare(db *sqlx.DB) (context.Context, error) {
	//テーブル初期化
	err := initUserTable(db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//テストユーザー作成
	userId, err := createUserForTest(db)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//セッション登録
	session := &pb.Session{
		Uuid:   "test123",
		Name:   Name,
		Userid: userId,
	}
	err = CreateSession(session, db)
	if err != nil {
		return nil, err
	}

	//メタデータ設定
	ctx := context.Background()
	m := map[string]string{"login_token": "test123"}
	md := metadata.New(m)
	ctx = metadata.NewIncomingContext(context.Background(), md)

	return ctx, nil
}

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
	q = `DELETE FROM session;`
	_, err = db.Exec(q)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//テスト共通処理
// func TestMain(m *testing.M) {
// 	prepare()
// 	// 開始処理
// 	log.Print("setup")
// 	// パッケージ内のテストの実行
// 	code := m.Run()
// 	// 終了処理
// 	log.Print("tear-down")
// 	// テストの終了コードで exit
// 	os.Exit(code)
// }

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
