package models

import (
	"database/sql"
)



type CatesModel struct {

}

func NewCatesModel(db *sql.DB) *CatesModel {
	return nil
}

type CateTreeItem struct {
	Oid       string
	Cat_name  string
	Intro     string
	Parent_id string
	Level     int
}

func (c *CatesModel) GetTree(rows []*Cate) []*CateTreeItem {
	return c.getTree(rows, "0", 0)
}

func (c *CatesModel) getTree(rows []*Cate, pid string, level int) []*CateTreeItem {
	var tree = make([]*CateTreeItem, 0)
	for _, row := range rows {
		if row.Parent_id == pid {
			tree = append(tree, &CateTreeItem{
				Oid:       row.Oid,
				Cat_name:  row.Cat_name,
				Intro:     row.Intro,
				Parent_id: row.Parent_id,
				Level:     level,
			})
			tree = append(tree, c.getTree(rows, row.Oid, level+1)...)
		}
	}
	return tree
}

func (c *CatesModel) GetFamily(rows []*Cate, catid string) []*Cate {
	var arr = make([]*Cate, 0)
	var isFind bool
	for catid != "0" {
		isFind = false
		for _, row := range rows {
			if row.Oid == catid {
				var tmp = make([]*Cate, len(arr)+1)
				tmp[0] = row
				copy(tmp[1:], arr)
				arr = tmp
				catid = row.Parent_id
				isFind = true //避免死循环
				break
			}
		}
		if !isFind {
			break
		}
	}
	return arr
}

func (c *CatesModel) GetChildCates(rows []*Cate, catid string) []string {
	var arr = make([]string, 0)
	for _, r := range rows {
		if r.Parent_id == catid {
			arr = append(arr, r.Oid)
			arr = append(arr, c.GetChildCates(rows, r.Oid)...)
		}
	}
	return arr
}

func (c *CatesModel) Find() []*Cate {

	return nil
}

type Cate struct {
	Id         int
	Oid        string
	Cat_name   string
	Intro      string
	Parent_id  string
	Created_at string
}

func (c *CatesModel) Query(sql string, params ...interface{}) (interface{}, error) {

	return nil, nil
}

func (c *CatesModel) Add(data map[string]interface{}) *DMLResult {

	return nil
}

func (c *CatesModel) DelByOid(oid string) *DMLResult {

	return nil
}

func (c *CatesModel) GetByOid(oid string) []*Cate {

	return nil
}

func (c *CatesModel) UpdateByOid(oid string, data map[string]interface{}) *DMLResult {

	return nil
}
