package connectdb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://10.190.221.93:30717")

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
}
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	//fmt.Println("ok")
	return mgoCli
}
