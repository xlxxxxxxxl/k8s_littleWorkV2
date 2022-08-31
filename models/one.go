package models

type Xl struct {
	Id   string `bson:"_id"`
	Name string
}

type Re struct {
	Data interface{} `json:"插入姓名为"`
	//ID   interface{} `json:"ID"`
	//Code    string
	Message string `json:"提示"`
}

type Showall struct {
	Data1    string `json:"所有保存的姓名为"`
	Message1 string `json:"提示"`
}

type Deleteall struct {
	Message2 string `json:"提示"`
	DeCount  int64  `json:"总计删除"`
}

type ChooseOne struct {
	Message3 string      `json:"提示"`
	D        interface{} `json:"幸运儿名字为"`
}

type LastLucky struct {
	//ID   string
	Name string `json:"上次中奖者为"`
	Time string `json:"中奖时间为"`
}
