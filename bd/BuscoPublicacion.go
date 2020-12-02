package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPublicacion funcion para buscar perfil*/
func BuscoPublicacion(ID string) (models.GraboPublicacion, error) {
	con, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("publicacion")

	var pub models.GraboPublicacion
	ObjID, _ := primitive.ObjectIDFromHex(ID)

	sentecia := bson.M{
		"_id": ObjID,
	}
	err := col.FindOne(con, sentecia).Decode(&pub)
	if err != nil {
		fmt.Println("Publicacion No encontrada" + err.Error())
		return pub, err
	}
	return pub, nil
}
