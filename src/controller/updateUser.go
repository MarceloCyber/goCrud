package controller

import "github.com/gin-gonic/gin"

func UpdateUser(c *gin.Context){
	c.JSON(200, gin.H{
		"Result": "Update User",
	})
}
