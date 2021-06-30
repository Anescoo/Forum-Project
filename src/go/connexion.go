package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	bdd "../bdd"
)

type DataError struct {
	PasswordError string
	TokenError string 
	UserNameError string
}


func Connexion(w http.ResponseWriter, req *http.Request) {

	t, errFiles := template.ParseFiles("./template/connexion.html", "./template/Header.html")

	if errFiles != nil {
		fmt.Println(errFiles.Error())
	}

	fmt.Print("Page de connexion ✔️ \n")

	// on récupère les valeurs (pseudo et mdp) que l'utilisateur rentre
	getPseudo := req.FormValue("pseudoConnexion")
	getMdp := req.FormValue("mdpConnexion")

	fmt.Println("Pseudo : ", getPseudo)
	fmt.Println("Mot de Passe :", getMdp)

	err := Login(w, getPseudo, getMdp)

	DataHTML := DataError {
		TokenError: "",
		PasswordError: "",
		UserNameError: "",
	}
	// Gestion d'erreurs si l'utilisateur se trompe.
	if err != 0 {
		if err == 1 {
			// DataHTML.TokenError = "Une erreur s'est produite, veuillez réessayer"
			// fmt.Println("Une erreur s'est produite")
		}else if err == 2 {
			// DataHTML.PasswordError = "Mot de passe incorrect"
			// fmt.Println("Mot de passe incorect")
		}else if err == 3 {
			// DataHTML.UserNameError = "Pseudo invalide"
			// fmt.Println("Pseudo invalide")
		}
	} else {
		uuidValue := sessionCookie(w, req)
		bdd.AddSession(getPseudo, uuidValue)
		time.Sleep(1 * time.Second)
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	}

	t.Execute(w, DataHTML)
}
