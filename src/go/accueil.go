package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	bdd "../bdd"
)

type test struct {
	UserName string
	Post string
	DatePoste string
}


func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/header.html")
	fmt.Print("Page d'accueil ✔️ \n")

	getPostValue := req.FormValue("PostValue")
	fmt.Println(getPostValue)
	bdd.MakeUser("Tao", "louis.teilliais@gmail.com", "Karim69lattrik")
	bdd.MakePoste("Tao",getPostValue,"test")
	var arr []string
	_, arr = bdd.GetPosteByID(5)
	fmt.Println(arr)
	
	p := test {
		UserName: arr[1],
		Post: arr[2],
		DatePoste : arr[5],
	}

	fmt.Println(p.DatePoste)
	
	for (i := 0; i < len()); i++){
		fmt.Println(p.Post)
	}


	if req.URL.Path == "/" { //verification de l'URL
	} else if req.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	t.Execute(w, p)
}
