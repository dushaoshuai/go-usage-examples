package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "api-examples/swag/docs"
)

type request struct {
	// 用户 ID
	UserId uint64 `json:"user_id,string"`

	// 请求来源，h5
	Source string `json:"source"`

	// 请求的 session ID，如果是登录用户需要带上
	Sid string `json:"sid"`

	// 不同请求对应的请求包体
	Data interface{} `json:"data"`
}

type queryReq struct {
	Phone string
	Age   int
}

type queryRsp struct {
	Phone  string
	Age    int
	Height int
}

// sdfkajdlajldjal ajsldj ajdlfaj
// @Summary  管理员更改自己密码
// @Tags     h5
// @Param    data  body      request{data=queryReq}  true  "请求参数"
// @Success  200   {object}  queryRsp
// @Router   /admin/account/change_pwd [post]
func foo() {}

// sjlajdlfjaldjfalj
// @Summary  创建创作参考
// @Tags     创作参考
// @Param    data  body      request{data=queryReq}  true  "请求参数"
// @Success  200   {object}  queryRsp
// @Router   /foo/bar/foo/bar [post]
func bar() {}

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":9091") // 8080
}
