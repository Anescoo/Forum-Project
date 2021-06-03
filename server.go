package main 


import (
	"net/http"
	"fmt"
	handlers "./src/go"
)

func main(){

	fmt.Print("Démarrage du serveur... 💬\n")

	http.Handle("/css/static/", http.StripPrefix("/css/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/timeline", handlers.Timeline)


	http.ListenAndServe(":8000", nil)
}