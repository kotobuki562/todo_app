## Golang で ToDo を作ってみる

`go run main.go`でサーバー起動

## SQLite3 の操作

`sqlite3 webapp.sql`で CLI 操作可能

終了したければ`.exit`

出来なかったら`;`をした後に`.exit`
これは直前の syntax が閉じられていないことが原因。

### 大まかな開発フロー

1. views にて UI の実装
2. controllers にてハンドラー関数の実装
3. controllers の server にて URL の登録

### 改善点

- cookie が nil の場合のエラー処理が出来ていない

JSON でデータをレスポンスする場合は

```main.go
// JSONデータをレスポンスに書き込む
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(user)
```

### package

go.mod を作成(パッケージを管理する)
`go mod init`

依存するパッケージをインストール
`go get -v`
すると go.mod にパッケージが追加され、go.sum が作成される

### heroku

heroku へログイン
`heroku login`

heroku を作成
`heroku create`

main ブランチを heroku にデプロイ
`git push heroku main`

postgresql を確認
`heroku pg:info`

postgresql の設定 init
`heroku addons:add heroku-postgresql`

`heroku pg:sql -c "\d"`

postgresql に入る
`heroku pg:sql`

ローカルの sql を実行する
`\i example.sql`

ポスグレから出る
`\q`
