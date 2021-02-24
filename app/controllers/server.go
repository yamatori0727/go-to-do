package controllers

import (
	"awesomeProject/awesomeProject/todo_app/app/models"
	"awesomeProject/awesomeProject/todo_app/config"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
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

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil{
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil{
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
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
	http.HandleFunc("/todos/new", todoNews)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
