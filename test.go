package main 

import (
    // "fmt"
    bdd "./src/bdd"
    // "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main(){

    // _, db := bdd.OpenDataBase()

    // db.Exec("DELETE FROM Poste WHERE ID = 3")
    // bdd.MakeUser("Louis", "louis.teilliais@gmail.com", "Karim69lattrik")
	// bdd.MakePoste("Louis", "Karim est le meilleur", "test")
	// // fmt.Println(bdd.GetPosteByID(2)[1])
    // db.Exec("SELECT * FROM Poste ORDER BY PosteTime ASC")
    // fmt.Println(bdd.GetAllPoste())
    // fmt.Println(bdd.GetPosteByID(5))
    // _, NbrPosts := bdd.GetAllPoste()
    // fmt.Println(len(NbrPosts))

    // db, _ := sql.Open("sqlite3", "src/bdd/DataBase.db")
    // fmt.Println(db.Query("SELECT * FROM User"))

    // fmt.Println(bdd.GetPosteByUser("Tao"))

    // bdd.UpdatePoste(15, "Allez les bleus")
    // bdd.Like(13, "Louis")
    // bdd.GetLikeNb(13)
    // fmt.Println(bdd.IsLike(13, "Louis"))
    // bdd.MakeCategorie("Voiture")
    // fmt.Println(bdd.GetPosteLikeByUser("Louis"))
    // fmt.Println(bdd.GetAllCategorie())
    // bdd.DeleteCategorire("")
    // bdd.MakeComment("Louis", "Je suis un commentaire", 13)
    // bdd.MakeUser("test", "test@gmail.com", "AZERTY")
    bdd.MakePoste("Louis", "test", "testcategorie")
}
