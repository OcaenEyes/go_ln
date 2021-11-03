package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/sonJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "cn",
			"name": "gzy",
		}
		c.JSON(200, data)
	})
	r.Run()
}
