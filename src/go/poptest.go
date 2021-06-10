package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Poptest(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("./template/poptest.html")
	fmt.Print("Poster une publication✔️ \n")

	t.Execute(w, nil)
}
