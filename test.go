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
	

}