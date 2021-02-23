package controllers

import (
	"net/http"
)

//HTMLの紐付けをし、表示する役割
func top(w http.ResponseWriter, r *http.Request){
	_, err := session(w, r)
	if err != nil{
		generateHTML(w, "Hello", "layout","public_navbar", "top")
	}else{
		http.Redirect(w, r, "/todos", 302)
	}

}

func index(w http.ResponseWriter, r *http.Request){
	_, err := session(w, r)
	if err != nil{
		http.Redirect(w, r, "/", 302)
	}else{
		generateHTML(w, nil, "layout", "index", "private_navbar")
	}
}
