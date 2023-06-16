package main

import (
	"fmt"
	"portfolio/model/data"
)

func main() {
	fmt.Println("Lancement du script d'installation des éléments de la base de donnée...")

	// Ouverture de la base de donnée
	data.OpenDatabase()

	// Utilisation de la base de donnée d'origine
	fmt.Println("Utilisation de la base de donnée portfolioDB")
	data.DatabaseExecute("USE portfolioDB")

	// Création de la table "PROJETS"
	fmt.Println("Création de la table : PROJETS")
	projetRequete := "CREATE TABLE Projets ( idProjet int primary key auto_increment, nomProjet varchar(255), descriptionProjet TEXT, imageBannerProjet BLOB, imageProfilProjet BLOB,languagesProjet TEXT, competencesProjet TEXT, outilsProjet TEXT, lienProjet TEXT);"
	data.DatabaseExecute(projetRequete)

	// Création de la table "EXPERIENCES"
	fmt.Println("Création de la table : EXPERIENCES")
	experienceRequete := "CREATE TABLE Experiences (idExperience int primary key auto_increment, intituleExperience varchar(255), lieuExperience varchar(255), organisation varchar(255), anneeDebutExperience DATE, anneeFinExperience DATE );"
	data.DatabaseExecute(experienceRequete)

	// Création de la table "FORMATIONS"
	fmt.Println("Création de la table : FORMATIONS")
	formationRequete := "CREATE TABLE Formations (idFormation int primary key auto_increment, nomFormation varchar(255), lieu varchar(255), diplome varchar(255), anneeDebutFormation DATE, anneeFinFormation DATE);"
	data.DatabaseExecute(formationRequete)

	// Création de la table "COMPETENCES"
	fmt.Println("Création de la table : COMPETENCES ")
	competenceRequete := "CREATE TABLE Competences (idCompetence int primary key auto_increment, nomCompetence varchar(255));"
	data.DatabaseExecute(competenceRequete)

}
