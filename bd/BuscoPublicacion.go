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

/*BuscoPublicacionSeguidor funcion para buscar perfil*/
// func BuscoPublicacionSeguidor(ID string, page int) (models.PublicacionesSeguidores, bool) {
// 	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	bd := MongoC.Database("RedSocial")
// 	col := bd.Collection("relacion")

// 	skip := (page - 1) * 20

// 	query := make([]bson.M, 0)
// 	query = append(query, bson.M{"$match": bson.M{"userid": ID}})
// 	query = append(query, bson.M{
// 		"$lookup": bson.M{
// 			"from":         "publicacion",
// 			"localField":   "userRelacionid",
// 			"foreignField": "userid",
// 			"as":           "publicacion",
// 		}})
// 	query = append(query, bson.M{"$unwind": "$publicacion"})
// 	query = append(query, bson.M{"$sort": bson.M{"publicacion.fecha": -1}})
// 	query = append(query, bson.M{"$skip": skip})
// 	query = append(query, bson.M{"$limit": 20})

// 	marcador, err := col.Aggregate(contt, query)
// 	var pub models.PublicacionesSeguidores

// 	err = marcador.All(contt, &pub)
// 	if err != nil {
// 		return pub, false
// 	}
// 	return pub, true
// }
