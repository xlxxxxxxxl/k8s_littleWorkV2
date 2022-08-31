package controllers

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"ks8-demo/connectdb"
	"ks8-demo/models"
)

type DeleteUser struct {
	beego.Controller
}

func (a *DeleteUser) Get() {
	client := connectdb.GetMgoCli()
	collection := client.Database("newdb").Collection("runood")
	res, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("1", err)
	}

	var result models.Deleteall
	result.DeCount = res.DeletedCount
	result.Message2 = "数据已删除，"
	a.Data["json"] = &result
	a.ServeJSON()
}
