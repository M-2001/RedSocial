package bd

import (
	"context"
	"log"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReadUnaPublicacion servira para leer una publicacion usuarios*/
func ReadUnaPublicacion(ID string) ([]*models.MostrarUnaPublicacion, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("RedSocial")
	col := db.Collection("publicacion")

	var respuesta []*models.MostrarUnaPublicacion
	/*query := bson.M{
		"_id": ID,
	}*/
	ObjID, err := primitive.ObjectIDFromHex(ID)

	//marcador:= col.FindOne(contt, bson.M{"_id":ObjID})

	marcador, err := col.Find(contt, bson.M{"_id": ObjID})
	if err != nil {
		log.Fatal(err.Error())
		return respuesta, false
	}
	for marcador.Next(context.TODO()) {
		var register models.MostrarUnaPublicacion
		err := marcador.Decode(&register)
		if err != nil {
			return respuesta, false
		}
		respuesta = append(respuesta, &register)
	}
	return respuesta, true
}
