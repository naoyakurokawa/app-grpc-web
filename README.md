## 起動方法

### client

```sh
cd client
npm run dev
```

### server

```sh
cd server
go run main.go
```

### proxy

```sh
docker-compose up
```

## Minio について

- 画像の表示
  - 管理画面 http://127.0.0.1:9000/ で、バケットポリシーを設定
  - prifix: \* とし、Read and Write
