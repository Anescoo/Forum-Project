package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)



func Profil(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/profil.html", "./template/header.html")
	fmt.Print("Page du profil ✔️ \n")
		
	
	t.Execute(w, nil)
}
