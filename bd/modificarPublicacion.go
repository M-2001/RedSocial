package bd

//"context"
//"time"

//"github.com/M-2001/RedSocial/models"
//"go.mongodb.org/mongo-driver/bson"
//"go.mongodb.org/mongo-driver/bson/primitive"

/*ModificarPublicacion sirve para modificar el registro de una publicacion*/
func ModificarPublicacion( /*pub models.GraboPublicacion, ID string, IDUsuario string) (bool, string, error*/ ) {
	/*conn, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("RedSocial")
	col := db.Collection("publicacion")

	register := make(map[string]interface{})

	if len(pub.UserID) > 0 {
		register["userid"] = pub.UserID
	}
	if len(pub.Publicacion) > 0 {
		register["publicacion"] = pub.Publicacion
	}
	if len(pub.Foto) > 0 {
		register["foto"] = pub.Foto
	}
	if len(pub.Tecnologias) > 0 {
		register["tecnologias"] = pub.Tecnologias
	}
	register["fechaPublicacion"] = pub.FechaPublicacion

	UpdateString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(conn, filter, UpdateString)
	if err != nil {
		return false,IDUsuario, err
	}
	return true,IDUsuario, nil*/
}
