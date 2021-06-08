<<<<<<< HEAD
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
	}

	t.Execute(w, nil)
}
=======
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
>>>>>>> 54a15459cdb70964db55023c3d0e1fce5387cfdb
