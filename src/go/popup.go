package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func POP(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/popup.html")
	fmt.Print("Page d'accueil ✔️ \n")

	if req.URL.Path == "/popup" { //verification de l'URL
	} else if req.URL.Path != "/popup" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	t.Execute(w, nil)
}
