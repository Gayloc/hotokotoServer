package server

import (
	"encoding/json"
	"log"
	"main/auth"
	"main/hitokoto"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartServer(database hitokoto.DataBase) {
	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) { //获得全部数据或者按类型获取数据
		result := database.FindByType(ctx.Query("type"))
		ctx.JSON(http.StatusOK, result)
	})

	engine.POST("/auth", auth.GetToken) //获取token

	engine.POST("/", auth.JWTAuthMiddleware(), func(ctx *gin.Context) { //添加数据
		hitokoto := hitokoto.Hitokoto{Id: 0, Hitokoto: "", HitokotoType: "", Reviewer: 0, From_who: "", Length: 0}
		err := ctx.ShouldBind(&hitokoto)
		if err != nil {
			log.Fatal(err.Error())
		}

		if database.AddItem(hitokoto) {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    ctx.MustGet("username").(string), //从上下文中读取用户名信息
				"message": "ok",
			})
			err = save(database) //执行保存
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    ctx.MustGet("username").(string),
				"message": "提交内容不符合规范",
			})
		}
	})

	engine.DELETE("/", auth.JWTAuthMiddleware(), func(ctx *gin.Context) { //删除数据
		id, err := strconv.Atoi(ctx.Query("id"))

		if database.DelItem(id) && (err == nil) {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    ctx.MustGet("username").(string),
				"message": "ok",
			})
			err = save(database) //执行保存
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    ctx.MustGet("username").(string),
				"message": "删除失败",
			})
		}
	})

	engine.PUT("/", auth.JWTAuthMiddleware(), func(ctx *gin.Context) { //修改数据
		hitokoto := hitokoto.Hitokoto{Id: 0, Hitokoto: "", HitokotoType: "", Reviewer: 0, From_who: "", Length: 0}
		err := ctx.ShouldBind(&hitokoto)
		if err != nil {
			log.Fatal(err.Error())
		}
		id, err := strconv.Atoi(ctx.Query("id"))

		if database.EditItem(id, hitokoto) && (err == nil) {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    ctx.MustGet("username").(string),
				"message": "ok",
			})
			err = save(database) //执行保存
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"user":    ctx.MustGet("username").(string),
				"message": "修改失败",
			})
		}
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func save(database hitokoto.DataBase) error {
	content, err := json.Marshal(database) //序列化
	if err != nil {
		return err
	}

	os.WriteFile("hitokoto.json", content, 0666) //保存（覆盖原有内容）
	if err != nil {
		return err
	}
	return nil
}
