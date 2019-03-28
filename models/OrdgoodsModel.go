package models

import "github.com/globalsign/mgo/bson"

type Ordgoods struct {
	Id       bson.ObjectId `bson:"_id"`
	OrdId string `bson:"ordId"`
	GoodsId string `bson:"goodsId"`
	GoodsName string `bson:"intro"`
	Price float64 `bson:"price"`
	Num float64 `bson:"num"`
}

type OrdgoodsModel struct {
	dbName string
	collectionName  string
}
func CreateOrdgoodsMode(db, cn string) *OrdgoodsModel{
	return &OrdgoodsModel{
		dbName:db,
		collectionName:cn,
	}
}

func (u *OrdgoodsModel)Insert(docs ...interface{}) error{
	return Insert(u.dbName,u.collectionName,docs)
}
func (u *OrdgoodsModel)FindOne(query,selector, result interface{}) error {
	return FindOne(u.dbName, u.collectionName,query,selector, result)
}