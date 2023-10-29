package server

import (
	"encoding/json"
	"log"
	"main/hitokoto"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartServer(database hitokoto.DataBase) {
	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) {
		result := database.FindByType(ctx.Query("type"))
		ctx.JSON(http.StatusOK, result)
	})

	engine.POST("/", authAdmin(), func(ctx *gin.Context) {
		hitokoto := hitokoto.Hitokoto{Id: 0, Hitokoto: "", HitokotoType: "", Reviewer: 0, From_who: "", Length: 0}
		err := ctx.ShouldBind(&hitokoto)
		if err != nil {
			log.Fatal(err.Error())
		}

		if database.AddItem(hitokoto) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
			err = save(database)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "提交内容不符合规范",
			})
		}
	})

	engine.DELETE("/", authAdmin(), func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))

		if database.DelItem(id) && (err == nil) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
			err = save(database)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "删除失败",
			})
		}
	})

	engine.PUT("/", authAdmin(), func(ctx *gin.Context) {
		hitokoto := hitokoto.Hitokoto{Id: 0, Hitokoto: "", HitokotoType: "", Reviewer: 0, From_who: "", Length: 0}
		err := ctx.ShouldBind(&hitokoto)
		if err != nil {
			log.Fatal(err.Error())
		}
		id, err := strconv.Atoi(ctx.Query("id"))

		if database.EditItem(id, hitokoto) && (err == nil) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
			err = save(database)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "修改失败",
			})
		}
	})

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func save(database hitokoto.DataBase) error {
	content, err := json.Marshal(database)
	if err != nil {
		return err
	}

	os.WriteFile("hitokoto.json", content, 0666)
	if err != nil {
		return err
	}
	return nil
}

func authAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("User-Group") != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "无权访问",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
