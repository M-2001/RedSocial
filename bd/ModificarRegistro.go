package bd

import (
	"context"
	"time"

	"github.com/M-2001/RedSocial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificarRegistro sirve para modificar el registro de un usuario*/
func ModificarRegistro(user models.Usuario, ID string) (bool, error) {
	conn, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("usuarios")

	register := make(map[string]interface{})
	if len(user.Nombre) > 0 {
		register["nombre"] = user.Nombre
	}
	if len(user.Apellido) > 0 {
		register["apellidos"] = user.Apellido
	}

	register["fechaN"] = user.FechaN

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	if len(user.Biografia) > 0 {
		register["biografia"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		register["ubicacion"] = user.Ubicacion
	}
	if len(user.SitioWeb) > 0 {
		register["sitioWeb"] = user.SitioWeb
	}
	UpdateString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(conn, filter, UpdateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
