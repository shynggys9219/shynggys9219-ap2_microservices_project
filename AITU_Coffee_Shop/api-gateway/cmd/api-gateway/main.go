package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		resp, err := http.Post("http://localhost:8080/api/v1/clients/", "application/json", c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err})
			return
		}
		defer resp.Body.Close()

		_, err = http.Post("http://localhost:8002/increment", "application/json", nil)
		if err != nil {
			log.Printf("http.Post: statistics: %v", err)
		}

		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)

	})

	r.POST("/login", func(c *gin.Context) {
		resp, err := http.Post("http://localhost:8001/login", "application/json", c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "user service down"})
			return
		}
		defer resp.Body.Close()
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	})

	r.GET("/stats", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:8002/stats")
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "stats service down"})
			return
		}
		defer resp.Body.Close()
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	})

	r.Run(":8000")
}
