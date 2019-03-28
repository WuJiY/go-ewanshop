package models

import "github.com/globalsign/mgo/bson"

type Ordinfos struct {
	Id       bson.ObjectId `bson:"_id"`
	OrdId string `bson:"ordId"`
	UserId string  `bson:"userId"`
	UserName string `bson:"userName"`
	Address string `bson:"address"`
	PayType string `bson:"payType"`
	PayState bool `bson:"payState"`
	Money float64 `bson:"money"`
	Fuyan string  `bson:"fuyan"`
}

type OrdinfosModel struct {
	dbName string
	collectionName  string
}
func CreateOrdinfosMode(db, cn string) *OrdinfosModel{
	return &OrdinfosModel{
		dbName:db,
		collectionName:cn,
	}
}

func (u *OrdinfosModel)Insert(docs ...interface{}) error{
	return Insert(u.dbName,u.collectionName,docs)
}
func (u *OrdinfosModel)FindOne(query,selector, result interface{}) error {
	return FindOne(u.dbName, u.collectionName,query,selector, result)
}