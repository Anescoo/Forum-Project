package main

import (
	"fmt"
	"net/http"

	handlers "./src/go"
)

func main() {

	fmt.Print("Démarrage du serveur... 💬\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/timeline", handlers.Timeline)
<<<<<<< HEAD
	http.HandleFunc("/profil", handlers.Profil)
	http.HandleFunc("/inscription", handlers.Inscription)
=======
	http.HandleFunc("/inscription", handlers.Inscription)

	http.HandleFunc("/profil", handlers.Profil)
>>>>>>> c29b4ba1f50299bc288a20b1fbae7577639c3cac

	http.ListenAndServe(":8000", nil)
}
