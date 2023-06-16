package model

import "time"

/*
Structure d'une "EXPERIENE"
*/
type Experience struct {
	Id           int
	Intitule     string
	Lieu         string
	Organisation string
	AnneeDebut   time.Time
	AnneeFin     time.Time
}
