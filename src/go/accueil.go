package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
<<<<<<< HEAD
=======
	// "time"
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285

	bdd "../bdd"
)

<<<<<<< HEAD
// Création d'une structure avec toutes les "infos" de notre post.
type PostData struct {
=======
// Création d'une structure avec toutes les "infos" de notre post. 
type PostData struct { 
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
	UserName string
	Post     string
	Date     string
	NbLike   int
	ID       string
	// Categorie []string
}

<<<<<<< HEAD
func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/header.html")
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
			bdd.MakePoste("Tao", string(getPostValue), string(getSelectValue))
=======
type LoginWrapper struct {
	IsLogged bool 
	Data interface{} 
}

func Accueil(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./template/Accueil.html", "./template/header.html")
	
	if err != nil {
		fmt.Print(err.Error)
	}

	fmt.Print("Page d'accueil ✔️ \n")

	getPostValue := req.FormValue("PostValue")
	getSelectValue := req.FormValue("selectCategorie")
	getIDLike := req.FormValue("like") 
	IdToLike, e := strconv.Atoi(getIDLike)
	getCategorieValue := req.FormValue("categorie")

	
	// Vérification si l'utilisateur est connecté
	if VerifyCookie(req) == true{
		if getPostValue != "" {
			bdd.MakePoste("Tao", string(getPostValue), string(getSelectValue)) 	
		}else if e == nil{
			bdd.Like(IdToLike, "Louis") 
		}else if getCategorieValue != ""{
			bdd.MakeCategorie(string(getCategorieValue))
		}else if e == nil {
			bdd.Dislike(IdToLike, "Louis")
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
		}
	}
	// else {
	// 	time.Sleep(5 * time.Second)
	// 	http.Redirect(w, req, "/connexion", http.StatusSeeOther)
	// }
<<<<<<< HEAD

	// Liker un post
	getIDLike := req.FormValue("like")
	IdToLike, e := strconv.Atoi(getIDLike)
	if e == nil {
		bdd.Like(IdToLike, "Louis")
	}

	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetAllPoste()

	// Parcourir et remplir notre tableau des données que l'on veut
	for _, post := range arr {
=======
	
	var arr [][]string 
	var posts []PostData 
	_, arr = bdd.GetAllPoste() 
	
	// Parcourir et remplir notre tableau des données que l'on veut  
	for _, post := range arr { 
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
		nbLike, _ := strconv.Atoi(post[0])
		p := PostData{
			ID:       post[0],
			UserName: post[1],
			Post:     post[2],
			NbLike:   bdd.GetLikeNb(nbLike),
			Date:     post[5],
<<<<<<< HEAD
=======
			// Categorie:post[3],
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
		}
		posts = append(posts, p)
	}


	uuidValue := readCookie(w, req)
	_, userValue := bdd.GetUserByUUID(uuidValue)
	fmt.Println(userValue)

	// Gestion de l'erreur 404
	if req.URL.Path == "/" {
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	pageData := LoginWrapper {
		IsLogged: VerifyCookie(req),
		Data: posts,
	}

	t.Execute(w, pageData)
}
