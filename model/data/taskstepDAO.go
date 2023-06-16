package data

import (
	"fmt"
)

func NewTaskstepDAO() {
	OpenDatabase()
}

func ShowAllTable() {
	res, _ := DatabaseQuery("SHOW TABLES")

	// tant qu'il y a des valeurs
	for res.Next() {
		var col1 string
		err := res.Scan(&col1)
		gestionErreur(err)
		fmt.Println(col1)
	}
}
