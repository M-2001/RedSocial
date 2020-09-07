package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultaRelacion consulta la relacion entre usuarios*/
func ConsultaRelacion(t models.RelationCollection) (bool, error) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoC.Database("RedSocial")
	col := bd.Collection("relacion")

	query := bson.M{
		"userid":         t.UserID,
		"userRelacionid": t.UserRelacionID,
	}
	var result models.RelationCollection
	fmt.Println(result)
	err := col.FindOne(contt, query).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
