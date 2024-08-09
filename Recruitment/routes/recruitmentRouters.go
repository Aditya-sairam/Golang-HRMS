package routes

import (
	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/middleware"
	controller "github.com/Aditya-sairam/golang-jwt-project/Recruitment/controllers"
	"github.com/gin-gonic/gin"
)

func JobRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/job_posting", controller.CreateJobPosting())

}
