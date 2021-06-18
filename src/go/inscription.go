package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	authent "./authent"
)

func Inscription(w http.ResponseWriter, req *http.Request) {

	fmt.Print("Page d'inscription ✔️ \n")

	t, _ := template.ParseFiles("./template/inscription.html", "./template/header.html")

	getPseudoInscription := req.FormValue("pseudoInscription")
	getEmailInscription := req.FormValue("emailInscription")
	getMdpInscription := req.FormValue("mdpInscription")

	fmt.Println("Pseudo : ", getPseudoInscription)
	fmt.Println("E-mail : ", getEmailInscription)
	fmt.Println("Mot de passe : ", getMdpInscription)

	err := authent.Register(getPseudoInscription, getEmailInscription, getMdpInscription)

	if err != 0 {
		if err == 1 {
			//mauvaise longueur pseudo
		}
		if err == 2 {
			//mot de passe trop court
		}
		if err == 3 {
			//mot de passe trop faible
		}
		if err == 4 {
			//adresse email invalide
		}
		if err == 5 {
			//username déja utilisé
		}
		if err == 6 {
			//email déja utilisé
		}
	}
	type HTMLData struct {
		pseudoWrong  string
		mdpWrong     string
		mdpWeak      string
		emailInvalid string
	}

	htmlData := HTMLData{
		pseudoWrong:  "",
		mdpWrong:     "",
		mdpWeak:      "",
		emailInvalid: "",
	}

	if err != 0 {
		if err == 1 {

			htmlData.pseudoWrong = "Votre pseudo est trop court"
			fmt.Println("Votre pseudo est trop court")
			//mauvaise longueur pseudo

		} else if err == 2 {
			htmlData.mdpWrong = "Votre mot de passe est trop court"
			fmt.Println("Votre mot de passe est trop court")
			//mot de passe trop court

		} else if err == 3 {
			htmlData.mdpWeak = "Votre mot de passe est trop faible"
			fmt.Println("Votre mot de passe est trop faible")
			//mot de passe trop faible

		} else if err == 4 {
			htmlData.emailInvalid = "Adresse mail invalide"
			fmt.Println("Adresse mail invalide")
			//adresse email invalide
		}
	}

	t.Execute(w, nil)
}
