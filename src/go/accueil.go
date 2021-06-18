package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	bdd "../bdd"
)

type PostData struct {
	UserName string
	Post     string
	Date     string
}

func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/header.html")
	fmt.Print("Page d'accueil ✔️ \n")

	getPostValue := req.FormValue("PostValue")
	bdd.MakeUser("Tao", "louis.teilliais@gmail.com", "Karim69lattrik")
	bdd.MakePoste("Tao", string(getPostValue), "test")

	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetAllPoste()

	for _, post := range arr {
		p := PostData{
			UserName: post[1],
			Post:     post[2],
			Date:     post[5],
		}
		posts = append(posts, p)
	}

	if req.URL.Path == "/" { //verification de l'URL
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	t.Execute(w, posts)
}
