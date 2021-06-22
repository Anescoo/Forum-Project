package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	authent "./authent"
)

func Connexion(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/connexion.html", "./template/header.html")
	fmt.Print("Page de connexion ✔️ \n")

	getPseudo := req.FormValue("pseudoConnexion") // on récupère les valeurs (pseudo et mdp) que l'utilisateur rentre 
	getMdp := req.FormValue("mdpConnexion")

	fmt.Println("Pseudo : ", getPseudo)
	fmt.Println("Mot de Passe :", getMdp)
	err := authent.Login(nil, getPseudo, getMdp) // Gestion d'erreur si l'utilisateur se trompe.
	if err != 0 {
		if err == 1 {
			//erreur de génération de token
		}
		if err == 2 {
			//mdp incorect
		}
		if err == 3 {
			//username invalide
		}
	}
	t.Execute(w, nil)
}
