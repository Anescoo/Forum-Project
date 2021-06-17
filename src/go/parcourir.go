package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Parcrir(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/parcourir.html")
	fmt.Print("Cherchez dans l'arborescence de votre disque✔️ \n")

	t.Execute(w, nil)
}
