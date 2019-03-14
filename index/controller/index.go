package controller

import (
	"MyBlog/admin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) ShowIndex() {
	o:=orm.NewOrm()
	var cates []models.Cate
	_,err:=o.QueryTable("cate").All(&cates)
	if err !=nil{
		beego.Info("查询所有栏目出错")
	}
	beego.Info(cates)
	c.Data["cates"]=cates
	c.TplName="index/index.html"
}
func (c *IndexController) ShowIndexarticle(){
	o:=orm.NewOrm()
	var cates []models.Cate
	_,err:=o.QueryTable("cate").All(&cates)
	if err !=nil{
		beego.Info("查询所有栏目出错")
	}
	beego.Info(cates)
	c.Data["cates"]=cates

	//更新查询cateid的文章
	cateid:=c.GetString("cateid")
	beego.Info(cateid)
	o = orm.NewOrm()
	article:=models.Article{}
	article.Cateid=cateid
	err =o.Read(&article,"Cateid")
	beego.Info(article)

	c.Data["sss"]=article


	c.TplName="index/article.html"
}
