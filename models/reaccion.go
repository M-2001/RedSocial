package models

/*Reaccion structura para obtener el formato que tendra la publicacion*/
type Reaccion struct {
	PublicacionID string `bson:"publicacionid" json:"publicacionid,omitempty"`
}
