package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harish1907/go-jwt-project/controllers"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("users/signup", controllers.SignUp)
	route.POST("users/login", controllers.Login)
}
