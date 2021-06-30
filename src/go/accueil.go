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
	// Categorie []string
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
	getCategorieValue := req.FormValue("categorie")

	getIDLike := req.FormValue("like")
	IdToLike, e := strconv.Atoi(getIDLike)

	getDislike := req.FormValue("Dislike")
	idToDislike, eDislike := strconv.Atoi(getDislike)

	uuidValue := readCookie(w, req)
	var userValue string
	errBdd, userValue = bdd.GetUserByUUID(uuidValue)
	ReturnError500(w, errBdd)

	// Vérification si l'utilisateur est connecté
	if VerifyCookie(req) == true {
		if getPostValue != "" {

			errBdd = bdd.MakePoste(userValue, string(getPostValue), string(getSelectValue))
			ReturnError500(w, errBdd)
		} else if e == nil {
			errBdd = bdd.Like(IdToLike, userValue)
			ReturnError500(w, errBdd)
			fmt.Println(userValue + " a Liker")
		} else if getCategorieValue != "" {
			errBdd = bdd.MakeCategorie(string(getCategorieValue))
			ReturnError500(w, errBdd)
		} else if eDislike == nil {
			bdd.Dislike(idToDislike, userValue)
		}
	}
	// else {
	// 		time.Sleep(5 * time.Second)
	// 		http.Redirect(w, req, "/connexion", http.StatusSeeOther)
	// 	}

	IdPostToComment := req.FormValue("ForId")
	fmt.Println(IdPostToComment)

	// Récupération de la valeur des commentaires
	getCommentValue := req.FormValue("commentaire")
	getCommentID := req.FormValue("IdComment")
	StrToInt, _ := strconv.Atoi(getCommentID)
	if getCommentValue != "" {

		fmt.Println(StrToInt)
		bdd.MakeComment(userValue, getCommentValue, StrToInt)

		ReturnError500(w, errBdd)

	}

	var arr [][]string
	var posts []PostData
	errBdd, arr = bdd.GetAllPoste()
	ReturnError500(w, errBdd)

	// Parcourir et remplir notre tableau des données que l'on veut
	for _, post := range arr {

		PostIdInt, _ := strconv.Atoi(post[0])
		_, coms := bdd.GetCommentByPoste(PostIdInt)

		errBdd, coms := bdd.GetCommentByPoste(commentID)
		ReturnError500(w, errBdd)

		p := PostData{
			ID:       	post[0],
			UserName: 	post[1],
			Post:     	post[2],
			NbLike:   	bdd.GetLikeNb(PostIdInt),
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

	pageData := LoginWrapper{
		IsLogged:      VerifyCookie(req),
		UserConnected: userValue,
		Data:          posts,
	}

	t.Execute(w, pageData)
}

func ReturnError500(w http.ResponseWriter, errBdd int) {
	if errBdd != 0 {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	} else {
		return
	}
}
