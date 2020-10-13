package models

/*Comentario structura para obtener el formato que tendra la publicacion*/
type Comentario struct {
	PublicacionID string `bson:"publicacionid" form-data:"publicacion,omitempty"`
	Comentario    string `bson:"comentario" form-data:"comentario,omitempty"`
}
