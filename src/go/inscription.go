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

    getPseudoInscription := req.FormValue("pseudoInscription") // Idem que connextion on récupère les valeurs que l'utilisateurs rentre 
    getEmailInscription := req.FormValue("emailInscription")
    getMdpInscription := req.FormValue("mdpInscription")

    fmt.Println("Pseudo : ", getPseudoInscription) // On les imprime dans la console 
    fmt.Println("E-mail : ", getEmailInscription)
    fmt.Println("Mot de passe : ", getMdpInscription)

    err := authent.Register(getPseudoInscription, getEmailInscription, getMdpInscription) // on lui applique la fonction Register() pour vérifier la validité des valeurs rentrés. 

    type HTMLData struct {
        pseudoWrong string
        mdpWrong string
        mdpWeak string
        emailInvalid string
    }

    htmlData := HTMLData{
        pseudoWrong : "",
        mdpWrong : "",
        mdpWeak : "",
        emailInvalid: "",
    }

    // Gestion des erreurs (pas encore faite)
    if err != 0 { 
        if err == 1 { 
            
            htmlData.pseudoWrong = "Votre pseudo est trop court"
            fmt.Println("Votre pseudo est trop court")
            //mauvaise longueur pseudo
        
        }else if err == 2 {
            htmlData.mdpWrong = "Votre mot de passe est trop court"
            fmt.Println("Votre mot de passe est trop court")
            //mot de passe trop court
        
        }else if err == 3 {
            htmlData.mdpWeak = "Votre mot de passe est trop faible"
            fmt.Println("Votre mot de passe est trop faible")
            //mot de passe trop faible
        
        }else if err == 4 {
            htmlData.emailInvalid = "Adresse mail invalide"
            fmt.Println("Adresse mail invalide")
            //adresse email invalide
        }
    }
    
    t.Execute(w, nil)
}
