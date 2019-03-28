package mongo_models



const (
	DbName = "test"
	UserCollection= "users"
	CateCollection ="cate"
	GoodsCollection = "goods"
	OrdgoodsCollection = "ordgoods"
	OrdInfoCollection = "ordinfo"
)

func CreateDB(){
	openDB();
}

