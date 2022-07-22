package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filename ...string,) {
	var files []string
	for _, file := range filename {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// アクセス権限をかける
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// URLの正規表現
var validPath = regexp.MustCompile("^/todos/(edit|update)/([0-9]+)$")

// URLをパースする。
// qoでIDを取得してfn(ハンドラ)を返している
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 受け取ったURLのpathを解析してvalidPathにマッチするか確認する
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		// エンドポイントの2番目にはIDが入っているのでintに変換する
		qi, err := strconv.Atoi(q[2])
		// qiがintでない場合はエラーを起こすのでエラーハンドリングを行う
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartMainServer() error {
	// viewsディレクトリにあるjsとcssを読み込む
	files := http.FileServer(http.Dir(config.Config.Static))
	// staticディレクトリは存在しないのでhttp.StripPrefixで取り除く
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URLの定義
	// topの定義はcontrollersのroute_mainで定義している
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	// validPathを使う都合で/todos/edit/にする
	// todoEdit関数は(w, r, int)を引数に取っている
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))

	// nilにするとデフォルトで404が返ってくる
	return http.ListenAndServe(":" + config.Config.Port, nil)
}