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
	
	
	// postdata := PostData
	
	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetPosteByUser("Tao")
	
	// fmt.Println(arr)
	for _, post := range arr {
		p := PostData {
			Post: post[2],
			// Date: post[5],
		}
		posts = append(posts, p)
	}
	// fmt.Println(posts)
	// fmt.Println(postdata)
	// fmt.Println(p.Post)

	t.Execute(w, posts)
}