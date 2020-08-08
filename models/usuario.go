package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario*/
type Usuario struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre    string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido  string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaN    time.Time          `bson:"fechaN" json:"fechaN,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biografia string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb  string             `bson:"sitioWeb" son:"sitioWeb,omitempty"`
}
