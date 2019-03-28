package models
const (
	DbName = "test"
	UserCollection= "user"
	CateCollection ="cate"
	GoodsCollection = "goods"
	OrdgoodsCollection = "ordgoods"
	OrdInfoCollection = "ordinfo"
)

func CreateDB(){
	openDB();
}

