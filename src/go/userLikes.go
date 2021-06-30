package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	bdd "../bdd"
)

func UserLikes(w http.ResponseWriter, req *http.Request) {

	fmt.Print("Page Mes Posts ✔️ \n")
	
	t, err := template.ParseFiles("./template/userLikes.html", "./template/header.html")

	if err != nil {
		fmt.Print(err.Error)
	}

	var arr [][]string                     // Création d'un tableau de tableau de string pour stocker tout les posts .
	var posts []PostData                   // Création d'un tableau qui prendra toutes les valeurs utiles pour un post.
	_, arr = bdd.GetPosteLikeByUser("The_Real_Legend") // On va push dans arr tout les posts qui ont été liké par l'utilisateur.

	// Boucle po_r récupérer les valeurs de chaque post et les envoyer dans le tableau posts.
	for _, post := range arr {
		p := PostData{
			ID:       post[0],
			UserName: post[1],
			Post:     post[2],
			Date:     post[5],
		}
		posts = append(posts, p)
	}

	t.Execute(w, posts)
}
