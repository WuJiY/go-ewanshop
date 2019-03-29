package models

import (
	msql "database/sql"
)

type UsersModel struct {

}

func NewUsersModel(db *msql.DB) *UsersModel {
	return nil
}

func (u *UsersModel) Count(where map[string]interface{}) int {

	return 0
}

func (u *UsersModel) Add(data map[string]interface{}) *DMLResult {

	return nil
}

func (u *UsersModel) DelByOid(oid string) *DMLResult {

	return nil
}

func (u *UsersModel) GetByOid(oid string) *User {

	return nil
}

func (u *UsersModel) Find(where, opt map[string]interface{}) []*User {

	return nil
}

type User struct {
	Id         int
	Oid        string
	UserName   string
	Phone      string
	Email      string
	Password   string
	Salt       string
	Address    string
	Created_at string
}

func (u *UsersModel) Query(sql string, params ...interface{}) (interface{}, error) {

	return nil, nil
}

func (u *UsersModel) UpdateByOid(oid string, data map[string]interface{}) *DMLResult {

	return nil
}
