package handlers 

import (
	"fmt"
	"net/http"
	"text/template"
)

func Connexion(w http.ResponseWriter, req *http.Request){

	t,_ := template.ParseFiles("./template/connexion.html", "./template/Header.html")
	fmt.Print("Page de connexion ✔️ \n")

	req.ParseForm()

	
	getPseudo := req.FormValue("pseudo")
	getMdp := req.FormValue("mdp")

	fmt.Println("Pseudo : " , getPseudo)
	fmt.Println("Mot de Passe :", getMdp)

	t.Execute(w, nil)
}

