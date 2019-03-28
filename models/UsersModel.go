package models

import "github.com/globalsign/mgo/bson"

type Address struct {
	Addresstype string `bson:"address"`
	Default string  `bson:default`
}

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	UserName string `bson:"userName"`
	Phone string `bson:"phone"`
	Email string `bson:"email"`
	Password string `bson:"password"`
	Salt string `bson:"salt"`
	Address string `bson:"address"`
	V string `bson:"__v"`
}



type UsersModel struct {
	dbName string
	collectionName  string

}

func CreateUsersMode(db, cn string) *UsersModel{
	return &UsersModel{
		dbName:db,
		collectionName:cn,
	}
}

func (u *UsersModel)Insert(docs ...interface{}) error{
	return Insert(u.dbName,u.collectionName,docs)
}
func (u *UsersModel)FindOne(query,selector, result interface{}) error {
	return FindOne(u.dbName, u.collectionName,query,selector, result)
}





