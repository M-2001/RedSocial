package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*ReaccionCollection sirve para dar like a una publicacion*/
type ReaccionCollection struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserID        string             `bson:"userid" json:"userId"`
	PublicacionID string             `bson:"publicacionid" json:"publicacionid"`
}
