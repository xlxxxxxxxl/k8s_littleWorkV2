package controllers

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"ks8-demo/connectdb"
	"ks8-demo/models"
)

type Lucky struct {
	beego.Controller
}

func (A *Lucky) Get() {
	client := connectdb.GetMgoCli()
	collection := client.Database("newdb").Collection("lucky")
	res, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("1", err)
	}
	var result models.LastLucky
	var alldb []models.LastLucky

	res2 := res.All(context.TODO(), &alldb)
	if res2 != nil {
		fmt.Println(res2)
	}

	var L int

	for i := range alldb {
		if i == len(alldb)-1 {
			L = i
		}
	}

	result.Name = alldb[L].Name
	result.Time = alldb[L].Time
	//fmt.Println(result.D)
	A.Data["json"] = &result
	A.ServeJSON()
	//result. = nil
}
