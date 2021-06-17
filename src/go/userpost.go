package handlers 

import (
	"fmt"
	"text/template"
	"net/http"
	
	bdd "../bdd"
)

func UserPost(w http.ResponseWriter, req *http.Request){
	
	fmt.Print("Page TL ✔️ \n")
	
	t, _ := template.ParseFiles("./template/userPost.html", "./template/header.html")
	
	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetPosteByUser("Tao")
	

	for _, post := range arr {
		p := PostData {
			Post: post[2],
			Date: post[5],
		}
		fmt.Println(p.Date)
		posts = append(posts, p)
	}

	t.Execute(w, posts)
}