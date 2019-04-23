/**
* BasicAuth 中间件
*/
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var pws = gin.H{
	"foo":    gin.H{"email": "foo@13sai.com", "phone": "1300000000"},
	"aaa": gin.H{"email": "aaa@13sai.com", "phone": "15666666666"},
	"sai":   gin.H{"email": "sai@13sai.com", "phone": "18399999999"},
}

func main() {
	r := gin.Default()

	auth := r.Group("/admin", gin.BasicAuth(gin.Accounts {
		"foo": "bar",
		"aaa": "111",
		"sai": "666",
	}))

	auth.GET("/pw", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if pw, ok := pws[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "pw": pw})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "pw": "error"})
		}
	})

	r.Run(":8010")
}