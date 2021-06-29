package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	bdd "../bdd"
	// authent "./authent"
)

// Création d'une structure avec toutes les "infos" de notre post.
type PostData struct {
	UserName 	   string
	Post     	   string
	Date     	   string
	NbLike         int
	ID             string
	CommentArr [][]string
	// Categorie []string
}

type ComsData struct {
	Pseudo		string
	Com			string
}

func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/Header.html")
	fmt.Print("Page d'accueil ✔️ \n")

	// Delete les posts

	getPostID := req.FormValue("delete")
	IdToSuppr, err := strconv.Atoi(getPostID)
	if err == nil {
		bdd.DeletePoste(IdToSuppr)
	}

	// getCategorieValue := req.FormValue("categorie")
	// bdd.MakeCategorie(string(getCategorieValue))

	// Gestion des cookies
	sessionCookie(w, req)

	// deleteCookie(w, req)

	getPostValue := req.FormValue("PostValue")
	getSelectValue := req.FormValue("selectCategorie")

	// Vérification si l'utilisateur est connecté
	if verifyCookie(w, req) == true {
		if getPostValue != "" {
			bdd.MakePoste("The_Real_Legend", string(getPostValue), string(getSelectValue))
		}
	}
	// else {
	// 	http.Redirect(w, req, "/connexion", http.StatusSeeOther)
	// }

	// Liker un post
	getIDLike := req.FormValue("like")
	IdToLike, e := strconv.Atoi(getIDLike)
	if e == nil {
		bdd.Like(IdToLike, "Louis")
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

	// var commentArr [][]string
	// _, commentArr = bdd.GetCommentByPoste(2)
	// fmt.Println(commentArr)


	// Parcourir et remplir notre tableau des données que l'on veut
	for _, post := range arr {
		nbLike, _ := strconv.Atoi(post[0])
		commentID, _ := strconv.Atoi(post[0])
		_, coms := bdd.GetCommentByPoste(commentID)
		p := PostData{
			ID:       	post[0],
			UserName: 	post[1],
			Post:     	post[2],
			NbLike:   	bdd.GetLikeNb(nbLike),
			Date:     	post[5],
			CommentArr: coms,
		}
		posts = append(posts, p)
	}

	// Gestion de l'erreur 404
	if req.URL.Path == "/" {
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	t.Execute(w, posts)
}
