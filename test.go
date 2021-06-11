package main 

import (
    "fmt"
    bdd "./src/bdd"
    "time"
    // "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main(){

    currentTime := time.Now()
    fmt.Println("Current Time in String: ", currentTime.String())
    // _, db := bdd.OpenDataBase()

    // db.Exec("DELETE FROM Poste WHERE ID = 23")
    // bdd.MakeUser("Tao", "louis.teilliais@gmail.com", "Karim69lattrik")
	// bdd.MakePoste("Tao", "Ce post s'affiche dans la console", "test")
	// fmt.Println(bdd.GetPosteByID(2)[1])
    fmt.Println(bdd.GetAllPoste())
    // fmt.Println(bdd.GetPosteByID(23))
    _, NbrPosts := bdd.GetAllPoste()
    fmt.Println(len(NbrPosts))
    // db, _ := sql.Open("sqlite3", "src/bdd/DataBase.db")
    // fmt.Println(db.Query("SELECT * FROM User"))
}
