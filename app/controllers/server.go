package controllers

import (
	"fmt"
	"html/template"
	"net/http"
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

	// nilにするとデフォルトで404が返ってくる
	return http.ListenAndServe(":" + config.Config.Port, nil)
}