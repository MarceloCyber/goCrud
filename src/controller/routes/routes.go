package routes

import (
	"github.com/MarceloCyber/goCrud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes( r *gin.RouterGroup) {
	r.GET("/", controller.Principal)
r.GET("/getUserById/:user_id", controller.FindUserById)
r.GET("/getUserByEmail/:user_email", controller.FindUserByEmail)
r.POST("/createUser", controller.CreateUser)
r.PUT("/updateUser/:user_id", controller.UpdateUser)
r.DELETE("/deleteUser/:user_id", controller.DeleteUser)
}