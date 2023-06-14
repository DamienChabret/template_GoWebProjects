package main

import (
	"fmt"
	"portfolio/model/data"
)

func main() {
	fmt.Println("Lancement du script d'installation des éléments de la base de donnée...")

	data.OpenDatabase()
}
