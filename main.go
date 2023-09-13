package main

import (
	"blog-sync/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	service.Init()
	r := gin.Default()
	r.GET("/sync", func(c *gin.Context) {
		doSync(c)
	})
	r.POST("/sync", func(c *gin.Context) {
		doSync(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func doSync(c *gin.Context) {
	action := c.Query("action")
	if action == "" {
		action = "blog"
	}
	service.DoPull(action)
	c.JSON(http.StatusOK, gin.H{
		"message": action + " sync OK",
	})
}
