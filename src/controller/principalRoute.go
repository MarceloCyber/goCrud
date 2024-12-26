package controller

import "github.com/gin-gonic/gin"

func Principal(c *gin.Context) {
	c.JSON(200, gin.H{
		"Result": "Principal",

	})
}