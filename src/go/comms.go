package handlers

import (
    "fmt"
    "net/http"
    "text/template"

    // bdd "../bdd"

)



func Comments(w http.ResponseWriter, req *http.Request) {

    t, errFiles := template.ParseFiles("./template/comms.html", "./template/header.html")
    
    if errFiles != nil {
		fmt.Print(errFiles.Error)
	}

    fmt.Print("Page des commentaires post ✔️ \n")
    
    // var arr [][]string
	// var posts []PostData
    // _, test := bdd.GetPosteByID(1)
    
    // fmt.Println(IdToComment)
    
    t.Execute(w, nil)
}