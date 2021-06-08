<<<<<<< HEAD
package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	bdd "../bdd"
)

func Connexion(w http.ResponseWriter, req *http.Request) {
	

	t, _ := template.ParseFiles("./template/connexion.html", "./template/header.html")
	fmt.Print("Page de connexion ✔️ \n")

	getPseudo := req.FormValue("pseudoConnexion")
	getMdp := req.FormValue("mdpConnexion")

	fmt.Println("Pseudo : " , getPseudo)
	fmt.Println("Mot de Passe :", getMdp)

	bdd.GetUser(getPseudo)
	bdd.GetUserHash(getPseudo)
	
	t.Execute(w, nil)
}
=======
package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	authent "./authent"
)

func Connexion(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/connexion.html", "./template/header.html")
	fmt.Print("Page de connexion ✔️ \n")

	getPseudo := req.FormValue("pseudoConnexion")
	getMdp := req.FormValue("mdpConnexion")

	fmt.Println("Pseudo : ", getPseudo)
	fmt.Println("Mot de Passe :", getMdp)
	err := authent.Login(nil, getPseudo, getMdp)
	if err != 0 {
		if err == 1 {
			//erreur de génération de token
		}
		if err == 2 {
			//mdp incorect
		}
	}
	t.Execute(w, nil)
}
>>>>>>> b627772974ef584a0276f9d77f4501ba12b128e2
