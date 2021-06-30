package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	bdd "../bdd"
)

func Profil(w http.ResponseWriter, req *http.Request) {
	
	fmt.Print("Page du profil ✔️ \n")
	
	t, err := template.ParseFiles("./template/profil.html", "./template/header.html")
	
	if err != nil {
		fmt.Print(err.Error)
	}
	uuidValue := readCookie(w, req)
	_, userValue := bdd.GetUserByUUID(uuidValue)

	pageData := LoginWrapper {
		IsLogged: VerifyCookie(req),
		UserConnected: userValue, 	
	}

	t.Execute(w, pageData)
}
