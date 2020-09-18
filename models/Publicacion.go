package models

/*Publicacion structura para obtener el formato que tendra la publicacion*/
type Publicacion struct {
	Publicacion string `bson:"publicacion" json:"publicacion"`
	//Foto        string `bson:"foto" json:"foto"`
	Tecnologias string `bson:"tecnologias" json:"tecnologias"`
}
