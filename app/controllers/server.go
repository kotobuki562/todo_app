package controllers

import (
	"fmt"
	"html/template"
	"net/http"
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

func StartMainServer() error {
	// viewsディレクトリにあるjsとcssを読み込む
	files := http.FileServer(http.Dir(config.Config.Static))
	// staticディレクトリは存在しないのでhttp.StripPrefixで取り除く
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URLの定義
	// topの定義はcontrollersのroute_mainで定義している
	http.HandleFunc("/", top)
	// nilにするとデフォルトで404が返ってくる
	return http.ListenAndServe(":" + config.Config.Port, nil)
}