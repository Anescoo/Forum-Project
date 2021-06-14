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

	t.Execute(w, nil)

}
