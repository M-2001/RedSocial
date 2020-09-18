package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertPublicacion servira para que el usuario inserte un tweet*/
func InsertPublicacion(t models.GraboPublicacion) (string, bool, error) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("publicacion")

	registro := bson.M{
		"userid":      t.UserId,
		"publicacion": t.Publicacion,
		"foto":        t.Foto,
		"tecnologias": t.Tecnologias,
		"fecha":       t.FechaPublicacion,
	}
	result, err := col.InsertOne(contt, registro)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
