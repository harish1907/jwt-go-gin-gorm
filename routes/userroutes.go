package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harish1907/go-jwt-project/controllers"
	"github.com/harish1907/go-jwt-project/middleware"
)

func UserRoutes(route *gin.Engine) {
	route.Use(middleware.Authenticate)
	route.GET("/users", controllers.GetUsers)
	route.GET("/users/:user_id", controllers.GetUser)
}
