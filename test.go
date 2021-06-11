package main

import (
	"fmt"

	bdd "./src/bdd"
)

func main() {
	fmt.Println(bdd.GetAllPoste())
	fmt.Println(bdd.GetPosteByID(1))
}
