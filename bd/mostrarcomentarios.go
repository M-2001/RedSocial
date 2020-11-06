package bd

import (
	"context"
	"log"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadComentarios servira para leer un nunmero de comentarios por publicacion*/
func ReadComentarios(IDP string) ([]*models.MostrarComentarios, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("RedSocial")
	col := db.Collection("comentario")

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
	//query := bson.M{
	//	"id": ID,
	//}
	//ObjID, err := primitive.ObjectIDFromHex(IDP)

	//marcador, err := col.Find(contt, bson.D{})

	//marcador, err := col.Find(contt, bson.M{"_id": ObjID}),
	marcador, err := col.Aggregate(contt, query)

	//marcador, err := col.Find(contt, query /*bson.M{"publicacionid": ObjID}*/)

	//marcador, err := col.Find(contt, query)

	err = marcador.All(contt, &respuesta)

	if err != nil {
		log.Fatal(err.Error())
		return respuesta, false
	}
	/*for marcador.Next(context.TODO()) {
		var register models.MostrarComentarios
		err := marcador.Decode(&register)
		if err != nil {
			return respuesta, false
		}
		respuesta = append(respuesta, &register)
	}*/
	return respuesta, true
}
