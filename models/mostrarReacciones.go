package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*MostrarReacciones sirve para dar like a una publicacion*/
type MostrarReacciones struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	PublicacionID  string             `bson:"publicacionid" json:"publicacionid,omitempty"`
	FecharReaccion time.Time          `bson:"fechareaccion" json:"fechareaccion,omitempty"`
}
