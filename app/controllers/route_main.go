package controllers

import (
	"log"
	"net/http"
	"text/template"
)

// w,rを受け取るとハンドラーとして定義できる
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "Hello")
}