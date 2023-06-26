package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/login", func(ctx *gin.Context) {
			ctx.AsciiJSON(http.StatusOK, gin.H{
				"msg": "success",
			})
		})
		v1.GET("/users/:name/*action", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":           "success",
				"user_name":     ctx.Param("name"),
				"action":        ctx.Param("action"),
				"uri_full_path": ctx.FullPath(),
			})
			// /v1/users/shaouai/show
			// {"action":"/show","msg":"success","uri_full_path":"/v1/users/:name/*action","user_name":"shaouai"}
		})
	}

	r.Run(":8080")
}
