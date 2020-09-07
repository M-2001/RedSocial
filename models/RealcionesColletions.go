package models

/*RelationCollection sirve para relacionar los demas usuarios*/
type RelationCollection struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelacionID string `bson:"userRelacionid" json:"userRelacionId"`
}
