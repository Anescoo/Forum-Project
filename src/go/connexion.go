package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	
)

func Connexion(w http.ResponseWriter, req *http.Request) {
	
	t, errFiles := template.ParseFiles("./template/connexion.html", "./template/header.html")

	if errFiles != nil {
		fmt.Print(errFiles.Error)
	}

	fmt.Print("Page de connexion ✔️ \n")
	
	
	// on récupère les valeurs (pseudo et mdp) que l'utilisateur rentre
	getPseudo := req.FormValue("pseudoConnexion")  
	getMdp := req.FormValue("mdpConnexion")
	
	fmt.Println("Pseudo : ", getPseudo)
	fmt.Println("Mot de Passe :", getMdp)
	
	
	err := Login(w, getPseudo, getMdp)

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

	}else {
		sessionCookie(w, req)
		// AJOUTER UUID A LA TABLE SESSION
	}
		
	t.Execute(w, nil)
}
