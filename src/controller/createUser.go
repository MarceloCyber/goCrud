package controller

import (

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func CreateUser(c *gin.Context){

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"Result":   "Created User",
	})
}