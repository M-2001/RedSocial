package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadAllUsers se enacrga de leer todos los usuarios en la red relacionados con un solo usuario*/
func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	contt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoC.Database("RedSocial")
	col := bd.Collection("usuarios")

	var results []*models.Usuario

	FOptions := options.Find()
	FOptions.SetSkip((page - 1) * 20)
	FOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	marcador, err := col.Find(contt, query, FOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var find, include bool

	for marcador.Next(contt) {
		var users models.Usuario
		err := marcador.Decode(&users)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var relacion models.RelationCollection
		relacion.UserID = ID
		relacion.UserRelacionID = users.ID.Hex()

		include = false
		find, err = ConsultaRelacion(relacion)
		if tipo == "new" && find == false {
			include = true
		}
		if tipo == "follow" && find == true {
			include = true
		}
		if relacion.UserRelacionID == ID {
			include = false
		}
		if include == true {
			users.Password = ""
			users.Biografia = ""
			users.SitioWeb = ""
			users.Ubicacion = ""
			users.Banner = ""
			users.Email = ""

			results = append(results, &users)
		}

	}
	err = marcador.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false

	}
	marcador.Close(contt)
	return results, true
}
