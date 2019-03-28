package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/wanchain/go-wanchain/console"
	"testing"
)
func TestFindOne(t *testing.T){
	openDB()
	userModel := CreateUsersMode("test","users");

    result := &User{}

   userModel.FindOne(bson.M{},bson.M{},result);
   console.log(result.UserName);


}
