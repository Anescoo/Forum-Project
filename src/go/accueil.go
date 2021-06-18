package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"strconv"
	bdd "../bdd"
)

type PostData struct {
	UserName string
	Post string
	Date string
	NbLike int
	ID string
	Categorie string
}



func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/header.html")
	fmt.Print("Page d'accueil ✔️ \n")

	getPostID := req.FormValue("delete") // récupérer id post
	IdToSuppr, err:= strconv.Atoi(getPostID)
	if err == nil {
		bdd.DeletePoste(IdToSuppr) // Appliquer la fonction de getBdd.go
	}
	
	getPostValue := req.FormValue("PostValue") // récupérer id post
	getCategorieValue := req.FormValue("categorie")
	if getPostValue != "" {
		bdd.MakePoste("Tao", string(getPostValue), string(getCategorieValue))
			// Appliquer la fonction de getBdd.go	
	}
	var tab []string
	for _, cat := range tab {
		categories := bdd.MakeCategorie(string(getCategorieValue))
		tab = append(tab, categories)
	}
	



	getIDLike := req.FormValue("like") // récupérer id post
	IdToLike, e := strconv.Atoi(getIDLike) 
	if e == nil{
		bdd.Like(IdToLike, "Louis") // Appliquer la fonction de getBdd.go
	}
	 

	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetAllPoste()
	for _, post := range arr {
		nbLike, _ := strconv.Atoi(post[0]) 
		p := PostData {
			ID: post[0],
			UserName: post[1],
			Post: post[2],
			Categorie : post[3],
			NbLike : bdd.GetLikeNb(nbLike),
			Date: post[5],
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
