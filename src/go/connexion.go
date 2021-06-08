package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	bdd "../bdd"
)

func Connexion(w http.ResponseWriter, req *http.Request) {
	

	t, _ := template.ParseFiles("./template/connexion.html", "./template/header.html")
	fmt.Print("Page de connexion ✔️ \n")

	getPseudo := req.FormValue("pseudoConnexion")
	getMdp := req.FormValue("mdpConnexion")

	fmt.Println("Pseudo : ", getPseudo)
	fmt.Println("Mot de Passe :", getMdp)

	bdd.GetUser(getPseudo)
	bdd.GetUserHash(getPseudo)
	
	t.Execute(w, nil)
}
