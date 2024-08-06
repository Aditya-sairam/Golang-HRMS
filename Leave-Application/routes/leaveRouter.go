package routes

import (
	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/middleware"
	controller "github.com/Aditya-sairam/golang-jwt-project/Leave-Application/controllers"
	"github.com/gin-gonic/gin"
)

func LeaveAppRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/leave_request/", controller.LeaveRequest())

}
