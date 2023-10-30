package main

import (
	"encoding/json"
	"main/hitokoto"
	"main/server"
	"os"
)

func main() {
	database := hitokoto.DataBase{}
	file, err := os.ReadFile("hitokoto.json") //读取本地文件
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &database) //反序列化
	if err != nil {
		panic(err)
	}
	server.StartServer(database)
}
