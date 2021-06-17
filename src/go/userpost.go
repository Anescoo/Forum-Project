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
	
	getPostID := req.FormValue("delete")
	IdToSuppr, err:= strconv.Atoi(getPostID)
	if err == nil {
		bdd.DeletePoste(IdToSuppr)
	}


	getNewValue := req.FormValue("sendUpdate")
	getPostIDupdate := req.FormValue("update")
	IdtoUpdate, err:= strconv.Atoi(getPostIDupdate)
	if err == nil {
		bdd.UpdatePoste(IdtoUpdate, getNewValue)
	}


	// fmt.Println(getNewValue)
	// fmt.Println(bdd.GetPosteByID(IdToSuppr))
	// fmt.Println(bdd.UpdatePoste(IdToSuppr, getNewValue))

	var arr [][]string
	var posts []PostData
	_, arr = bdd.GetPosteByUser("Tao")
	

	for _, post := range arr {
		p := PostData {
			ID: post[0],
			Post: post[2],
			Date: post[5],
		}
		// fmt.Println(p.Date)
		posts = append(posts, p)
	}

	t.Execute(w, posts)
}