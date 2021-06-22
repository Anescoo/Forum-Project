package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"strconv"
	bdd "../bdd"
	// authent "./authent"
)

type PostData struct { // on créer une structure avec toutes les "infos" de notre post. 
	UserName string
	Post string
	Date string
	NbLike int
	ID string
	// Categorie []string
}

// type CategorieStruct struct {
// 	Categorie []string 
// }

func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/header.html")
	fmt.Print("Page d'accueil ✔️ \n")

	getPostID := req.FormValue("delete") // récupérer id post
	IdToSuppr, err:= strconv.Atoi(getPostID)
	if err == nil {
		bdd.DeletePoste(IdToSuppr) // Appliquer la fonction de getBdd.go
	}

	// getCategorieValue := req.FormValue("categorie")
	// bdd.MakeCategorie(string(getCategorieValue))
	
	getPostValue := req.FormValue("PostValue") 
	getSelectValue := req.FormValue("selectCategorie")
	if getPostValue != "" {
		bdd.MakePoste("Tao", string(getPostValue), string(getSelectValue)) // Appliquer la fonction de getBdd.go	
	}
	
	getIDLike := req.FormValue("like") // récupérer id post
	IdToLike, e := strconv.Atoi(getIDLike) 
	if e == nil{
		bdd.Like(IdToLike, "Louis") // Appliquer la fonction de getBdd.go
	}
	
	var arr [][]string // on créer un tableau de tableau de string car la fonction GetAllPoste() en renvoie un 
	var posts []PostData // on créer un tableau de notre structure PostData
	_, arr = bdd.GetAllPoste() // on asigne a notre tableau la fonction 
	
	for _, post := range arr { // on parcours notre "Grand tableau"
		nbLike, _ := strconv.Atoi(post[0])
		p := PostData { // on va récupérer a chaque index les données que l'on veut
			ID: post[0],
			UserName: post[1],
			Post: post[2],
			NbLike : bdd.GetLikeNb(nbLike),
			Date: post[5],
		}
		posts = append(posts, p) // on remplit notre tableau des valeurs de "p" 
	}

	// authent.sessionCookie()

	// Gestion de l'erreur 404
	if req.URL.Path == "/" { //verification de l'URL
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	t.Execute(w, posts)
}
