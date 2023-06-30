package main

import (
	gin "github.com/dushaoshuai/explore-gin-routes-tree"
)

func doc(ctx *gin.Context)           {}
func docInstall(ctx *gin.Context)    {}
func docGetStart(ctx *gin.Context)   {}
func docCmd(ctx *gin.Context)        {}
func docUpload(ctx *gin.Context)     {}
func refMod(ctx *gin.Context)        {}
func refSpec(ctx *gin.Context)       {}
func userLogin(ctx *gin.Context)     {}
func userLogout(ctx *gin.Context)    {}
func companyLogin(ctx *gin.Context)  {}
func companyLogout(ctx *gin.Context) {}

func main() {
	r := gin.Default()

	r.GET("/doc", doc)

	docGroup := r.Group("/doc")
	{
		docGroup.GET("/install", docInstall)
		docGroup.GET("/tutorial/getting-started", docGetStart)
		docGroup.GET("/cmd", docCmd)
		docGroup.POST("/upload", docUpload)
	}

	refGroup := r.Group("/ref")
	{
		refGroup.GET("/mod", refMod)
		refGroup.GET("/spec", refSpec)
	}

	r.POST("/user/login/:user", userLogin)
	r.POST("/user/logout/:user", userLogout)
	r.POST("/company/login/*company", companyLogin)
	r.POST("/company/logout/*company", companyLogout)

	r.Run() // default :8080
}
