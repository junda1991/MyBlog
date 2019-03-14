package models

import (
	"github.com/astaxie/beego/orm"
)
import _ "github.com/go-sql-driver/mysql"

type Admin struct {
	Id int
	Username string
	Password string
}

func init(){
	//设置数据库的基本信息
	orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/myblog?charset=utf8")
	//映射model数据
	orm.RegisterModel(new(Admin),new(Cate),new(Article))
	//生成表
	orm.RunSyncdb("default",false,true)
}