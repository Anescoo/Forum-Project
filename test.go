package main 

import (
	"fmt"
	bdd "./src/bdd"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main(){

	
	bdd.MakeUser("Louis", "louis.teilliais@gmail.com", "ATYyuiop1234")
	fmt.Println(bdd.GetUser("Louis"))

}