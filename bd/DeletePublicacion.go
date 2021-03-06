package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeletePublicacion servira para borrar un publicacion en especifico*/
func DeletePublicacion(ID string, UserID string) error {
	contt, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("publicacion")

	objID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(contt, query)
	return err
}
