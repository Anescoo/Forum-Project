package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Connexion(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/connexion.html", "./template/header.html")
	fmt.Print("Page de connexion ✔️ \n")
	t.Execute(w, nil)
}
