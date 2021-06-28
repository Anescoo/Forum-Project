package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	bdd "../bdd"
)

func UserPost(w http.ResponseWriter, req *http.Request) {

	fmt.Print("Page TL ✔️ \n")

	t, _ := template.ParseFiles("./template/userPost.html", "./template/Header.html")

	// Gestion de la supression des posts
	getPostID := req.FormValue("delete")
	IdToSuppr, err := strconv.Atoi(getPostID)
	if err == nil {
		bdd.DeletePoste(IdToSuppr)
	}

	getNewValue := req.FormValue("sendUpdate")
	getPostIDupdate := req.FormValue("update")
	// Gestion de l'update des posts
	IdtoUpdate, err := strconv.Atoi(getPostIDupdate)
	if err == nil {
		bdd.UpdatePoste(IdtoUpdate, getNewValue)
	}

	// même methode que dans "Accueil.go" mais pour les posts de l'utilisateurs on prend juste les valeurs dont on a besoin
	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetPosteByUser("Tao")

	for _, post := range arr {
		p := PostData{
			ID:   post[0],
			Post: post[2],
			Date: post[5],
		}
		posts = append(posts, p)
	}

	t.Execute(w, posts)
}
