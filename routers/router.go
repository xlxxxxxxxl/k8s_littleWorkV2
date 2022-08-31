package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"ks8-demo/connectdb"
	"ks8-demo/controllers"
)

func init() {
	connectdb.GetMgoCli()
	beego.Router("/", &controllers.MainController{})
	beego.Router("/addUser", &controllers.AddUser{})
	beego.Router("/delete", &controllers.DeleteUser{})
	beego.Router("/showAll", &controllers.ShowAll{})
	beego.Router("/chooseOne", &controllers.ChooseOne{})
	beego.Router("/showLast", &controllers.Lucky{})
}
