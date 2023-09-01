package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/sayhello/:name", func(ctx *gin.Context) {
		name := ctx.Params.ByName("name")
		if name == "" {
			name = "user default"
		}
		ctx.JSON(http.StatusOK, DNSDiscoverySayHello(name))
	})

	return r
}
