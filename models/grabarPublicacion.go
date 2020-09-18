package models

import "time"

/*GraboPublicacion estrutura para insertar un publicacion en bson*/
type GraboPublicacion struct {
	UserId           string    `bson:"userid" json:"userid,omitempty"`
	Publicacion      string    `bson:"publicacion" json:"publicacion,omitempty"`
	Foto             string    `bson:"foto" json:"foto"`
	Tecnologias      string    `bson:"tecnologias" json:"tecnologias,,omitempty"`
	FechaPublicacion time.Time `bson:"fechaPublicacion" json:"fechaPublicacion,omitempty"`
}
