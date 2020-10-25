/*package bd

import (
	"context"
	"log"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadReacciones servira para leer un nunmero de reacciones por publicacion
func ReadReacciones(IDP string) ([]*models.ReaccionCollection, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("RedSocial")
	col := db.Collection("reaccion")

	var respuesta []*models.MostrarComentarios

	query := make([]bson.M, 0)
	query = append(query, bson.M{"$match": bson.M{"publicacionid": IDP}})
	query = append(query, bson.M{
		"$lookup": bson.M{
			"from":         "publicacion",
			"localField":   "publicacionid",
			"foreignField": "_id",
			"as":           "comentariobyPublicacion",
		}})

	query = append(query, bson.M{"$unwind": "$comentario"})
	query = append(query, bson.M{"$sort": bson.M{"fechacomentario": -1}})

	marcador, err := col.Aggregate(contt, query)


	err = marcador.All(contt, &respuesta)

	if err != nil {
		log.Fatal(err.Error())
		return respuesta, false
	}
	return respuesta, true
}*/
