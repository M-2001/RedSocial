package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*GraboPublicacion estrutura para insertar un publicacion en bson*/
type GraboPublicacion struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID      string             `bson:"userid" json:"userid,omitempty"`
	Publicacion string             `bson:"publicacion" json:"publicacion,omitempty"`
	//Foto             string             `bson:"foto" json:"foto"`
	Code             string    `bson:"code" json:"code,omitempty"`
	Tecnologias      string    `bson:"tecnologias" json:"tecnologias,omitempty"`
	FechaPublicacion time.Time `bson:"fechaPublicacion" json:"fechaPublicacion,omitempty"`
}
