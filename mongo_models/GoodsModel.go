package mongo_models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Goods struct {
	Id       bson.ObjectId `bson:"_id"`
	Goods_name string `bson:"goods_name"`
	Cat_id  string `bson:"cat_id"`
	Shop_price float64 `bson:"shop_price"`
	Goods_img string `bson:"goods_img"`
	Goods_desc string `bson:"goods_desc"`
	Goods_number float64 `bson:"goods_number"`
	Is_best bool `bson:"is_best"`
	Is_new bool `bson:"is_new"`
	Is_hot bool `bson:"is_hot"`
	Is_on_sale bool `bson:"is_on_sale"`
	CreatedAt time.Time `bson:"createdAt"`
}

type GoodsModel struct {
	dbName string
	collectionName  string
}
func CreateGoodstMode(db, cn string) *GoodsModel{
	return &GoodsModel{
		dbName:db,
		collectionName:cn,
	}
}

func (u *GoodsModel)Insert(docs ...interface{}) error{
	return Insert(u.dbName,u.collectionName,docs)
}
func (u *GoodsModel)FindOne(query,selector, result interface{}) error {
	return FindOne(u.dbName, u.collectionName,query,selector, result)
}
