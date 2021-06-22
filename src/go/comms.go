package handlers

import (
    "fmt"
    "net/http"
    "text/template"
)



func Comments(w http.ResponseWriter, req *http.Request) {

    t, _ := template.ParseFiles("./template/comms.html", "./template/header.html")
    fmt.Print("Page des commentaires post ✔️ \n")
    
	
    
    t.Execute(w, nil)
}