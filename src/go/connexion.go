package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Connexion(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/connexion.html", "./template/header.html")
	fmt.Print("Page de connexion ✔️ \n")

	// on récupère les valeurs (pseudo et mdp) que l'utilisateur rentre
	getPseudo := req.FormValue("pseudoConnexion")
	getMdp := req.FormValue("mdpConnexion")

	fmt.Println("Pseudo : ", getPseudo)
	fmt.Println("Mot de Passe :", getMdp)

	err := Login(w, getPseudo, getMdp)
	// if Login(w, getPseudo, getMdp) == 2 {
	// 	sessionCookie(w, req)
	// }

	// Gestion d'erreurs si l'utilisateur se trompe.
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

	// quand l'utilisateur se connecte

	// sessionCookie(w, req)

	t.Execute(w, nil)
}
