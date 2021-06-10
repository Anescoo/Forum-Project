package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	bdd "../bdd"
)

func Accueil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/Accueil.html", "./template/header.html")
	fmt.Print("Page d'accueil ✔️ \n")

	htmlData := text {
	post : "",
	}

	htmlData.post = bdd.GetPosteByID(1)


	if req.URL.Path == "/" { //verification de l'URL
	} else if req.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	t.Execute(w, nil)

}
