package mongo_models

import "github.com/globalsign/mgo/bson"

type Cat struct {
	Id       bson.ObjectId `bson:"_id"`
	Cat_name string `bson:"cat_name"`
	Intro string `bson:"intro"`
	Parent_id string `bson:"parent_id"`
}

type CatModel struct {
	dbName string
	collectionName  string
}
func CreateCatMode(db, cn string) *CatModel{
	return &CatModel{
		dbName:db,
		collectionName:cn,
	}
}

func (u *CatModel)Insert(docs ...interface{}) error{
	return Insert(u.dbName,u.collectionName,docs)
}
func (u *CatModel)FindOne(query,selector, result interface{}) error {
	return FindOne(u.dbName, u.collectionName,query,selector, result)
}
