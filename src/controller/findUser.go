package controller

import "github.com/gin-gonic/gin"

func FindUserById(c *gin.Context) {
c.JSON(200, gin.H{
	"Result": "Find User By Id",
})
}

func FindUserByEmail(c *gin.Context) {
c.JSON(200, gin.H{
	"Result": "Find User By Email",
})

}