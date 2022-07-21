package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	// GETならテンプレートファイルをレスポンス
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		// formで飛んできたデータを解析する
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// ユーザーを作成する
		user := models.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		// ユーザーを保存する
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		
		// 正常に終了した場合はトップページにリダイレクトさせる
		http.Redirect(w, r, "/", 302)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
		_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		// cookieを作成する
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	session := models.Session{UUID: cookie.Value}
	err = session.DeleteSessionByUUID()
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/login", 302)
}