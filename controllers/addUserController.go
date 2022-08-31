package controllers

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	"ks8-demo/connectdb"
	"ks8-demo/models"
)

type AddUser struct {
	beego.Controller
}

func (A *AddUser) Get() {
	var result models.Re
	addName := A.GetString("name")

	name := models.Xl{addName, addName}
	client := connectdb.GetMgoCli()
	collection := client.Database("newdb").Collection("runood")
	_, err := collection.InsertOne(context.TODO(), name)
	if err != nil {
		//result.ID = res.InsertedID
		result.Data = addName
		result.Message = "姓名已经存在 插入失败"
	} else {
		//result.ID = res.InsertedID
		result.Data = addName
		result.Message = "插入成功"
	}
	A.Data["json"] = result
	A.ServeJSON()
}
