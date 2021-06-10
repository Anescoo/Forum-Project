package main 

import (
    "fmt"
    bdd "./src/bdd"
    // "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main(){

    
    bdd.MakeUser("Tao", "louis.teilliais@gmail.com", "Karim69lattrik")
	bdd.MakePoste("Tao", "Ce post s'affiche dans la console", "test")
	fmt.Println(bdd.GetPosteByID(2)[1])


    // db, _ := sql.Open("sqlite3", "src/bdd/DataBase.db")
    // fmt.Println(db.Query("SELECT * FROM User"))
}