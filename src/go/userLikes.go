package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	// "strconv"
	bdd "../bdd"
)

func UserLikes(w http.ResponseWriter, req *http.Request) {

	fmt.Print("Page Mes Likes ✔️ \n")

	t, err := template.ParseFiles("./template/userLikes.html", "./template/Header.html")

	if err != nil {
		fmt.Println(err.Error())
	}



	var errBdd int

	uuidValue := readCookie(w, req)
	var userValue string
	errBdd, userValue = bdd.GetUserByUUID(uuidValue)
	ReturnError500(w, errBdd)

	var arr [][]string                              // Création d'un tableau de tableau de string pour stocker tout les posts .
	var posts []PostData                            // Création d'un tableau qui prendra toutes les valeurs utiles pour un post.
	errBdd, arr = bdd.GetPosteLikeByUser(userValue) // On va push dans arr tout les posts qui ont été liké par l'utilisateur.
	ReturnError500(w, errBdd)


	// Boucle pour récupérer les valeurs de chaque post et les envoyer dans le tableau posts.
	for _, post := range arr {
		p := PostData{
			ID:       post[0],
			UserName: post[1],
			Post:     post[2],
			Date:     post[5],
		}
		posts = append(posts, p)
	}

	pageData := LoginWrapper {
		IsLogged: VerifyCookie(req),
		UserConnected: userValue, 
		Data: posts,
	}

	t.Execute(w, pageData)
}
