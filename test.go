<<<<<<< HEAD
package main 

import (
	"fmt"
	bdd "./src/bdd"
	// "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main(){

	
	fmt.Println(bdd.GetAllUser())


	fmt.Println(bdd.GetPosteByID(2))
	

=======
package main 

import (
    "fmt"
    bdd "./src/bdd"
    // "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main(){

    // _, db := bdd.OpenDataBase()

    // db.Exec("DELETE FROM Poste WHERE ID = 23")
    // bdd.MakeUser("Tao", "louis.teilliais@gmail.com", "Karim69lattrik")
	// bdd.MakePoste("Tao", "Ce post s'affiche dans la console", "test")
	// fmt.Println(bdd.GetPosteByID(2)[1])
    fmt.Println(bdd.GetAllPoste())
    fmt.Println(bdd.GetPosteByID(23))
    _, NbrPosts := bdd.GetAllPoste()
    fmt.Println(len(NbrPosts))
    // db, _ := sql.Open("sqlite3", "src/bdd/DataBase.db")
    // fmt.Println(db.Query("SELECT * FROM User"))
>>>>>>> 4a706add021e8ce8384cee8ac19cad24217ce935
}