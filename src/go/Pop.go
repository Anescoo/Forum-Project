package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Pop(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/pop-up.html")
	fmt.Print("Poster une publication✔️ \n")

	t.Execute(w, nil)
}
