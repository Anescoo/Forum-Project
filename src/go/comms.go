package handlers

import (
    "fmt"
    "net/http"
    "text/template"

    // bdd "../bdd"

)



func Comments(w http.ResponseWriter, req *http.Request) {

    t, _ := template.ParseFiles("./template/comms.html", "./template/header.html")
    fmt.Print("Page des commentaires post ✔️ \n")
    
    // var arr [][]string
	// var posts []PostData
    // _, test := bdd.GetPosteByID(1)
    
    // fmt.Println(IdToComment)
    
    t.Execute(w, nil)
}