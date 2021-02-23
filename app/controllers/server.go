package controllers

import (
	"awesomeProject/awesomeProject/todo_app/app/models"
	"awesomeProject/awesomeProject/todo_app/config"
	"fmt"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string){
	var files []string
	for _, file :=range filenames{
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error){
	cookie, err := r.Cookie("_cookie")
	if err == nil{
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok{
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}





func StartMainServer() error {
	//静的ファイル（CSSとJS）を設定し、紐づける役割
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//urlとハンドラーの紐付けをし、返す役割
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
