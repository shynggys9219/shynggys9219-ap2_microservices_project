package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

var userCount int32 = 0

func main() {
	r := gin.Default()

	r.POST("/increment", func(c *gin.Context) {
		atomic.AddInt32(&userCount, 1)
		c.Status(http.StatusOK)
	})

	r.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"registered_users": atomic.LoadInt32(&userCount),
		})
	})

	r.Run(":8002")
}
