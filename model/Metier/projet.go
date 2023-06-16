package metier

/*
Structure d'un "PROJET"
*/
type Projet struct {
	Id           int
	Nom          string
	Description  string
	ImageBanner  []byte
	ImageProfile []byte
	Languages    []string
	Competences  []string
	Outils       []string
	Lien         string
}
