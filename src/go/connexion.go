package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	bdd "../bdd"
)

type DataError struct {
	MessageError string
	NbErr int
}


func Connexion(w http.ResponseWriter, req *http.Request) {

	t, errFiles := template.ParseFiles("./template/connexion.html", "./template/Header.html")

	if errFiles != nil {
		fmt.Println(errFiles.Error())
	}

	var errBdd int

	fmt.Print("Page de connexion ✔️ \n")

	// on récupère les valeurs (pseudo et mdp) que l'utilisateur rentre
	getPseudo := req.FormValue("pseudoConnexion")
	getMdp := req.FormValue("mdpConnexion")

	fmt.Println("Pseudo : ", getPseudo)
	fmt.Println("Mot de Passe :", getMdp)

	err := Login(w, getPseudo, getMdp)

	DataHTML := DataError {
		MessageError: "",
		NbErr: err,
	}

	// Gestion d'erreurs si l'utilisateur se trompe.
	if err != 0 {
		if err == 1 {
			DataHTML.MessageError = "Une erreur s'est produite, veuillez réessayer"
		}else if err == 2 {
			DataHTML.MessageError = "Mot de passe incorrect"
		}else if err == 3 {
			DataHTML.MessageError = "Pseudo invalide"
		}
	} else {
		uuidValue := sessionCookie(w, req)
		errBdd = bdd.AddSession(getPseudo, uuidValue)
		ReturnError500(w, errBdd)
		time.Sleep(1 * time.Second)
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	}

	t.Execute(w, DataHTML)
}
