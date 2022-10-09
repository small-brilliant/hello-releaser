package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 520, "msg": "谭谭我好喜欢你吖"})
	})
	r.Run(":8080")
}
