package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
	bdd "../bdd"
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
<<<<<<< HEAD
	// if Login(w, getPseudo, getMdp) == 2 {
	// 	sessionCookie(w, req)
	// }
=======
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285

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
		uuidValue := sessionCookie(w, req)
		bdd.AddSession(getPseudo, uuidValue)
		time.Sleep(2 * time.Second)
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	}
<<<<<<< HEAD

	// quand l'utilisateur se connecte

	// sessionCookie(w, req)

=======
			
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
	t.Execute(w, nil)
}
