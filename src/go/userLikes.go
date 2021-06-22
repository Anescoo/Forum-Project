package handlers

import (
	"fmt"
	"text/template"
	"net/http"
	
	bdd "../bdd"
)

func UserLikes(w http.ResponseWriter, req *http.Request){
	
	fmt.Print("Page Mes Posts ✔️ \n")
	
	t, _ := template.ParseFiles("./template/userLikes.html", "./template/header.html")

	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetPosteLikeByUser("Tao")

	for _, post := range arr {
		p := PostData {
			ID: post[0],
			UserName: post[1],
			Post: post[2],
			Date: post[5],
		}
		// fmt.Println(p.Date)
		posts = append(posts, p)
	}


	t.Execute(w, posts)
}