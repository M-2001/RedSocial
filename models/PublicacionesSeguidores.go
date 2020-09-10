package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*PublicacionesSeguidores estructura para devolver las publicaciones */
type PublicacionesSeguidores struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserRelacionID string             `bson:"userRelacionid" json:"userRelationid,omitempty"`
	UserID         string             `bson:"userid" json:"userid,omitempty"`
	Publicacion    struct {
		Publicacionn string    `bson:"publicacion" json:"publicacion,omitempty"`
		Fecha        time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID           string    `bson:"_id" json:"_id,omitempty"`
	}
}
