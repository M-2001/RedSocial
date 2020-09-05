package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*MostrarPublicaciones estructura que se encargara de dar estructura a una lista de publicaciones*/
type MostrarPublicaciones struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID           string             `bson:"userid" json:"userid,omitempty"`
	Publicacion      string             `bson:"publicacion" json:"publicacion,omitempty"`
	FechaPublicacion time.Time          `bson:"fechaPublicacion" json:"fechaPublicacion"`
}
