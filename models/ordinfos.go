package models

import (
	"database/sql"
)

type OrdinfosModel struct {

}

func NewOrdinfosModel(db *sql.DB) *OrdinfosModel {
	return nil
}

func (o *OrdinfosModel) Add(data map[string]interface{}) *DMLResult {

	return nil
}

func (o *OrdinfosModel) GetByOid(oid string) *Ordinfo {

	return nil
}

func (o *OrdinfosModel) Find(where, opt map[string]interface{}) []*Ordinfo {

	return nil
}

type Ordinfo struct {
	Id         int
	Oid        string
	OrdId      string
	UserId     string
	UserName   string
	Address    string
	PayType    string
	PayState   int
	Money      string
	Fuyan      string
	Created_at string
}

func (o *OrdinfosModel) Query(sql string, params ...interface{}) (interface{}, error) {

	return nil, nil
}
