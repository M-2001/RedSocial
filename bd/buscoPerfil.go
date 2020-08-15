package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil funcion para buscar perfil*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	con, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	ObjID, _ := primitive.ObjectIDFromHex(ID)

	sentecia := bson.M{
		"_id": ObjID,
	}
	err := col.FindOne(con, sentecia).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro No encontrado" + err.Error())
	}
	return perfil, nil
}
