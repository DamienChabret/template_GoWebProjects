
CREATE TABLE Projet ( 
   idProjet int primary key auto_increment, 
   nomProjet varchar(255), 
   descriptionProjet TEXT, 
   imageBannerProjet BLOB, 
   imageProfilProjet BLOB, 
   languagesProjet TEXT, 
   competencesProjet TEXT, 
   outilsProjet TEXT, 
   lienProjet TEXT
   );

CREATE TABLE Experiences (
   idExperience int primary key auto_increment, 
   intituleExperience varchar(255), 
   lieuExperience varchar(255), 
   organisation varchar(255), 
   anneeDebutExperience DATE, 
   anneeFinExperience DATE);

CREATE TABLE Formations (
   idFormation int primary key auto_increment, 
   nomFormation varchar(255), lieu varchar(255), 
   diplomeFormation varchar(255), 
   anneeDebutFormation DATE, 
   anneeFinFormation DATE);

CREATE TABLE Competences (
   idCompetence int primary key auto_increment, 
   nomCompetence varchar(255)
   );