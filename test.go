package main 

import (
    "fmt"
    bdd "./src/bdd"
    // "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main(){

    _, db := bdd.OpenDataBase()

    // db.Exec("DELETE FROM Poste WHERE ID = 3")
    // bdd.MakeUser("Tao", "louis.teilliais@gmail.com", "Karim69lattrik")
	// bdd.MakePoste("Louis", "Karim est le meilleur", "test")
	// fmt.Println(bdd.GetPosteByID(2)[1])
    db.Exec("SELECT * FROM Poste ORDER BY DatePoste DESCENDING")
    fmt.Println(bdd.GetAllPoste())
    // fmt.Println(bdd.GetPosteByID(5))
    _, NbrPosts := bdd.GetAllPoste()
    fmt.Println(len(NbrPosts))

    // db, _ := sql.Open("sqlite3", "src/bdd/DataBase.db")
    // fmt.Println(db.Query("SELECT * FROM User"))
}