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
		bdd.Like(IdToLike, "Tao") // Appliquer la fonction de getBdd.go
	}

	// Récupération de la valeur des commentaires
	getCommentValue := req.FormValue("commentaire")
	getCommentID := req.FormValue("IdComment")
	StrToInt, _ := strconv.Atoi(getCommentID) 
	if getCommentValue != "" {
		bdd.MakeComment("Tao", getCommentValue, StrToInt)
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
			NbLike : bdd.GetLikeNb(nbLike),
			Date: post[5],
		}
		posts = append(posts, p)
	}

	// sessionCookie()

	// Gestion de l'erreur 404
	if req.URL.Path == "/" { //verification de l'URL
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	t.Execute(w, posts)
}
