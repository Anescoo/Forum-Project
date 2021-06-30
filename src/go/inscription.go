package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type ErrorData struct {
	Message string
	NbErr int
}

func Inscription(w http.ResponseWriter, req *http.Request) {

	fmt.Print("Page d'inscription ✔️ \n")

	t, errFiles := template.ParseFiles("./template/inscription.html", "./template/Header.html")

	if errFiles != nil {
		fmt.Println(errFiles.Error())
	}

	// Idem que connextion on récupère les valeurs que l'utilisateurs rentre
	getPseudoInscription := req.FormValue("pseudoInscription")
	getEmailInscription := req.FormValue("emailInscription")
	getMdpInscription := req.FormValue("mdpInscription")

	fmt.Println("Pseudo : ", getPseudoInscription)
	fmt.Println("E-mail : ", getEmailInscription)
	fmt.Println("Mot de passe : ", getMdpInscription)

	// Application de la fonction Register() pour vérifier la validité des valeurs rentrés.
	err := Register(getPseudoInscription, getEmailInscription, getMdpInscription, w)

	HTMLError := ErrorData{
		Message: "",
		NbErr: err,
	}

	// Gestion des erreurs (pas encore faite)
	if err != 0 {
		if err == 1 {
			HTMLError.Message = "Votre pseudo est trop court"
		} else if err == 2 {
			HTMLError.Message = "Votre mot de passe est trop court"

		} else if err == 3 {
			HTMLError.Message = "Votre mot de passe est trop faible"

		} else if err == 4 {
			HTMLError.Message = "Adresse mail invalide"
		}
		fmt.Println(HTMLError.Message)
	}
	t.Execute(w, HTMLError)
}
