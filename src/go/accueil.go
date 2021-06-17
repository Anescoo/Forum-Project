package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"strconv"
	bdd "../bdd"
)

type StructOfStruct struct {
	
}

type PostData struct {
	UserName string
	Post string
	Date string
	ID string
}



func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/header.html")
	fmt.Print("Page d'accueil ✔️ \n")

	getPostID := req.FormValue("test")
	IdToSuppr, err:= strconv.Atoi(getPostID)
	if err == nil {
		bdd.DeletePoste(IdToSuppr)
	}
	getPostValue := req.FormValue("PostValue")
	
	if getPostValue != "" {
		bdd.MakePoste("Louis", string(getPostValue),"test")		
	}

	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetAllPoste()
	for _, post := range arr {
		p := PostData {
			ID: post[0],
			UserName: post[1],
			Post: post[2],
			Date: post[5],
		}
		posts = append(posts, p)
	}
	
	if req.URL.Path == "/" { //verification de l'URL
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	t.Execute(w, posts)
}
