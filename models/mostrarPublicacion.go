package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*MostrarUnaPublicacion estructura que se encargara de mostrar una publicacione*/
type MostrarUnaPublicacion struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID           string             `bson:"userid" json:"userid,omitempty"`
	Publicacion      string             `bson:"publicacion" json:"publicacion,omitempty"`
	Foto             string             `bson:"foto" json:"fotos,omitemptys"`
	Code             string             `bson:"code" json:"code,omitempty"`
	Tecnologias      string             `bson:"tecnologias" json:"tecnologias,omitempty"`
	FechaPublicacion time.Time          `bson:"fecha" json:"fecha"`
}
