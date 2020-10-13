package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DoCommentPublication servira para que el ussario pueda comentar una publicacion
func DoCommentPublication(t models.CommentPublications) (string, bool, error) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("comentario")

	registro := bson.M{
		"userid":          t.UserID,
		"publicacionid":   t.PublicacionID,
		"comentario":      t.Comentario,
		"fechacomentario": t.FechaComentario,
	}
	result, err := col.InsertOne(contt, registro)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
