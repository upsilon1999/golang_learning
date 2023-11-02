package three

/*
	①引入第三方包，观察项目结构和go.mod文件的变化
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getData() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}