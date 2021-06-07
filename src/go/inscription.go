package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	bdd "../bdd"
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
	
	bdd.MakeUser(getPseudoInscription, getEmailInscription, getMdpInscription)
	

	t.Execute(w, nil)

}
