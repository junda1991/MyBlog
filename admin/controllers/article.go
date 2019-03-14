package controllers

import (
	"MyBlog/admin/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController)ShowArticle()  {
	c.Ctx.WriteString("hello")
}
//article/list
func (c *ArticleController)ShowArticlelist(){
	o:=orm.NewOrm()
	var articles []models.Article
	_,err:=o.QueryTable("Article").All(&articles)
	if err !=nil{
		beego.Info("查询所有文章出错")
	}
	//beego.Info(articles)
	c.Data["articles"]=articles
	c.TplName="article/list.html"
}
//article/add
func (c *ArticleController)ShowAddarticle() {
	o:=orm.NewOrm()
	var cates []models.Cate
	_,err:=o.QueryTable("cate").All(&cates)
	if err !=nil{
		beego.Info("查询所有栏目出错")
	}
	beego.Info(cates)
	c.Data["cates"]=cates
	c.TplName="article/add.html"
}
//post //article/add
func (c *ArticleController)HandleAddarticle() {

	//1拿到数据
	title := c.GetString("title")
	author := c.GetString("author")
	desc := c.GetString("desc")
	cateid := c.GetString("cateid")
	content := c.GetString("content")

	f, h, err := c.GetFile("pic")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("pic", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	//beego.Info(title,author,desc,content,cateid,h.Filename)

	////2对数据进行校验
	if title == "" ||author == ""||desc == ""||cateid == ""||content == ""{
		beego.Info("数据为空")
		c.Data["msg"]="数据有为空,重新添加"
		c.TplName="article/error.html"
		return
	}
	////3插入数据库
	o := orm.NewOrm()
	ar := models.Article{}
	ar.Title=title
	ar.Author=author
	ar.Desc=desc
	ar.Cateid=cateid
	ar.Content=content
	ar.Pic= h.Filename

	err = o.Read(&ar, "Title")
	if err != nil {
		beego.Info("查询失败")
		//4插入
		_,err:=o.Insert(&ar)
		if err!=nil{
			beego.Info("插入失败",err)
		}
		c.Redirect("/article/list",302)
		return
	}
	//c.Ctx.WriteString("添加失败,已有用户名<a>aaa</a>")
	c.Data["msg"]="添加失败,已有标题名名"
	c.TplName="article/error.html"
}

//article/del
func (c *ArticleController)HandleDelarticle(){
	//1获取id
	id,err:=c.GetInt("id")
	//beego.Info(id,err)
	//2执行删除
	o:=orm.NewOrm()
	ca:=models.Article{Id:id}
	err=o.Read(&ca)
	if err!=nil{
		beego.Info("查询错误")
	}
	o.Delete(&ca)
	c.Redirect("/article/list",302)
}

//get article/update
func (c *ArticleController)ShowUpdatearticle(){
	//1获取id
	id,err:=c.GetInt("id")
	beego.Info(id,err)
	//查询
	o := orm.NewOrm()
	ar := models.Article{Id: id}

	err = o.Read(&ar)
	if err!=nil {
		fmt.Println("查询不到",err)
	}
	beego.Info(ar)
	c.Data["ar"]=ar


	o=orm.NewOrm()
	var cates []models.Cate
	_,err=o.QueryTable("cate").All(&cates)
	if err !=nil{
		beego.Info("查询所有栏目出错")
	}
	beego.Info(cates)
	c.Data["cates"]=cates
	c.TplName="article/edit.html"
}

func (c *ArticleController)HandleUpdatearticle(){
	id, _:=c.GetInt("id")
	title:=c.GetString("title")
	author := c.GetString("author")
	desc := c.GetString("desc")
	cateid := c.GetString("cateid")
	content := c.GetString("content")

	f, h, err := c.GetFile("pic")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("pic", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建


	if title==""{
		c.Data["msg"]="栏目不能为空"
		c.Redirect("/article/list",302)
		return
	}
	o := orm.NewOrm()
	ar := models.Article{}
	ar.Id=id
	ar.Title=title
	ar.Author=author
	ar.Desc=desc
	ar.Cateid=cateid
	ar.Content=content
	ar.Pic= h.Filename

	_,err=o.Update(&ar)
	if err!=nil{
		beego.Info("更新失败",err)
	}

	c.Redirect("/article/list",302)
	beego.Info(id)

}

