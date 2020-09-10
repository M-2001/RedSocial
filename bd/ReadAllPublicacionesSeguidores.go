package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadPublicacionesSeguidores servira para mostrar las publicaciones de los seguidores relacionados*/
func ReadPublicacionesSeguidores(ID string, page int) ([]models.PublicacionesSeguidores, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoC.Database("RedSocial")
	col := bd.Collection("relacion")

	skip := (page - 1) * 20

	query := make([]bson.M, 0)
	query = append(query, bson.M{"$match": bson.M{"userid": ID}})
	query = append(query, bson.M{
		"$lookup": bson.M{
			"from":         "publicacion",
			"localField":   "userRelacionid",
			"foreignField": "userid",
			"as":           "publicacion",
		}})
	query = append(query, bson.M{"$unwind": "$publicacion"})
	query = append(query, bson.M{"$sort": bson.M{"publicacion.fecha": -1}})
	query = append(query, bson.M{"$skip": skip})
	query = append(query, bson.M{"$limit": 20})

	marcador, err := col.Aggregate(contt, query)
	var result []models.PublicacionesSeguidores

	err = marcador.All(contt, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
