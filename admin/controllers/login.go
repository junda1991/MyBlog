package controllers

import (
	"MyBlog/admin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}
//登录
func (c *LoginController) ShowLogin() {
	c.TplName="admin/login.html"
}
func (c *LoginController) Handlelogin() {
	username:=c.GetString("username")
	password:=c.GetString("password")
	//c.Ctx.WriteString(username+password)
	beego.Info(username,password)

	o := orm.NewOrm()

	ad := models.Admin{}
	ad.Username=username
	err := o.Read(&ad,"username")
	if err!=nil {
		c.Data["msg"]="用户名不存在"
		c.TplName="admin/loginerror.html"
		return
	}
	if ad.Password!=password{
		c.Data["msg"]="密码错误"
		c.TplName="admin/loginerror.html"
		return
	}
	beego.Info("查询成功")
	c.SetSession("username",username)
	//登录后台
	c.Redirect("/admin",302)

}
//退出登录
func (c *LoginController) Handlelogout(){
	//删除session
	c.DelSession("username")
	c.Redirect("/login",302)
}
