/*
 * html模板调用
 */
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("tpl/*")
    r.GET("/sai", func (c *gin.Context) {
        c.JSON(200, gin.H {
            "msg":"wow",
        })
    })

    r.GET("/test", func (c *gin.Context) {
        c.HTML(http.StatusOK, "test.tpl", gin.H {
            "title":"sai",
            "data":"what",
        })
    })

    r.Run(":8089")
}