package routes

import (
	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/middleware"
	controller "github.com/Aditya-sairam/golang-jwt-project/Recruitment/controllers"
	"github.com/gin-gonic/gin"
)

func JobRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("recruitment/job_posting", controller.CreateJobPosting())
	incomingRoutes.GET("recruitment/job_list", controller.ListJobs())
	incomingRoutes.GET("/recruitment/:job_id", controller.GetJob())

}
