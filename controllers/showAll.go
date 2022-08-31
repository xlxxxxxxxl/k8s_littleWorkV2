package controllers

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"ks8-demo/connectdb"
	"ks8-demo/models"
)

type ShowAll struct {
	beego.Controller
}

func (c *ShowAll) Get() {
	client := connectdb.GetMgoCli()
	collection := client.Database("newdb").Collection("runood")
	res, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("1", err)
	}
	var result models.Showall
	var alldb []models.Xl

	//result = new(models.Showall)

	res2 := res.All(context.TODO(), &alldb)
	//fmt.Println("所有人名如下：\n")
	if res2 != nil {
		fmt.Println("2", res2)
	}
	for i := range alldb {
		result.Data1 += "  " + alldb[i].Name
	}
	//result.Data = alldb[].Name
	result.Message1 = "显示成功"
	c.Data["json"] = &result
	c.ServeJSON()
}
