package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	// "time"

	bdd "../bdd"
)

// Création d'une structure avec toutes les "infos" de notre post.
type PostData struct {
	UserName   string
	Post       string
	Date       string
	NbLike     int
	ID         string
	CommentArr [][]string
	Categorie string
}

type LoginWrapper struct {
	IsLogged      bool
	UserConnected string
	Data          interface{}
}

func Accueil(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./template/Accueil.html", "./template/Header.html")

	if err != nil {
		fmt.Println(err.Error())
	}

	var errBdd int

	fmt.Print("Page d'accueil ✔️ \n")

	getPostValue := req.FormValue("PostValue")

	
	getSelectValue := req.FormValue("selectCategorie")
	

	getIDLike := req.FormValue("like")
	IdToLike, e := strconv.Atoi(getIDLike)

	getDislike := req.FormValue("Dislike")
	idToDislike, eDislike := strconv.Atoi(getDislike)

	uuidValue := readCookie(w, req)
	var userValue string
	errBdd, userValue = bdd.GetUserByUUID(uuidValue)
	if ReturnError500(w, errBdd) {
		return
	}

	// Vérification si l'utilisateur est connecté
	if VerifyCookie(req) == true {
		if getPostValue != "" {
			errBdd = bdd.MakePoste(userValue, string(getPostValue), string(getSelectValue))
			if ReturnError500(w, errBdd) {
				return
			}
		} else if e == nil {
			errBdd = bdd.Like(IdToLike, userValue)
			if ReturnError500(w, errBdd) {
				return
			}
		} else if getSelectValue != "" {
			errBdd = bdd.MakeCategorie(string(getSelectValue))
			if ReturnError500(w, errBdd) {
				return
			}
		} else if eDislike == nil {
			bdd.Dislike(idToDislike, userValue)
		}
	}else {
		fmt.Println("User is not connected")
	}

	IdPostToComment := req.FormValue("ForId")
	fmt.Println(IdPostToComment)

	// Récupération de la valeur des commentaires
	getCommentValue := req.FormValue("commentaire")
	getCommentID := req.FormValue("IdComment")
	StrToInt, _ := strconv.Atoi(getCommentID)
	if getCommentValue != "" {

		errBdd = bdd.MakeComment(userValue, getCommentValue, StrToInt)
		if ReturnError500(w, errBdd) {
			return
		}

	}

	var arr [][]string
	var posts []PostData
	errBdd, arr = bdd.GetAllPoste()
	if ReturnError500(w, errBdd) {
		return
	}

	// Parcourir et remplir notre tableau des données que l'on veut
	for _, post := range arr {

		PostIdInt, _ := strconv.Atoi(post[0])
		errBdd, coms := bdd.GetCommentByPoste(PostIdInt)

		if ReturnError500(w, errBdd) {
			return
		}

		p := PostData{
			ID:       	post[0],
			UserName: 	post[1],
			Post:     	post[2],
			NbLike:   	bdd.GetLikeNb(PostIdInt),
			Date:     	post[5],
			CommentArr: coms,
			Categorie: post[3],
		}
		posts = append(posts, p)
	}

	// Gestion de l'erreur 404
	if req.URL.Path == "/" {
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	pageData := LoginWrapper {
		IsLogged: VerifyCookie(req),
		UserConnected: userValue, 
		Data: posts,

	}

	t.Execute(w, pageData)
}

func ReturnError500(w http.ResponseWriter, errBdd int) bool {
	if errBdd != 0 {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return true
	} else {
		return false
	}
}
