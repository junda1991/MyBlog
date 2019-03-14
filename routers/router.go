package routers

import (
	"MyBlog/admin/controllers"
	"MyBlog/index/controller"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterFunc)
	beego.InsertFilter("/cate/*",beego.BeforeRouter,FilterFunc)
	beego.InsertFilter("/article/*",beego.BeforeRouter,FilterFunc)

    beego.Router("/admin", &controllers.MainController{})
	//beego.Router("/login", &controllers.MainController{},"get:ShowLogin")

	beego.Router("/admin/list", &controllers.MainController{},"get:ShowAdminlist")
	beego.Router("/admin/add", &controllers.MainController{},"get:ShowAddadmin;post:HandleAddadmin")
	beego.Router("/admin/del", &controllers.MainController{},"get:HandleDel")
	beego.Router("/admin/update", &controllers.MainController{},"get:ShowUpdate;post:HandleUpdate")


	beego.Router("/cate",&controllers.CateController{},"get:ShowCate")
	beego.Router("/cate/list", &controllers.CateController{},"get:ShowCatelist")
	beego.Router("/cate/add", &controllers.CateController{},"get:ShowAddcate;post:HandleAddcate")
	beego.Router("/cate/del", &controllers.CateController{},"get:HandleDelcate")
	beego.Router("/cate/update", &controllers.CateController{},"get:ShowUpdatecate;post:HandleUpdatecate")

	beego.Router("/article",&controllers.ArticleController{},"get:ShowArticle")
	beego.Router("/article/list", &controllers.ArticleController{},"get:ShowArticlelist")
	beego.Router("/article/add", &controllers.ArticleController{},"get:ShowAddarticle;post:HandleAddarticle")
	beego.Router("/article/del", &controllers.ArticleController{},"get:HandleDelarticle")
	beego.Router("/article/update", &controllers.ArticleController{},"get:ShowUpdatearticle;post:HandleUpdatearticle")


	beego.Router("/login", &controllers.LoginController{},"get:ShowLogin;post:Handlelogin")
	beego.Router("/logout", &controllers.LoginController{},"get:Handlelogout")

	beego.Router("/index",&controller.IndexController{} ,"get:ShowIndex")
	beego.Router("/",&controller.IndexController{} ,"get:ShowIndex")
	beego.Router("/index/article",&controller.IndexController{} ,"get:ShowIndexarticle")


}
var FilterFunc = func(ctx *context.Context){
	username:=ctx.Input.Session("username")
	if username==nil{
		ctx.Redirect(302,"/login")
	}
}