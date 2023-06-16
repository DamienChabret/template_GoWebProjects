package metier

import "time"

/*
Structure d'une "FORMATION"
*/
type Formation struct {
	Id         int
	Nom        string
	Lieu       string
	Diplome    string
	AnneeDebut time.Time
	AnneeFin   time.Time
}
