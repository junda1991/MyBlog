package models

import _ "github.com/go-sql-driver/mysql"

//文章结构体
type Article struct {
	Id int
	Title string
	Author string
	Desc string
	Pic  string
	Content string
	Cateid string
}

//类型表
type Cate struct {
	Id int
	Catename string
	//Article []*Article `orm:"reverse(many)"`
}