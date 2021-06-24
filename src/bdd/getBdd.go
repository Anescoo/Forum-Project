package bdd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "src/bdd/DataBase.db") //ouverture de la bdd
	if err != nil {                                       //cas d'erreur
		fmt.Println("Error database")
		return 500, nil //renvoie du code d'erreur
	}
	return 0, db //renvoie la base de donner
}

/////////////////////////////////User///////////////////////////////////////////////////////////////////////////////

func MakeUser(pseudo string, email string, hash string) int {
	_, db := OpenDataBase()                                                                //ouverture de la bdd
	result, err := db.Prepare("INSERT INTO User (Pseudo, Email, HashMDP) VALUES(?, ?, ?)") //Preparation de la commande
	if err != nil {
		fmt.Println("Error MakeUser") //cas d'erreur
		return 500
	} else {
		result.Exec(pseudo, email, hash) //creation du l'utilisateur
	}
	db.Close()
	return 0
}

func GetUser(username string) (int, [2]string) {

	_, db := OpenDataBase()

	var resultFunc [2]string
	result, err := db.Query("SELECT Pseudo, Email FROM User WHERE Pseudo = ?", username) // recherche de l'utilisateur avec le pseudo voulue
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	}

	var Pseudo string
	var Email string

	for result.Next() { //recuperation des info sur la bdd
		result.Scan(&Pseudo, &Email)
	}

	resultFunc[0] = Pseudo
	resultFunc[1] = Email

	db.Close()
	return 0, resultFunc //renvoie l'utilisateur
}

func GetAllUser() (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT Pseudo, Email FROM User")
	if err != nil {
		fmt.Println("Error GetAllUser")
		fmt.Println(err.Error())

		return 500, resultFunc
	}

	var Pseudo string
	var Email string

	for result.Next() {
		result.Scan(&Pseudo, &Email)
		temp := []string{Pseudo, Email}
		resultFunc = append(resultFunc, temp)
	}

	db.Close()
	return 0, resultFunc
}

func GetUserHash(username string) (int, string) {

	_, db := OpenDataBase()
	result, err := db.Query("SELECT HashMDP FROM User WHERE Pseudo = ?", username) //renvoie du hash de l'utilisateur voulue
	if err != nil {
		fmt.Println("Error GetUserHash")
		db.Close()
		return 500, "error 500"
	} else {

		var HashMDP string

		result.Next() //lecture de la bdd
		result.Scan(&HashMDP)

		db.Close()
		return 0, HashMDP //renvoie le hash
	}
}

////////////////////////////////////Poste///////////////////////////////////////////////////////////////////////////////////////

func MakePoste(user string, contenue string, categorie string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Poste (PseudoUser, Contenue, Categorie, nbLike) VALUES(?,?,?,?)")
	if err != nil {
		fmt.Println("Error MakePoste")
		fmt.Println(err.Error())
		return 500
	} else {
		_, err := result.Exec(user, contenue, categorie, 0)
		if err != nil {
			fmt.Println(err.Error())
			db.Close()
			return 500
		} else {
			db.Close()
			return 0
		}
	}
}

func GetPosteByID(id int) (int, []string) {
	_, db := OpenDataBase()
	var resultFunc []string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike, DatePoste FROM Poste WHERE id = ?", id)
	if err != nil {
		fmt.Println("Error GetPosteById")
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {
		fmt.Println("Starting GetPosteById")

		var ID int
		var PseudoUser string
		var Contenue string
		var Categorie string
		var DatePoste string
		var nbLike int

		result.Next()
		result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike, &DatePoste)

		resultFunc = append(resultFunc, strconv.Itoa(ID))
		resultFunc = append(resultFunc, PseudoUser)
		resultFunc = append(resultFunc, Contenue)
		resultFunc = append(resultFunc, Categorie)
		resultFunc = append(resultFunc, strconv.Itoa(nbLike))
		resultFunc = append(resultFunc, DatePoste)

		db.Close()
		return 0, resultFunc
	}
}

func GetAllPoste() (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike, DatePoste FROM Poste ORDER BY DatePoste DESC")
	if err != nil {
		fmt.Println("Error GetAllPoste")
		fmt.Println(err.Error())
		return 500, resultFunc
	}

	var ID int
	var PseudoUser string
	var Contenue string
	var Categorie string
	var DatePoste string
	var nbLike int

	for result.Next() {
		result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike, &DatePoste)
		temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, Categorie, strconv.Itoa(nbLike), DatePoste}
		resultFunc = append(resultFunc, temp)
	}
	db.Close()
	return 0, resultFunc
}

func GetPosteByCategorie(categorie string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, nbLike, DatePoste FROM Poste WHERE Categorie = ?", categorie)
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {

		var ID int
		var PseudoUser string
		var Contenue string
		var nbLike int
		var DatePoste string

		for result.Next() {
			result.Scan(&ID, &PseudoUser, &Contenue, &nbLike, &DatePoste)
			temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, strconv.Itoa(nbLike), DatePoste}
			resultFunc = append(resultFunc, temp)
		}

		db.Close()
		return 0, resultFunc
	}
}

