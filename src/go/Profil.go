package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Profil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/profil.html", "./template/Header.html")
	fmt.Print("Page du compte utilisateur ✔️ \n")
	t.Execute(w, nil)
}
