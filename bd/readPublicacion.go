package bd

import (
	"context"
	"log"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadPublicaciones servira para leer un nunmero de publicaciones realizadas por los usuarios*/
func ReadPublicaciones(ID string, pagina int64) ([]*models.MostrarPublicaciones, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("RedSocial")
	col := db.Collection("publicacion")

	var respuesta []*models.MostrarPublicaciones
	query := bson.M{
		"userid": ID,
	}
	Options := options.Find()
	Options.SetLimit(20)
	Options.SetSort(bson.D{{Key: "fechaPublicacion", Value: -1}})
	Options.SetSkip((pagina - 1) * 20)

	marcador, err := col.Find(contt, query, Options)
	if err != nil {
		log.Fatal(err.Error())
		return respuesta, false
	}
	for marcador.Next(context.TODO()) {
		var register models.MostrarPublicaciones
		err := marcador.Decode(&register)
		if err != nil {
			return respuesta, false
		}
		respuesta = append(respuesta, &register)
	}
	return respuesta, true
}
