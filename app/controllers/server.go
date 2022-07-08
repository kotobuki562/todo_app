package controllers

import (
	"net/http"
	"todo_app/config"
)

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