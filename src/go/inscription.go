package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Inscription(w http.ResponseWriter, req *http.Request) {

	fmt.Print("Page d'inscription ✔️ \n")

	t, _ := template.ParseFiles("./template/inscription.html", "./template/header.html")

	t.Execute(w, nil)

}
