package routes

import (
	controller "github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/controllers"
	middleware "github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
