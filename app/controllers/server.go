package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
	// URLの定義
	http.HandleFunc("/", top)
	// nilにするとデフォルトで404が返ってくる
	return http.ListenAndServe(":" + config.Config.Port, nil)
}