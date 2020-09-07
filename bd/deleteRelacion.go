package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
)

/*DeleteRelacion servira para boorar la relacion de la bd*/
func DeleteRelacion(t models.RelationCollection) (bool, error) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(contt, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
