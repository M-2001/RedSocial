package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReaccionPublicacion servira para que el usuario pueda reaccionar a una publicacion
func ReaccionPublicacion(t models.ReaccionCollection) (string, bool, error) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("reaccion")

	registro := bson.M{
		"userid":        t.UserID,
		"publicacionid": t.PublicacionID,
		"fechareaccion": time.Now(),
	}
	result, err := col.InsertOne(contt, registro)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

/*DELETE REACTION*/

/*DeleteReaccion servira para boorar la relacion de la bd*/
func DeleteReaccion(ID string, UserID string) error {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("reaccion")

	objID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(contt, query)
	return err
}
