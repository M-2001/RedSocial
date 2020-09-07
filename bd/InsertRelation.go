package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
)

/*InsertRelation servira para insertar una nueva relacion en la base de datos*/
func InsertRelation(t models.RelationCollection) (bool, error) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("relacion")

	_, err := col.InsertOne(contt, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
