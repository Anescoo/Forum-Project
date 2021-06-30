package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

type HTMLData struct {
	pseudoWrong  string
	mdpWrong     string
	mdpWeak      string
	emailInvalid string
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
	err := Register(getPseudoInscription, getEmailInscription, getMdpInscription)

	ErrorHtml := HTMLData{
		pseudoWrong:  "",
		mdpWrong:     "",
		mdpWeak:      "",
		emailInvalid: "",
	}

	// Gestion des erreurs (pas encore faite)
	if err != 0 {
		if err == 1 {
			ErrorHtml.pseudoWrong = "Votre pseudo est trop court"
			fmt.Println("Votre pseudo est trop court")

		} else if err == 2 {
			ErrorHtml.mdpWrong = "Votre mot de passe est trop court"
			fmt.Println("Votre mot de passe est trop court")
			
		} else if err == 3 {
			ErrorHtml.mdpWeak = "Votre mot de passe est trop faible"
			fmt.Println("Votre mot de passe est trop faible")

		} else if err == 4 {
			ErrorHtml.emailInvalid = "Adresse mail invalide"
			fmt.Println("Adresse mail invalide")
		}
	}
	t.Execute(w, ErrorHtml)
}
