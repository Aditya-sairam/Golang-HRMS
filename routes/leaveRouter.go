package routes

import (
	controller "github.com/Aditya-sairam/golang-jwt-project/controllers"
	"github.com/Aditya-sairam/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func LeaveRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/leave_request/", controller.LeaveRequest())

}
