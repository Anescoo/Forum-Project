package main 

import (
    "fmt"
    bdd "./src/bdd"
    // "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main(){

    
    bdd.MakeUser("Louis1", "louis.teilliais@gmail.com", "Karim69lattrik")
    fmt.Println(bdd.GetUser("Louis1"))
    fmt.Println(bdd.GetUserHash("Louis1"))


    // db, _ := sql.Open("sqlite3", "src/bdd/DataBase.db")
    // fmt.Println(db.Query("SELECT * FROM User"))
}