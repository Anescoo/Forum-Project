package handlers

import (
    "fmt"
    "net/http"
    "strconv"
    "text/template"

    bdd "../bdd"
)

func UserPost(w http.ResponseWriter, req *http.Request) {

	fmt.Print("Page TL ✔️ \n")

	t, errFiles := template.ParseFiles("./template/userPost.html", "./template/Header.html")

	if errFiles != nil {
		fmt.Print(errFiles.Error())
	}

	var errBdd int

	// Gestion de la supression des posts
	getPostID := req.FormValue("delete")
	IdToSuppr, err := strconv.Atoi(getPostID)
	getNewValue := req.FormValue("sendUpdate")
	getPostIDupdate := req.FormValue("update")
	IdtoUpdate, errUpdate := strconv.Atoi(getPostIDupdate)

	// Vérification de l'user
	uuidValue := readCookie(w, req)
	errBdd, userValue := bdd.GetUserByUUID(uuidValue)
	if ReturnError500(w, errBdd) {
		return
	}

	if VerifyCookie(req) {
		if err == nil {
			errBdd = bdd.DeletePoste(IdToSuppr)
			if ReturnError500(w, errBdd) {
				return
			}
		} else if errUpdate == nil {
			errBdd = bdd.UpdatePoste(IdtoUpdate, getNewValue)
			if ReturnError500(w, errBdd) {
				return
			}
		}
	}

	// même methode que dans "Accueil.go" mais pour les posts de l'utilisateurs on prend juste les valeurs dont on a besoin
	var arr [][]string
	var posts []PostData

	errBdd, arr = bdd.GetPosteByUser(userValue)
	if ReturnError500(w, errBdd) {
		return
	}

	for _, post := range arr {
		p := PostData{
			ID:   post[0],
			Post: post[2],
			Date: post[5],
			Categorie: post[3],
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
