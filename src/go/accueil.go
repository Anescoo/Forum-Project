package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"strconv"
	bdd "../bdd"
	// authent "./authent"
)
// Création d'une structure avec toutes les "infos" de notre post. 
type PostData struct { 
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
	
	// récupérer id post
	getPostID := req.FormValue("delete") 
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

	// récupérer id post
	getIDLike := req.FormValue("like") 
	IdToLike, e := strconv.Atoi(getIDLike) 
	if e == nil{
		bdd.Like(IdToLike, "Louis") // Appliquer la fonction de getBdd.go
	}
	
	var arr [][]string 
	var posts []PostData 
	_, arr = bdd.GetAllPoste() 
	
	// Parcourir et remplir notre tableau des données que l'on veut  
	for _, post := range arr { 
		nbLike, _ := strconv.Atoi(post[0])
		p := PostData { 
			ID: post[0],
			UserName: post[1],
			Post: post[2],
			NbLike : bdd.GetLikeNb(nbLike),
			Date: post[5],
		}
		posts = append(posts, p) 
	}

	// authent.sessionCookie()

	// Gestion de l'erreur 404
	if req.URL.Path == "/" { 
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	t.Execute(w, posts)
}
