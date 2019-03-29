package models

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"log"
	"time"
)

type GoodsModel struct {
	Collection string

}

func NewGoodsModel() *GoodsModel {
	return &GoodsModel{
		Collection:"goods",
	}
}

func (g *GoodsModel) GetByOid(oid string) *Good {

	return nil
}

func (g *GoodsModel) Find(where, opt map[string]interface{}) []*Good {
	var goods []*Good
	err := FindAll(DbName,g.Collection,bson.M{},bson.M{},&goods)
	if(err != nil){
		fmt.Println("err is", err)
	}
	return goods

}

func (g *GoodsModel) Add(data map[string]interface{}){
	err := Insert(DbName,g.Collection,&data)
	if(err != nil ){
		fmt.Println("err is ", err)
	}else{
		//log.Info("add one cate in db")
		log.Println("add one goods in db")
	}

}

func (g *GoodsModel) DelByOid(oid string) *DMLResult {

	return nil
}

func (g *GoodsModel) Count(where map[string]interface{}) int {

	return 0
}

type Good struct {
	Id           int
	Oid          string
	Goods_name   string
	Cat_id       string
	Shop_price   string
	Goods_img    string
	Goods_desc   string
	Goods_number int
	Is_best      int
	Is_new       int
	Is_hot       int
	Is_on_sale   int
	Created_at   time.Time  `bson:"createAt"`
	Updated_at   time.Time
}

func (g *GoodsModel) Query(sql string, params ...interface{}) (interface{}, error) {

	return nil, nil
}

func (g *GoodsModel) DecrGoodsNum(oid string, num int) *DMLResult {

	return nil
}
