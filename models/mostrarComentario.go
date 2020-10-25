package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*MostrarComentarios estructura que se encargara de mostrar una publicacione*/
type MostrarComentarios struct {
	ID              primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID          string             `bson:"userid" json:"userid,omitempty"`
	PublicacionID   string             `bson:"publicacionid" json:"publicacionid,omitempty"`
	Comentario      string             `bson:"comentario" json:"comentario,omitempty"`
	FechaComentario time.Time          `bson:"fechacomentario" json:"fechacomentario,omitempty"`
}
