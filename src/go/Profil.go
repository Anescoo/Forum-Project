package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Profil(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("./template/profil.html", "./template/header.html")
	
	if err != nil {
		fmt.Print(err.Error)
	}


	fmt.Print("Page du profil ✔️ \n")

	
	t.Execute(w, nil)
}
