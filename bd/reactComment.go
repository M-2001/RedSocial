package bd

import (
	"context"
	"log"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReactComment servira para que el usuario pueda reaccionar a una publicacion
func ReactComment(t models.ReactComment) (string, bool, error) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	DB := "RedSocial"
	db := MongoC.Database(DB)
	col := db.Collection("reactComment")

	registro := bson.M{
		"userid":       t.UserID,
		"commentid":    t.CommentID,
		"reactComment": t.ReactComment,
	}
	result, err := col.InsertOne(contt, registro)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

//ReadReactComment servira para leer un nunmero de reacciones por comentaerios
func ReadReactComment(IDp string) ([]*models.ReactComment, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("RedSocial")
	col := db.Collection("reactComment")

	var respuesta []*models.ReactComment

	query := make([]bson.M, 0)
	query = append(query, bson.M{"$match": bson.M{"commentid": IDp}})
	query = append(query, bson.M{
		"$lookup": bson.M{
			"from":         "comentario",
			"localField":   "commentid",
			"foreignField": "_id",
			"as":           "reaccionbyComment",
		}})

	query = append(query, bson.M{"$unwind": "$commentid"})
	//query = append(query, bson.M{"$sort": bson.M{"fechareaccion": -1}})

	marcador, err := col.Aggregate(contt, query)

	err = marcador.All(contt, &respuesta)

	if err != nil {
		log.Fatal(err.Error())
		return respuesta, false
	}

	return respuesta, true
}

/*DeleteReactComment servira para borrar un publicacion en especifico*/
func DeleteReactComment(ID string, UserID string) error {
	contt, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("reactComment")

	objID, _ := primitive.ObjectIDFromHex(ID)

	query := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(contt, query)
	return err
}
