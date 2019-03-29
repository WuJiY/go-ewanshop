package models

import (
	"database/sql"
)

type GoodsModel struct {

}

func NewGoodsModel(db *sql.DB) *GoodsModel {
	return nil
}

func (g *GoodsModel) GetByOid(oid string) *Good {

	return nil
}

func (g *GoodsModel) Find(where, opt map[string]interface{}) []*Good {

	return nil
}

func (g *GoodsModel) Add(data map[string]interface{}) *DMLResult {

	return nil
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
	Created_at   string
	Updated_at   string
}

func (g *GoodsModel) Query(sql string, params ...interface{}) (interface{}, error) {

	return nil, nil
}

func (g *GoodsModel) DecrGoodsNum(oid string, num int) *DMLResult {

	return nil
}
