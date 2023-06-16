package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Si pas installer -> Installer package MySQL = go get -u github.com/go-sql-driver/mysql
)

var user string          // utilisateur de la base de donnée
var password string      // password de l'utilisateur
var host string          // host (localhost)
var database_used string // database utilisé
var port string          // port utilisé

var db *sql.DB

/*
Créer une connexion à la base de donnée
*/
func OpenDatabase() {
	// Initier les variables
	fmt.Println("Initialisation des variables")
	initVar()

	// Se connecte à la base de donnée
	fmt.Println("Ouverture de la connexion")
	var err error
	db, err = sql.Open("mysql", user+":"+password+"@("+host+":"+port+")"+"/"+database_used) // "user:password@tcp(host:port)/database"
	gestionErreur(err)

	// defer db.Close()
	fmt.Println("Ouverture de la base de donnée effectué avec succès")
}

/*
Execute une requête SQL sans retourner des données
*/
func DatabaseExecute(requete string) {
	db.Exec(requete)
}

/*
Execute une requête SQL qui renvoie des données
@return *sql.Rows résultat de la requête
*/
func DatabaseQuery(requete string) (*sql.Rows, error) {
	return db.Query(requete)
}

/* Initie les variables de connexion à la base de donnée */
func initVar() {
	user = "root"
	password = ""
	host = "localhost"
	database_used = "taskstep" // portfolioDB
	port = "3306"

}

// Gère l'erreur
func gestionErreur(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