func GetPosteByUser(UserPseudo string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike, DatePoste FROM Poste WHERE PseudoUser = ? ORDER BY DatePoste DESC", UserPseudo)
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {

		var ID int
		var PseudoUser string
		var Contenue string
		var Categorie string
		var nbLike int
		var DatePoste string

		for result.Next() {
			result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike, &DatePoste)
			temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, Categorie, strconv.Itoa(nbLike), DatePoste}
			resultFunc = append(resultFunc, temp)
		}

		db.Close()
		return 0, resultFunc
	}
}

func UpdatePoste(id int, update string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("UPDATE Poste SET Contenue = ? WHERE ID = ?")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(update, id)
		db.Close()
		return 0
	}
}

func DeletePoste(id int) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("DELETE FROM Poste WHERE ID = ?")
	if err != nil {
		fmt.Println(err.Error())
		return 500
	} else {
		_, err := result.Exec(id)
		if err != nil {
			fmt.Println(err.Error())
			return 500
		} else {
			db.Close()
			return 0
		}
	}
}

///////////////////////Like///////////////////////////////////////////////////////////////////////

func Like(idPost int, PseudoUser string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Like (PseudoUser, idPoste) VALUES(?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return 500
	} else {
		result.Exec(PseudoUser, idPost)
		db.Close()
		return 0
	}
}

func Unlike(idPost int, PseudoUser string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("DELET FROM Like WHERE idPoste = ? AND PseudoUser =?")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(idPost, PseudoUser)
		db.Close()
		return 0
	}
}

func GetPosteLikeByUser(UserPseuso string) (int, [][]string) {
	_, db := OpenDataBase()
	var ResultFunc [][]string
	var IdLiker []int
	temp, err := db.Query("SELECT idPoste FROM Like WHERE PseudoUser = ?", UserPseuso)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, ResultFunc
	} else {
		var IDPoste int

		for temp.Next() {
			temp.Scan(&IDPoste)
			IdLiker = append(IdLiker, IDPoste)
		}

		db.Close()
		for i := 0; i < len(IdLiker); i++ {
			_, db := OpenDataBase()
			result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike, DatePoste FROM Poste WHERE ID = ? ORDER BY DatePoste DESC ", IdLiker[i])
			if err != nil {
				fmt.Println(err.Error())
				db.Close()
				return 500, ResultFunc
			} else {

				var ID int
				var PseudoUser string
				var Contenue string
				var Categorie string
				var nbLike int
				var DatePoste string

				for result.Next() {
					result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike, &DatePoste)
					temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, Categorie, strconv.Itoa(nbLike), DatePoste}
					ResultFunc = append(ResultFunc, temp)
				}
			}
		}
		return 0, ResultFunc
	}
}

func IsLike(idPoste int, PseudoUser string) (int, bool) {
	_, db := OpenDataBase()
	_, err := db.Query("SELECT PseudoUser FROM Like WHERE PseudoUser = ? AND IdPoste = ?", PseudoUser, idPoste)
	if err != nil {
		db.Close()
		return 500, false
	} else {
		db.Close()
		return 0, true
	}
}

func GetLikeNb(id int) int {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT idPoste FROM Like WHERE idPoste = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		var resultFunc []int
		var idPoste int

		for result.Next() {
			result.Scan(&idPoste)
			resultFunc = append(resultFunc, idPoste)
		}

		db.Close()
		// fmt.Println(len(resultFunc))
		return len(resultFunc)
	}
}

////////////////////////Categorie//////////////////////////////////////////////////////////////////////////////

func MakeCategorie(name string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Categorie (Name) Values(?)")
	if err != nil {
		db.Close()
		return 500
	} else {
		result.Exec(name)
	}

	db.Close()
	return 0
}

func GetAllCategorie() (int, []string) {
	_, db := OpenDataBase()
	var ResultFunc []string
	result, err := db.Query("SELECT Name FROM Categorie")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, ResultFunc
	} else {
		var NameCategorie string

		for result.Next() {
			result.Scan(&NameCategorie)
			ResultFunc = append(ResultFunc, NameCategorie)
		}

		db.Close()
		return 0, ResultFunc
	}
}

func NbPosteByCategorie(name string) (int, int) {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT ID FROM Poste WHERE Categorie = ?", name)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 0, 500
	} else {
		var temp []int
		var ID int

		for result.Next() {
			result.Scan(&ID)
			temp = append(temp, ID)
		}
		return 0, len(temp)
	}
}

func DeleteCategorire(name string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("DELETE FROM Categorie WHERE Name = ? ")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(name)
		db.Close()
		return 0
	}
}

//////////////////////////////////////////Commentaire//////////////////////////////////////////////////////

func MakeComment(Username string, Contenue string, IDPoste int) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Commentaire (User, Contenue, IDPoste) VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(Username, Contenue, IDPoste)
		db.Close()
		return 0
	}
}

func GetCommentByPoste(idPoste int) (int, [][]string) {
	_, db := OpenDataBase()
	var ResultFunc [][]string
	result, err := db.Query("SELECT User, Contenue FROM Commentaire WHERE IDPoste = ?", idPoste)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, ResultFunc
	} else {
		var User string
		var Contenue string

		for result.Next() {
			result.Scan(&User, &Contenue)
			temp := []string{User, Contenue}
			ResultFunc = append(ResultFunc, temp)
		}
		db.Close()
		return 0, ResultFunc
	}
}
