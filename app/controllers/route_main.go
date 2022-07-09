package controllers

import (
	"net/http"
)

// w,rを受け取るとハンドラーとして定義できる
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
}