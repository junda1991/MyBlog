package main

import (
	_ "MyBlog/admin/models"
	_ "MyBlog/routers"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	beego.AddFuncMap("ShowPrePage",HandlePrePage)
	beego.AddFuncMap("ShowNextPage",HandleNextPage)

	beego.Run()
}
func HandlePrePage(data int)string{

	pageindex:=data-1
	pageindex1:=strconv.Itoa(pageindex)
	return pageindex1
}
func HandleNextPage(data int)string{

	pageindex:=data+1
	pageindex1:=strconv.Itoa(pageindex)
	return pageindex1
}
