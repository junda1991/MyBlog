package main

import (
	_ "MyBlog/routers"
	"github.com/astaxie/beego"
	_ "MyBlog/admin/models"
)

func main() {
	beego.Run()
}

