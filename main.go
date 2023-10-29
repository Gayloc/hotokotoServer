package main

import (
	"encoding/json"
	"main/hitokoto"
	"main/server"
	"os"
)

func main() {
	database := hitokoto.DataBase{}
	file, err := os.ReadFile("hitokoto.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &database)
	if err != nil {
		panic(err)
	}
	server.StartServer(database)
}
