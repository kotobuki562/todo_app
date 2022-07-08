package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
	// URLの定義
	// topの定義はcontrollersのroute_mainで定義している
	http.HandleFunc("/", top)
	// nilにするとデフォルトで404が返ってくる
	return http.ListenAndServe(":" + config.Config.Port, nil)
}