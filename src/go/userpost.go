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
	
	t, _ := template.ParseFiles("./template/userPost.html", "./template/header.html")
	
	getPostID := req.FormValue("test")
	IdToSuppr, err:= strconv.Atoi(getPostID)
	if err == nil {
		bdd.DeletePoste(IdToSuppr)
	}

	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetPosteByUser("Louis")
	

	for _, post := range arr {
		p := PostData {
			ID: post[0],
			Post: post[2],
		}
		posts = append(posts, p)
	}
	

	t.Execute(w, posts)
}