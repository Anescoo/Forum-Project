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
	UserName string
	Post     string
	Date     string
	NbLike   int
	ID       string
	// Categorie []string
}

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
		}
	}
	// else {
	// 	time.Sleep(5 * time.Second)
	// 	http.Redirect(w, req, "/connexion", http.StatusSeeOther)
	// }
	
	var arr [][]string 
	var posts []PostData 
	_, arr = bdd.GetAllPoste() 
	
	// Parcourir et remplir notre tableau des données que l'on veut  
	for _, post := range arr { 
		nbLike, _ := strconv.Atoi(post[0])
		p := PostData{
			ID:       post[0],
			UserName: post[1],
			Post:     post[2],
			NbLike:   bdd.GetLikeNb(nbLike),
			Date:     post[5],
			// Categorie:post[3],
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
