package controllers

import (
	"log"
	"net/http"
)

// w,rを受け取るとハンドラーとして定義できる
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
		return
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		// 得たsessionでuserを取得する
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// userのtodos取得のメソッドを叩いてtodosを取得する
		todos, _ := user.GetTodosByUser()
		// Userのstructに[]Todoを加えて返す
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}