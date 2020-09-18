package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificarPublicacion sirve para modificar el registro de una publicacion*/
func ModificarPublicacion(pub models.GraboPublicacion, ID string) (bool, error) {
	conn, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("publicacion")

	register := make(map[string]interface{})

	if len(pub.UserId) > 0 {
		register["userid"] = pub.UserId
	}
	if len(pub.Publicacion) > 0 {
		register["publicacion"] = pub.Publicacion
	}

	register["fechaPublicacion"] = pub.FechaPublicacion

	if len(pub.Foto) > 0 {
		register["foto"] = pub.Foto
	}

	UpdateString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(conn, filter, UpdateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
