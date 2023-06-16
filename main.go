package main

import (
	"fmt"
	"net/http"
	handlers "portfolio/controller"
)

func main() {
	fmt.Println("Démarage du serveur")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/me", handlers.Me)
	http.HandleFunc("/experiences", handlers.Experience)
	http.HandleFunc("/formation", handlers.Formation)
	http.HandleFunc("/project", handlers.Project)
	http.HandleFunc("/contact", handlers.Contact)

	http.ListenAndServe(":5500", nil)
}
