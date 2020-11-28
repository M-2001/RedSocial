package models

//"go.mongodb.org/mongo-driver/bson/primitive"

/*ReactComment sirve para dar like a una publicacion*/
type ReactComment struct {
	UserID       string `bson:"userid" json:"userId,omitempty"`
	CommentID    string `bson:"comentarioid" json:"commentid,omitempty"`
	ReactComment string `bson:"reactcomment" json:"reactcomment,omitempty"`
}
