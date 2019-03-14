package controllers

import (
	"MyBlog/admin/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//Session
	username:=c.GetSession("username")
	beego.Info(username)
	if username ==nil{
		c.Redirect("/login",302)
	}
	c.Data["username"]=username
	c.TplName = "admin/index.html"

}

//addmin/list
func (c *MainController) ShowAdminlist() {
	o:=orm.NewOrm()
	var admins []models.Admin
	_,err:=o.QueryTable("Admin").All(&admins)
	if err !=nil{
		beego.Info("查询所有文章出错")
	}
	beego.Info(admins)
	c.Data["sss"]=admins
	c.TplName="admin/list.html"
}
//addmin/add
func (c *MainController) ShowAddadmin() {
	c.TplName="admin/add.html"
}

//post addmin/add
func (c *MainController) HandleAddadmin() {
	//c.TplName="admin/add.html"
	//username := c.GetString("username")
	//if username == "" {
	//	c.Ctx.WriteString("username is empty")
	//	return
	//}
	//password := c.GetString("password")
	//if password == "" {
	//	c.Ctx.WriteString("password is empty")
	//	return
	//}
	////1有orm对象
	//o:=orm.NewOrm()
	////2有一个要插入数据的结构体对象
	//admin:=models.Admin{}
	////3对结构体赋值
	//admin.Username=username
	//admin.Password=password
	////4插入
	//_,err:=o.Insert(&admin)
	//if err!=nil{
	//	beego.Info("插入失败",err)
	//}
	//1拿到数据
	username := c.GetString("username")
	password := c.GetString("password")
	beego.Info(username, password)
	//c.Ctx.WriteString("提交后的页面")
	//2对数据进行校验
	if username == "" || password == "" {
		beego.Info("数据为空")
		c.Data["msg"]="数据为空,请重新注册"
		c.TplName="admin/error.html"
		return
	}
	//3插入数据库
	o := orm.NewOrm()
	admin := models.Admin{}
	admin.Username=username
	admin.Password = password
	err := o.Read(&admin, "username")
	if err != nil {
		beego.Info("查询失败")
		//4插入
		_,err:=o.Insert(&admin)
		if err!=nil{
			beego.Info("插入失败",err)
		}
		c.Redirect("/admin/list",302)
		return
	}
	//c.Ctx.WriteString("添加失败,已有用户名<a>aaa</a>")
	c.Data["msg"]="添加失败,已有用户名"
	c.TplName="admin/error.html"

}

//admin/del
func (c *MainController) HandleDel(){
	//1获取id
	id,err:=c.GetInt("id")
	beego.Info(id,err)
	//2执行删除
	o:=orm.NewOrm()
	ad:=models.Admin{Id:id}
	err=o.Read(&ad)
	if err!=nil{
		beego.Info("查询错误")
	}
	o.Delete(&ad)
	c.Redirect("/admin/list",302)
}

func (c *MainController) ShowUpdate(){
	//1获取id
	id,err:=c.GetInt("id")
	beego.Info(id,err)
	//查询
	o := orm.NewOrm()
	user := models.Admin{Id: id}

	err = o.Read(&user)
	if err!=nil {
		fmt.Println("查询不到",err)
	}
	beego.Info(user)
	c.Data["ssw"]=user
	c.TplName="admin/edit.html"
}

func (c *MainController) HandleUpdate(){
	id, _:=c.GetInt("id")
	username:=c.GetString("username")
	password:=c.GetString("password")
	if username==""||password==""{
		c.Data["msg"]="用户名或者密码不能为空"
		c.Redirect("/admin/list",302)
		return
	}
	o := orm.NewOrm()
	user := models.Admin{Id: id}
	if o.Read(&user) == nil {
		user.Username = username
		user.Password= password
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
		}
		c.Redirect("/admin/list",302)
	}
	beego.Info(id,username,password)
}