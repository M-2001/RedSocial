package bd

import (
	"context"
	"log"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadReaccionesDeUnaPublicacion servira para leer una publicacion usuarios*/
func ReadReaccionesDeUnaPublicacion( /*ID string*/ ) ([]*models.ReaccionCollection, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("RedSocial")
	col := db.Collection("reaccion")

	var respuesta []*models.ReaccionCollection
	/*query := bson.M{
		"_id": ID,
	}*/
	//ObjID, err := primitive.ObjectIDFromHex(ID)

	//marcador:= col.FindOne(contt, bson.M{"_id":ObjID})

	marcador, err := col.Find(contt, bson.D{ /*"_id": ObjID*/ })
	if err != nil {
		log.Fatal(err.Error())
		return respuesta, false
	}
	for marcador.Next(context.TODO()) {
		var register models.ReaccionCollection
		err := marcador.Decode(&register)
		if err != nil {
			return respuesta, false
		}
		respuesta = append(respuesta, &register)
	}
	return respuesta, true
}
