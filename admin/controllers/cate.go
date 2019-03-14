package controllers

import (
	"MyBlog/admin/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CateController struct {
	beego.Controller
}

func (c *CateController)ShowCate()  {
	c.Ctx.WriteString("cate")
}

//cate/list
func (c *CateController) ShowCatelist() {
	o:=orm.NewOrm()
	var cates []models.Cate
	_,err:=o.QueryTable("cate").All(&cates)
	if err !=nil{
		beego.Info("查询所有文章出错")
	}
	beego.Info(cates)
	c.Data["sss"]=cates
	c.TplName="cate/list.html"
}

//cate/add
func (c *CateController) ShowAddcate() {
	c.TplName="cate/add.html"
}
//post cate/add 路径
func (c *CateController) HandleAddcate() {

	//1拿到数据
	catename := c.GetString("catename")

	beego.Info(catename)
	//c.Ctx.WriteString("提交后的页面")
	//2对数据进行校验
	if catename == "" {
		beego.Info("数据为空")
		c.Data["msg"]="数据为空,重新添加"
		c.TplName="cate/error.html"
		return
	}
	//3插入数据库
	o := orm.NewOrm()
	cate := models.Cate{}
	cate.Catename=catename

	err := o.Read(&cate, "Catename")
	if err != nil {
		beego.Info("查询失败")
		//4插入
		_,err:=o.Insert(&cate)
		if err!=nil{
			beego.Info("插入失败",err)
		}
		c.Redirect("/cate/list",302)
		return
	}
	//c.Ctx.WriteString("添加失败,已有用户名<a>aaa</a>")
	c.Data["msg"]="添加失败,已有栏目名"
	c.TplName="cate/error.html"

}

//cate/del
func(c *CateController) HandleDelcate(){
	//1获取id
	id,err:=c.GetInt("id")
	beego.Info(id,err)
	//2执行删除
	o:=orm.NewOrm()
	ca:=models.Cate{Id:id}
	err=o.Read(&ca)
	if err!=nil{
		beego.Info("查询错误")
	}
	o.Delete(&ca)
	c.Redirect("/cate/list",302)
}

//get cate/update
func(c *CateController) ShowUpdatecate(){
	//1获取id
	id,err:=c.GetInt("id")
	beego.Info(id,err)
	//查询
	o := orm.NewOrm()
	cate := models.Cate{Id: id}

	err = o.Read(&cate)
	if err!=nil {
		fmt.Println("查询不到",err)
	}
	beego.Info(cate)
	c.Data["cate"]=cate


	c.TplName="cate/edit.html"
}
//post cate/update
func (c *CateController)HandleUpdatecate()  {
	id, _:=c.GetInt("id")
	catename:=c.GetString("catename")

	if catename==""{
		c.Data["msg"]="栏目不能为空"
		c.Redirect("/cate/list",302)
		return
	}
	o := orm.NewOrm()
	cate := models.Cate{}
	cate.Id=id
	cate.Catename=catename
	err := o.Read(&cate, "catename")
	if err != nil {
		beego.Info("查询失败")

		//4插入
		_,err:=o.Update(&cate)
		if err!=nil{
			beego.Info("更新失败",err)
		}

	//o := orm.NewOrm()
	//cate := models.Cate{Id: id}
	//if o.Read(&cate) == nil {
	//	cate.Catename = catename
	//	if num, err := o.Update(&cate); err == nil {
	//		fmt.Println(num)
	//	}
	//	c.Redirect("/cate/list",302)
	//}
		c.Redirect("/cate/list",302)
	beego.Info(id,catename)
}
	c.Data["msg"]="栏目不重复"
	c.TplName="cate/error2.html"
	}