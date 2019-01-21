package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Love Yourself"
	c.Data["Email"] = "gorishabhio@gmail.com"
	c.TplName = "index.tpl"
}


type LoveController struct {
	beego.Controller
}

func ( c *LoveController) Get() {
	c.Data["LoveQuote"] = "self love is the best form of love"
	c.Data["CreateLove"] = "Learn to create love with these books!"
	c.TplName = "love.tpl"
}
