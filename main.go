package main

import (
	"encoding/json"
	"main/hitokoto"
	"main/server"
	"os"
)

func main() {
	database := hitokoto.DataBase{}
	if !FileExists("hitokoto.json") {
		DefaultDatabase, err := os.ReadFile("DefaultDatabase.json")
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("hitokoto.json", DefaultDatabase, 0666)
		if err != nil {
			panic(err)
		}
	}
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

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
