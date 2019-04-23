/*
 * http方法
 */
 package main

 import (
	"net/http"
	"github.com/gin-gonic/gin"
 )

func main() {
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run(":9002")
}