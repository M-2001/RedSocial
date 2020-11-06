package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*MostrarReacciones sirve para mostrar reaciones de una publicacion*/
type MostrarReacciones struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID        string             `bson:"userid" json:"userid,omitempty"`
	PublicacionID string             `bson:"publicacionid" json:"publicacionid,omitempty"`
	FechaReaccion time.Time          `bson:"fechareaccion" json:"fechareaccion,omitempty"`
}
