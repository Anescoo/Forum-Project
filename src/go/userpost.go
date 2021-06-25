package handlers 

import (
	"fmt"
	"text/template"
	"net/http"
	"strconv"
	
	bdd "../bdd"
)

func UserPost(w http.ResponseWriter, req *http.Request){
	
	fmt.Print("Page TL ✔️ \n")
	
	t, errFiles := template.ParseFiles("./template/userPost.html", "./template/header.html")
	
	if errFiles != nil {
		fmt.Print(errFiles.Error)
	}
	
	// Gestion de la supression des posts 
	getPostID := req.FormValue("delete")
	IdToSuppr, err:= strconv.Atoi(getPostID)
	getNewValue := req.FormValue("sendUpdate")
	getPostIDupdate := req.FormValue("update") 
	IdtoUpdate, errUpdate:= strconv.Atoi(getPostIDupdate)
	
	if VerifyCookie(w ,req) == true {
		if err == nil {
			bdd.DeletePoste(IdToSuppr)
		}else if errUpdate == nil {
			bdd.UpdatePoste(IdtoUpdate, getNewValue)
		}
	}
	
	
	// même methode que dans "Accueil.go" mais pour les posts de l'utilisateurs on prend juste les valeurs dont on a besoin  
	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetPosteByUser("Tao")
	
	
	for _, post := range arr {
		p := PostData {
			ID: post[0],
			Post: post[2],
			Date: post[5],
		}
		posts = append(posts, p)
	}
	
	t.Execute(w, posts)
}