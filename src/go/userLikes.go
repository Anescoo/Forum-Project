package handlers

import (
	"fmt"
	"text/template"
	"net/http"
	
	// bdd "../bdd"
)

func UserLikes(w http.ResponseWriter, req *http.Request){
	
	fmt.Print("Page Mes Posts ✔️ \n")
	
	t, _ := template.ParseFiles("./template/userPost.html", "./template/header.html")

	t.Execute(w, nil)
}