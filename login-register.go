package forum

import (
	"html/template"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/user/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		usernameorlogin := r.Form["username"]
		password := r.Form["password"]
		testemail := RemoveFromString(usernameorlogin[0])
		if IsEmail(testemail) {

		}
		SendToDBLogin(usernameorlogin[0], password[0])
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("/user/register.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		email := r.Form["email"]
		username := r.Form["username"]
		password := r.Form["password"]
		SendToDBRegister(email[0], username[0], password[0])
	}
}

func IsEmail(str string) bool {
	bool := false
	if strings.Contains(str, "@") {
		bool = true
	}
	return bool
}
