package models

type OrdgoodsModel struct {
	Collection string

}

func NewOrdgoodsModel() *OrdgoodsModel {
	return &OrdgoodsModel{
		Collection:"ordgoods",
	}
}

func (o *OrdgoodsModel) Add(data map[string]interface{}) *DMLResult {

	return nil
}

type Ordgood struct {
	Id         int
	Oid        string
	OrdId      string
	GoodsId    string
	GoodsName  string
	Price      string
	Num        int
	Created_at string
}

func (o *OrdgoodsModel) Query(sql string, params ...interface{}) (interface{}, error) {

	return nil, nil
}
