package models

import "github.com/globalsign/mgo/bson"

type Cate struct {
	Id       bson.ObjectId `bson:"_id"`
	Cat_name string `bson:"cat_name"`
	Intro string `bson:"intro"`
	Parent_id string `bson:"parent_id"`
}

type CatesModel struct {
	dbName string
	collectionName  string
}
func CreateCatMode(db, cn string) *CatesModel{
	return &CatesModel{
		dbName:db,
		collectionName:cn,
	}
}

func (u *CatesModel)Insert(docs ...interface{}) error{
	return Insert(u.dbName,u.collectionName,docs)
}
func (u *CatesModel)FindOne(query,selector, result interface{}) error {
	return FindOne(u.dbName, u.collectionName,query,selector, result)
}


func (u *CatesModel)Find() *[]Cate{
	return nil
}
func (u *CatesModel)GetChildCates(allCate *[]Cate, catid string) []string {
	return nil
}