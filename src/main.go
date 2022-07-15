package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通过字典模拟 DB
//var db = make(map[string]string)
var db = map[string]string{"one": "1", "foo": "bar"}

//一定要手动在goland 中手动打开go module开关，否则不会识别下载下来的包
func setupRouter() *gin.Engine {
	// 初始化 Gin 框架默认实例，该实例包含了路由、中间件以及配置信息
	router := gin.Default()

	// Ping 测试路由
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 获取用户数据路由 这个是url里面的参数，而不是params里面的参数
	router.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// 需要 HTTP 基本授权认证的子路由群组设置
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // 用户名:foo 密码:bar
		"manu": "123", // 用户名:manu 密码:123
	}))

	// 保存用户信息路由
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// 解析并验证 JSON 格式请求数据
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return router
}

func main() {
	// 设置路由信息
	router := setupRouter()
	// 启动服务器并监听 8080 端口
	router.Run(":1234")
}
