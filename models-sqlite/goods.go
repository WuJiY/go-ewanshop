package models

import (
	"database/sql"
)

type GoodsModel struct {
	db *sql.DB
}

func NewGoodsModel(db *sql.DB) *GoodsModel {
	return &GoodsModel{db}
}

func (g *GoodsModel) GetByOid(oid string) *Good {
	var gs = g.Find(map[string]interface{}{"oid": oid}, nil)
	if 0 < len(gs) {
		return gs[0]
	}
	return nil
}

func (g *GoodsModel) Find(where, opt map[string]interface{}) []*Good {
	var sql = "select * from goods"
	r, err := Find(sql, where, opt, g.Query)
	if err != nil {
		panic(err)
	}
	return r.([]*Good)
}

func (g *GoodsModel) Add(data map[string]interface{}) *DMLResult {
	var ret, err = Add(g.db, "goods", data)
	if err != nil {
		panic(err)
	}
	return ret
}

func (g *GoodsModel) DelByOid(oid string) *DMLResult {
	var r, err = DML(g.db, "delete from goods where oid=?", oid)
	if err != nil {
		panic(err)
	}
	return r
}

func (g *GoodsModel) Count(where map[string]interface{}) int {
	var ret, err = Count(g.db, "goods", where)
	if err != nil {
		panic(err)
	}
	return ret
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
	rows, err := g.db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]*Good, 0)
	for rows.Next() {
		var good = &Good{}
		var err = rows.Scan(&good.Id, &good.Oid, &good.Goods_name, &good.Cat_id, &good.Shop_price, &good.Goods_img, &good.Goods_desc,
			&good.Goods_number, &good.Is_best, &good.Is_new, &good.Is_hot, &good.Is_on_sale, &good.Created_at, &good.Updated_at)
		if err != nil {
			return nil, err
		}
		ret = append(ret, good)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (g *GoodsModel) DecrGoodsNum(oid string, num int) *DMLResult {
	var ret, err = DML(g.db, "update goods set goods_number=goods_number-? where oid=? and goods_number>=?", num, oid, num)
	if err != nil {
		panic(err)
	}
	return ret
}
